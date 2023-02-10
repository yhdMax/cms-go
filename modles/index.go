package modles

type Info struct {
	Name     string `db:"name" json:"name"`
	PassWord string `db:"password" json:"password"`
	ID       string `db:"id" json:"-"`
	Status   string `db:"status" json:"-"`
	Group    string `db:"group" json:"-"`
}

type UserInfo struct {
	UserName string `db:"name" json:"name"`
	PassWord string `db:"password" json:"password"`
	Code     string `db:"code" json:"code"`
}
