package model

type Video struct {
	IDModel
	UserModel
	Channel     string `gorm:"type:varchar(255) not null default '';comment:存储渠道"`
	BusinessId  string `gorm:"type:varchar(255) not null default '';comment:第三方业务id;index"`
	Sha1        string `gorm:"type:varchar(255) not null default '';comment:文件sha1;index"`
	Url         string `gorm:"type:varchar(255);not null default '';comment:存储地址"`
	Size        int64  `gorm:"type:int(11);not null default 0;comment:大小"`
	Duration    int64  `gorm:"type:int(11);not null default 0;comment:时长"`
	Status      string `gorm:"type:varchar(11);not null default '';comment:状态: -1.blocked,0.wait,1.normal"`
	ContentType string `gorm:"type:varchar(255);not null default '';comment:内容类型"`
	Extra       string `gorm:"type:varchar(255);not null default '';comment:拓展字段"`
	DateModel
	DeletedModel
}

type File struct {
	IDModel
	Channel     string `gorm:"type:varchar(255) not null default '';comment:存储渠道"`
	Sha1        string `gorm:"type:varchar(255) not null default '';comment:文件sha1;index"`
	Url         string `gorm:"type:varchar(255);not null default '';comment:存储地址"`
	ContentType string `gorm:"type:varchar(255);not null default '';comment:内容类型"`
	Size        int64  `gorm:"type:int(11);not null default 0;comment:大小"`
	Extra       string `gorm:"type:varchar(255);not null default '';comment:拓展字段"`
	Callback    string `gorm:"type:varchar(1024);not null default '';comment:回调内容"`
	DateModel
	DeletedModel
}

type DownloadLog struct {
	IDModel
	UserModel
	Type   string `gorm:"type:varchar(255) not null default '';comment:类型:image.图片,video.视频"`
	FileId int64  `gorm:"type:int(255);not null default 0;comment:资源id;index:idx_app_file"`
	DateModel
}
