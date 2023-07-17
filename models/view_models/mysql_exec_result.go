package view_models

type MysqlExecResult struct {
	ExecSql     string                   `json:"exec_sql" form:"exec_sql"`
	ColumnNames []string                 `json:"column_names" form:"column_names"`
	Rows        []map[string]interface{} `json:"rows" form:"rows"`
	IsErr       bool                     `json:"is_err" form:"is_err"`
	ErrMsg      string                   `json:"err_msg" form:"err_msg"`
}
