package entity

// Agreement 对应数据表 t_agreement
type Agreement struct {
	Id           int64  `orm:"id"`
	Name         string `orm:"name"`
	MajorVersion int    `orm:"major_version"`
	MinorVersion int    `orm:"minor_version"`
	PatchVersion int    `orm:"patch_version"`
	VersionCode  int    `orm:"version_code"`
	Status       int    `orm:"status"`
	Content      string `orm:"content"`
	PublishedAt  int64  `orm:"published_at"`
	CreatedAt    int64  `orm:"created_at"`
	UpdatedAt    int64  `orm:"updated_at"`
}

// UserAgreement 对应数据表 t_user_agreement
type UserAgreement struct {
	Id            int64  `orm:"id"`
	UserID        string `orm:"user_id"`
	AgreementID   int64  `orm:"agreement_id"`
	AgreementName string `orm:"agreement_name"`
	VersionCode   int    `orm:"version_code"`
	Agreed        int    `orm:"agreed"`
	CreatedAt     int64  `orm:"created_at"`
}
