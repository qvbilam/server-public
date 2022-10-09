package business

import (
	"file/global"
	"file/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImageBusiness struct {
	ID          int64
	UserID      int64
	BusinessID  string
	Sha1        string
	Url         string
	Size        int64
	Width       int64
	Height      int64
	Status      string
	ContentType string
	Expand      string
}

func (b *ImageBusiness) Exists() bool {
	i := model.Image{}
	if res := global.DB.Where("business_id", b.BusinessID).Select("id").First(&i); res.RowsAffected == 0 {
		return false
	}
	return true
}

func (b *ImageBusiness) GetById() *model.Image {
	i := model.Image{}
	if res := global.DB.Where("id", b.ID).First(&i); res.RowsAffected == 0 {
		return nil
	}
	return &i
}

func (b *ImageBusiness) GetByBusinessId() *model.Image {
	i := model.Image{}
	if res := global.DB.Where("business_id", b.ID).First(&i); res.RowsAffected == 0 {
		return nil
	}
	return &i
}

func (b *ImageBusiness) Create() (*model.Image, error) {
	i := model.Image{
		UserModel:   model.UserModel{UserID: b.UserID},
		BusinessId:  b.BusinessID,
		Sha1:        b.Sha1,
		Url:         b.Url,
		Size:        b.Size,
		Width:       b.Width,
		Height:      b.Height,
		Status:      b.Status,
		ContentType: b.ContentType,
		Expand:      b.Expand,
	}
	if res := global.DB.Save(&i); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "创建失败")
	}

	return &i, nil
}

func (b *ImageBusiness) Update() error {
	i := model.Image{}
	if res := global.DB.Where("id", b.ID).First(&i); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "视频不存在")
	}

	b.EntityToUpdateModel(&i)
	if res := global.DB.Save(&i); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "修改视频信息失败")
	}
	return nil
}

func (b *ImageBusiness) UpdateByBusinessId() error {
	i := model.Image{}
	if res := global.DB.Where("business_id", b.BusinessID).First(&i); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "视频不存在")
	}

	b.EntityToUpdateModel(&i)
	if res := global.DB.Save(&i); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "修改视频信息失败")
	}
	return nil
}

func (b *ImageBusiness) EntityToUpdateModel(image *model.Image) {
	if b.Url != "" {
		image.Url = b.Url
	}

	if b.Size != 0 {
		image.Size = b.Size
	}

	if b.Width != 0 {
		image.Width = b.Width
	}

	if b.Height != 0 {
		image.Height = b.Height
	}

	if b.Status != "" {
		image.Status = b.Status
	}

	if b.ContentType != "" {
		image.ContentType = b.ContentType
	}
}
