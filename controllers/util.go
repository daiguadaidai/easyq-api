package controllers

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/models/external"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/utils/sqlparser"
	"github.com/daiguadaidai/easyq-api/views/request"
	"github.com/daiguadaidai/easyq-api/views/response"
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

func (this *UtilController) DBResult() (*response.UtilDBQueryResultResponse, error) {
	// 随机生成字段名数
	randN := utils.RandN(100)
	if randN == 0 {
		return &response.UtilDBQueryResultResponse{
			ColumnNames: make([]string, 0, 0),
			Rows:        make([]map[string]interface{}, 0, 0),
			Sql:         "SELECT * FROM tbl_cnt_0",
		}, nil
	}

	sqlStr := fmt.Sprintf("SELECT * FROM tbl_cnt_%v", randN)

	// 生成字段名 col_name_rand
	columnNames := make([]string, 0, randN)
	for i := 0; i < randN; i++ {
		columnNames = append(columnNames, fmt.Sprintf("col_name_%v", i))
	}

	// 生成随机内容
	randRowCnt := utils.RandN(1000)
	if randRowCnt == 0 {
		randRowCnt = 100
	}
	rows := make([]map[string]interface{}, 0, randRowCnt)
	for i := 0; i < randRowCnt; i++ {
		row := make(map[string]interface{})
		randLen := utils.RandN(300)
		if randLen == 0 {
			randLen = 50
		}

		for _, columnName := range columnNames {
			row[columnName] = utils.RandString(randLen)
		}

		rows = append(rows, row)
	}

	return &response.UtilDBQueryResultResponse{
		Sql:         sqlStr,
		ColumnNames: columnNames,
		Rows:        rows,
	}, nil
}
