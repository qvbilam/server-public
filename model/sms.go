package model

import "time"

type SmsChannel struct {
	IDModel
	Name            string `gorm:"name"`
	Type            string
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	Sign            string `json:"sign"`
	TemplateCode    string `json:"template_code"`
	Count           int64  `json:"count"`
	ErrorCount      int64  `json:"error_count"`
	DateModel
}

type SmsCodeRecord struct {
	IDModel
	ChannelID int64      `gorm:"type:varchar(255) not null default '';comment:渠道id;index"`
	Mobile    string     `gorm:"type:varchar(255) not null default '';comment:手机号;index"`
	Type      string     `gorm:"type:varchar(255) not null default '';comment:验证码类型"`
	Code      string     `gorm:"type:varchar(255) not null default '';comment:验证码"`
	Param     string     `gorm:"type:varchar(255) not null default '';comment:参数"`
	Response  string     `gorm:"type:varchar(255) not null default '';comment:响应"`
	IsUsed    bool       `gorm:"type:tinyint(1) not null default 0;comment:是否使用"`
	ExpiredAt *time.Time `gorm:"comment:过期时间"`
	DateModel
}
