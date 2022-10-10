package business

import (
	"file/global"
	"file/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VideoBusiness struct {
	ID          int64
	UserID      int64
	BusinessID  string
	Sha1        string
	Url         string
	Size        int64
	Duration    int64
	Status      string
	ContentType string
	Expand      string
	Channel     string
}

func (b *VideoBusiness) Exists() bool {
	v := model.Video{}
	condition := model.Video{}
	if b.ID != 0 {
		condition.ID = b.ID
	}
	if b.BusinessID != "" {
		condition.BusinessId = b.BusinessID
	}
	if b.Sha1 != "" {
		condition.Sha1 = b.Sha1
	}
	if res := global.DB.Where(condition).Select("id").First(&v); res.RowsAffected == 0 {
		return false
	}
	return true
}

func (b *VideoBusiness) GetById() *model.Video {
	v := model.Video{}
	if res := global.DB.Where("id", b.ID).First(&v); res.RowsAffected == 0 {
		return nil
	}
	return &v
}

func (b *VideoBusiness) GetByBusinessId() *model.Video {
	v := model.Video{}
	if res := global.DB.Where("business_id", b.ID).First(&v); res.RowsAffected == 0 {
		return nil
	}
	return &v
}

func (b *VideoBusiness) Create() (*model.Video, error) {
	v := model.Video{
		UserModel:   model.UserModel{UserID: b.UserID},
		BusinessId:  b.BusinessID,
		Sha1:        b.Sha1,
		Url:         b.Url,
		Size:        b.Size,
		Duration:    b.Duration,
		Status:      b.Status,
		ContentType: b.ContentType,
		Expand:      b.Expand,
		Channel:     b.Channel,
	}
	if res := global.DB.Save(&v); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "创建失败")
	}

	return &v, nil
}

func (b *VideoBusiness) Update() error {
	v := model.Video{}
	if res := global.DB.Where("id", b.ID).First(&v); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "视频不存在")
	}

	b.EntityToUpdateModel(&v)
	if res := global.DB.Save(&v); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "修改视频信息失败")
	}
	return nil
}

func (b *VideoBusiness) UpdateByBusinessId() error {
	v := model.Video{}
	if res := global.DB.Where("business_id", b.BusinessID).First(&v); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "视频不存在")
	}

	b.EntityToUpdateModel(&v)
	if res := global.DB.Save(&v); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "修改视频信息失败")
	}
	return nil
}

func (b *VideoBusiness) EntityToUpdateModel(video *model.Video) {
	if b.Url != "" {
		video.Url = b.Url
	}

	if b.Size != 0 {
		video.Size = b.Size
	}

	if b.Duration != 0 {
		video.Duration = b.Duration
	}

	if b.Status != "" {
		video.Status = b.Status
	}

	if b.ContentType != "" {
		video.ContentType = b.ContentType
	}
}
