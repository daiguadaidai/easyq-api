package external

type SqlFingerprint struct {
	GroupId             int    `json:"group_id"`
	OriStatment         string `json:"ori_statment"`
	PlaseholderStatment string `json:"plaseholder_statment"`
	Fingerprint         string `json:"fingerprint"`
}
