package model

type File struct {
	IDModel
	UserModel
	AppId       int64  `gorm:"type:int(11);not null default 0;comment:应用;index:idx_app_sha"`
	FileSHA     string `gorm:"type:varchar(255);not null default '';comment:应用;index:idx_app_sha"`
	FilePath    string `gorm:"type:varchar(255);not null default '';comment:文件路径"`
	FileSize    int64  `gorm:"type:int(11);not null default 0;comment:文件大小"`
	ContentType string `gorm:"type:varchar(255);not null default '';comment:内容类型"`
	Visible
	DateModel
	DeletedModel
}

type Config struct {
	IDModel
	AppId       int64  `gorm:"type:int(11);not null default 0;comment:应用;uniqueIndex"`
	Key         string `gorm:"type:varchar(255);not null default '';comment:应用id"`
	Secrect     string `gorm:"type:varchar(255);not null default '';comment:密钥"`
	Host        string `gorm:"type:varchar(255);not null default '';comment:地址"`
	ExpireTime  int64  `gorm:"type:int(11);not null default 0;comment:文件大小"`
	UploadDir   string `gorm:"type:varchar(255);not null default '';comment:文件目录"`
	CallBackUrl string `gorm:"type:varchar(255);not null default '';comment:回调地址"`
}

type UploadLog struct {
	IDModel
	UserModel
	AppId   int64  `gorm:"type:int(11);not null default 0;comment:应用;index:idx_app_file"`
	FileId  int64  `gorm:"type:int(255);not null default 0;comment:资源id;index:idx_app_file"`
	Channel string `gorm:"type:varchar(255);not null default '';comment:通道;"`
	DateModel
}

type DownLoadLog struct {
	IDModel
	UserModel
	AppId  int64 `gorm:"type:int(11);not null default 0;comment:应用;index:idx_app_file"`
	FileId int64 `gorm:"type:int(255);not null default 0;comment:资源id;index:idx_app_file"`
	DateModel
}
