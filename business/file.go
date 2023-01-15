package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"public/global"
	"public/model"
)

type FileBusiness struct {
	ID          int64
	Sha1        string
	Size        int64
	Url         string
	ContentType string
	Extra       string
	Channel     string
}

func (b *FileBusiness) Exists() bool {
	i := model.File{}
	if res := global.DB.Where("sha1", b.Sha1).Select("id").First(&i); res.RowsAffected == 0 {
		return false
	}
	return true
}

func (b *FileBusiness) GetById() *model.File {
	i := model.File{}
	if res := global.DB.Where("id", b.ID).First(&i); res.RowsAffected == 0 {
		return nil
	}
	return &i
}

func (b *FileBusiness) GetBySha1() *model.File {
	i := model.File{}
	if res := global.DB.Where("sha1", b.Sha1).First(&i); res.RowsAffected == 0 {
		return nil
	}
	return &i
}

func (b *FileBusiness) Create() (*model.File, error) {
	tx := global.DB.Begin()
	f := model.File{}
	res := tx.Where(&model.File{Sha1: b.Sha1}).Find(&tx)
	if res.RowsAffected != 0 {
		tx.Rollback()
		return &f, nil
	}

	f.Size = b.Size
	f.Sha1 = b.Sha1
	f.Url = b.Url
	f.ContentType = b.ContentType
	f.Extra = b.Extra
	f.Channel = b.Channel

	if res := global.DB.Save(&f); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "创建失败")
	}

	return &f, nil
}

func (b *FileBusiness) Update() error {
	i := model.File{}
	if res := global.DB.Where("id", b.ID).First(&i); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "视频不存在")
	}

	b.EntityToUpdateModel(&i)
	if res := global.DB.Save(&i); res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "修改视频信息失败")
	}
	return nil
}

func (b *FileBusiness) EntityToUpdateModel(file *model.File) {
	if b.Url != "" {
		file.Url = b.Url
	}

	if b.ContentType != "" {
		file.ContentType = b.ContentType
	}
}
