package controllers

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/models/external"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/utils/sqlparser"
	"github.com/daiguadaidai/easyq-api/views/request"
)

type UtilController struct {
	Ctx *contexts.GlobalContext
}

func NewDefaultUtilController(ctx *contexts.GlobalContext) *UtilController {
	return &UtilController{Ctx: ctx}
}

func NewUtilController(ctx *contexts.GlobalContext) *UtilController {
	return &UtilController{Ctx: ctx}
}

// 加密
func (this *UtilController) Encrypt(req *request.UtilEncreptRequest) (string, error) {
	return utils.Encrypt(req.Data)
}

// 解密
func (this *UtilController) Decrypt(req *request.UtilDecryptRequest) (string, error) {
	return utils.Decrypt(req.Data)
}

func (this *UtilController) GetSqlFingerprints(req *request.UtilSqlFingerprintRequest) ([]*external.SqlFingerprint, error) {
	fingers := make([]*external.SqlFingerprint, 0, len(req.Statements))
	for i, statement := range req.Statements {
		// 解析每一条输入到sql
		sqlStrs, err := sqlparser.SqlToMulti(statement)
		if err != nil {
			return nil, fmt.Errorf("解析第%v条sql出错. %v, 输入sql为: %v", i, err.Error(), statement)
		}

		// 输入的一条sql可能包含多条sql
		for _, sqlStr := range sqlStrs {
			ormalized, digest := sqlparser.NormalizeDigest(sqlStr)
			finger := &external.SqlFingerprint{
				GroupId:             i,
				OriStatment:         sqlStr,
				PlaseholderStatment: ormalized,
				Fingerprint:         digest,
			}
			fingers = append(fingers, finger)
		}
	}

	return fingers, nil
}
