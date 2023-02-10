package dao

import (
	"vue-next-admin-go/db"
	"vue-next-admin-go/modles"
)

type UserInfoModel modles.UserInfo

func (that *UserInfoModel) Info() (info modles.Info, err error) {
	query := `select * from testops.user where name=? and password=?`
	err = db.DB.Get(&info, query, that.UserName, that.PassWord)
	return
}

func (that *UserInfoModel) List() (list []*modles.Info, err error) {
	query := `select * from testops.user where 1=1`

	err = db.DB.Select(&list, query)
	return
}
