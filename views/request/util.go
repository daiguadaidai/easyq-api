package request

type UtilEncreptRequest struct {
	Data string `json:"data" form:"data" binding:"required"`
}

type UtilDecryptRequest struct {
	Data string `json:"data" form:"data" binding:"required"`
}

type UtilSqlFingerprintRequest struct {
	Statements []string `json:"statements" form:"statements" binding:"required"`
}
