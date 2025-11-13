package entity

// Agreement 对应数据表 t_agreement
type Agreement struct {
	Id           int64  `orm:"id"`
	Name         string `orm:"name"`
	MajorVersion string `orm:"major_version"`
	MinorVersion string `orm:"minor_version"`
	PatchVersion string `orm:"patch_version"`
	Version      int    `orm:"version"`
	Content      string `orm:"content"`
	CreatedAt    int64  `orm:"created_at"`
	UpdatedAt    int64  `orm:"updated_at"`
}

// UserAgreement 对应数据表 t_user_agreement
type UserAgreement struct {
	Id            int64  `orm:"id"`
	UserID        string `orm:"user_id"`
	AgreementID   int64  `orm:"agreement_id"`
	AgreementName string `orm:"agreement_name"`
	Agreed        int    `orm:"agreed"`
	CreatedAt     int64  `orm:"created_at"`
}
