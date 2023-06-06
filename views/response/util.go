package response

type UtilDBQueryResultResponse struct {
	ColumnNames []string `json:"column_names"`
	Rows []map[string]interface{} `json:"rows"`
	Sql string `json:"sql"`
}
