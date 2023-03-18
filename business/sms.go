package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"public/global"
	"public/model"
	"public/utils"
	"public/utils/ali/sms"
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
	b.Code = utils.RandomNumber(4)

	// todo 验证ip频率
	// todo 验证手机号频率
	tx := global.DB.Begin()
	// 获取通道
	channel := model.SmsChannel{}
	if res := tx.First(&channel); res.RowsAffected == 0 {
		tx.Rollback()
		return status.Errorf(codes.NotFound, "缺少短信网关配置")
	}

	// todo 默认只支持阿里云短信
	c := sms.Client{
		AccessKeyId:     channel.AccessKeyID,
		AccessKeySecret: channel.AccessKeySecret,
		SignName:        channel.Sign,
		TemplateCode:    channel.TemplateCode,
	}

	res, err := c.SendCode(b.Mobile, b.Code)
	if err != nil {
		smsChannelFail(channel.ID, tx)
		tx.Rollback()
		return status.Errorf(codes.Internal, "短信网关异常")
	}

	// 增加记录
	expiredAt := time.Now().Add(time.Minute * 5)
	if res := tx.Model(&model.SmsCodeRecord{
		ChannelID: channel.ID,
		Mobile:    b.Mobile,
		Type:      b.Type,
		Code:      b.Code,
		Response:  *res,
		ExpiredAt: &expiredAt,
	}); res.RowsAffected == 0 {
		tx.Rollback()
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

func smsChannelSuccess(id int64, tx *gorm.DB) int64 {
	if tx == nil {
		tx = global.DB
	}

	res := tx.Where(&model.SmsChannel{IDModel: model.IDModel{ID: id}}).Updates(map[string]interface{}{
		"count": gorm.Expr("count + ?", 1),
	})
	return res.RowsAffected
}

func smsChannelFail(id int64, tx *gorm.DB) int64 {
	if tx == nil {
		tx = global.DB
	}

	res := tx.Where(&model.SmsChannel{IDModel: model.IDModel{ID: id}}).Updates(map[string]interface{}{
		"count":       gorm.Expr("count + ?", 1),
		"error_count": gorm.Expr("error_count + ?", 1),
	})
	return res.RowsAffected
}
