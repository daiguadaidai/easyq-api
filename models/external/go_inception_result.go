package external

import (
	"fmt"
	"strconv"
)

const (
	GoInceptionErrorLevelSuccess = iota
	GoInceptionErrorLevelWarn
	GoInceptionErrorLevelError
)

type GoInceptionResult struct {
	OrderId      int64   `json:"order_id"`
	Stage        string  `json:"stage"`
	ErrorLevel   int64   `json:"error_level"`
	StageStatus  string  `json:"stage_status"`
	ErrorMessage string  `json:"error_message"`
	ExecSql      string  `json:"exec_sql"`
	LogicSql     string  `json:"logic_sql"`
	AffectedRows int64   `json:"affected_rows"`
	Sequence     string  `json:"sequence"`
	BackupDbname string  `json:"backup_dbname"`
	ExecuteTime  float64 `json:"execute_time"`
	Sqlsha1      string  `json:"sqlsha1"`
	BackupTime   float64 `json:"backup_time"`
}

func NewGoInceptionResult(orderId, stage, errorLevel, stageStatus, errorMessage, execSql, affectedRows, sequence, backupDbname, executeTime, sqlsha1, backupTime string) (*GoInceptionResult, error) {
	newOrderId, err := strconv.ParseInt(orderId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("orderId转化成int64出错", err.Error())
	}
	newErrorLevel, err := strconv.ParseInt(errorLevel, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("errorLevel转化成int64出错", err.Error())
	}
	NewAffectedRows, err := strconv.ParseInt(affectedRows, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("affectedRows转化成int64出错", err.Error())
	}
	NewExecuteTime, err := strconv.ParseFloat(executeTime, 64)
	if err != nil {
		return nil, fmt.Errorf("executeTime转化成int64出错", err.Error())
	}
	NewBackupTime, err := strconv.ParseFloat(backupTime, 64)
	if err != nil {
		return nil, fmt.Errorf("backupTime转化成int64出错", err.Error())
	}

	return &GoInceptionResult{
		OrderId:      newOrderId,
		Stage:        stage,
		ErrorLevel:   newErrorLevel,
		StageStatus:  stageStatus,
		ErrorMessage: errorMessage,
		ExecSql:      execSql,
		LogicSql:     execSql,
		AffectedRows: NewAffectedRows,
		Sequence:     sequence,
		BackupDbname: backupDbname,
		ExecuteTime:  NewExecuteTime,
		Sqlsha1:      sqlsha1,
		BackupTime:   NewBackupTime,
	}, nil
}
