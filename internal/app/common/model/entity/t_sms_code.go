package entity

type TSmsCode struct {
	Id           int64  `orm:"id"`
	BusinessType int    `orm:"business_type"`
	Phone        string `orm:"phone"`
	Code         string `orm:"code"`
	Status       int    `orm:"status"`
	CreatedAt    int64  `orm:"created_at"`
	ExpiredAt    int64  `orm:"expired_at"`
	UpdatedAt    int64  `orm:"updated_at"`
}
