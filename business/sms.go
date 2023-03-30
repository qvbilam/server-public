package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"public/cache"
	"public/global"
	"public/model"
	"public/utils"
	"time"
)

type SmsBusiness struct {
	Mobile    string `json:"mobile"`
	Type      string `json:"type"`
	Code      string `json:"code"`
	CheckCode string `json:"check_code"`
	ClientIP  string `json:"client_ip"`
}

func (b *SmsBusiness) Send() error {
	// 随机验证码
	b.Code = utils.RandomNumber(4)

	if !utils.IsMobile(b.Mobile) {
		return status.Errorf(codes.InvalidArgument, "手机号错误")
	}

	rs := cache.RedisServer{}
	if !rs.SetMobileSmsLock(b.Mobile, b.Type, 50) {
		return status.Errorf(codes.ResourceExhausted, "请稍后再试")
	}

	tx := global.DB.Begin()
	// 获取通道
	channel := model.SmsChannel{}
	if res := tx.First(&channel); res.RowsAffected == 0 {
		tx.Rollback()
		rs.DelMobileSmsLock(b.Mobile, b.Type)
		return status.Errorf(codes.NotFound, "缺少短信网关配置")
	}

	// 暂只支持阿里云短信
	//c := sms.Client{
	//	AccessKeyId:     channel.AccessKeyID,
	//	AccessKeySecret: channel.AccessKeySecret,
	//	SignName:        channel.Sign,
	//	TemplateCode:    channel.TemplateCode,
	//}

	//res, err := c.SendCode(b.Mobile, b.Code)
	testRes := "test"
	res := &testRes
	var err error

	if err != nil {
		smsChannelFail(channel.ID, nil)
		tx.Rollback()
		rs.DelMobileSmsLock(b.Mobile, b.Type)
		return status.Errorf(codes.Internal, "短信网关异常"+err.Error())
	}

	// 增加记录
	expiredAt := time.Now().Add(time.Minute * 5)
	if res := tx.Save(&model.SmsCodeRecord{
		ChannelID: channel.ID,
		Mobile:    b.Mobile,
		Type:      b.Type,
		Code:      b.Code,
		Response:  *res,
		ExpiredAt: &expiredAt,
	}); res.RowsAffected == 0 {
		tx.Rollback()
		smsChannelFail(channel.ID, nil)
		return status.Errorf(codes.Internal, "发送失败")
	}

	// 更新渠道
	if count := smsChannelSuccess(channel.ID, tx); count == 0 {
		tx.Rollback()
		return status.Errorf(codes.Internal, "发送异常")
	}

	tx.Commit()
	return nil
}

func (b *SmsBusiness) Check() error {
	entity := model.SmsCodeRecord{}
	if res := global.DB.Where(&model.SmsCodeRecord{Mobile: b.Mobile, Type: b.Type}).Last(&entity); res.RowsAffected == 0 {
		return status.Errorf(codes.InvalidArgument, "验证码错误")
	}

	if entity.IsUsed == true {
		return status.Errorf(codes.InvalidArgument, "验证码已使用")
	}
	if entity.Code != b.CheckCode {
		return status.Errorf(codes.InvalidArgument, "验证码错误")
	}
	if entity.ExpiredAt.Unix() <= time.Now().Unix() {
		return status.Errorf(codes.InvalidArgument, "验证码过期")
	}

	entity.IsUsed = true
	global.DB.Save(entity)

	return nil
}

func smsChannelSuccess(id int64, tx *gorm.DB) int64 {
	if tx == nil {
		tx = global.DB
	}
	res := tx.Model(&model.SmsChannel{}).Where(&model.SmsChannel{IDModel: model.IDModel{ID: id}}).Updates(map[string]interface{}{
		"count": gorm.Expr("count + ?", 1),
	})
	return res.RowsAffected
}

func smsChannelFail(id int64, tx *gorm.DB) int64 {
	if tx == nil {
		tx = global.DB
	}

	res := tx.Model(&model.SmsChannel{}).Where(&model.SmsChannel{IDModel: model.IDModel{ID: id}}).Updates(map[string]interface{}{
		"count":       gorm.Expr("count + ?", 1),
		"error_count": gorm.Expr("error_count + ?", 1),
	})
	return res.RowsAffected
}
