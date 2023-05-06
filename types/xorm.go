package types

import (
	timex "github.com/liuxiaobopro/gobox/time"
)

type ModelExtends struct {
	CreateAt  timex.JsonTime `xorm:"comment('创建时间') DATETIME created" json:"create_at"`
	UpdateAt  timex.JsonTime `xorm:"comment('更新时间') DATETIME updated" json:"update_at"`
	CreateUid int            `xorm:"comment('创建人id') INT(11) default 0" json:"create_uid"`
	UpdateUid int            `xorm:"comment('更新人id') INT(11) default 0" json:"update_uid"`
}

type ModelCTExtends struct {
	CreateAt timex.JsonTime `xorm:"comment('创建时间') DATETIME created" json:"create_at"`
}

type ModelCUExtends struct {
	CreateAt timex.JsonTime `xorm:"comment('创建时间') DATETIME created" json:"create_at"`
	UpdateAt timex.JsonTime `xorm:"comment('更新时间') DATETIME updated" json:"update_at"`
}
