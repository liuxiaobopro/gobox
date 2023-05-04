package types

import (
	timex "github.com/liuxiaobopro/gobox/time"
)

type ModelExtends struct {
	CreateAt  timex.JsonTime `xorm:"comment('创建时间') DATETIME created" json:"createAt"`
	UpdateAt  timex.JsonTime `xorm:"comment('更新时间') DATETIME updated" json:"updateAt"`
	CreateUid int            `xorm:"comment('创建人id') INT(11) default 0" json:"createUid"`
	UpdateUid int            `xorm:"comment('更新人id') INT(11) default 0" json:"updateUid"`
}

type ModelCTExtends struct {
	CreateAt timex.JsonTime `xorm:"comment('创建时间') DATETIME created" json:"createAt"`
}

type ModelCUExtends struct {
	CreateAt timex.JsonTime `xorm:"comment('创建时间') DATETIME created" json:"createAt"`
	UpdateAt timex.JsonTime `xorm:"comment('更新时间') DATETIME updated" json:"updateAt"`
}
