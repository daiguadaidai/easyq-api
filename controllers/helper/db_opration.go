package helper

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/config"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/gdbc"
	"strconv"
	"strings"
)

var ignoreDatabases = map[string]struct{}{
	"information_schema": {},
	"__tencentdb__":      {},
	"xa":                 {},
	"mysql":              {},
	"performance_schema": {},
	"sys":                {},
	"sysdb":              {},
}

// 过滤掉可以忽略掉数据库
func FilterIgnoreDatabases(databases []string) []string {
	newDatabases := make([]string, 0, len(databases))
	for _, database := range databases {
		if _, ok := ignoreDatabases[database]; !ok {
			newDatabases = append(newDatabases, database)
		}
	}

	return newDatabases
}

func FindDBNameByVipPort(ctx *contexts.GlobalContext, vipPort string) ([]string, error) {
	items := strings.Split(vipPort, ":")
	if len(items) < 2 {
		return nil, fmt.Errorf("错误的VipPort. vipPort: %v", vipPort)
	}
	host := items[0]
	port, err := strconv.ParseInt(items[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("解析VipPort. 端口出错: %v. %v", vipPort, err)
	}

	return FindDBNameByHostPort(ctx, host, port)
}

func FindDBNameByHostPort(ctx *contexts.GlobalContext, host string, port int64) ([]string, error) {
	cfg := config.NewMysqlConfig(host, port, ctx.Cfg.ApiConfig.AdminMysqlUser, ctx.Cfg.ApiConfig.AdminMysqlPassword, "")
	db, err := gdbc.GetMySQLDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("临时获取数据库链接出错. host:port: %v:%v. %v", host, port, err)
	}
	defer db.Close()

	dbNames, err := dao.NewDBOperationDao(db).ShowDatabases()
	if err != nil {
		return nil, fmt.Errorf("查询 SHOW DATABASES 出错. %v", err)
	}

	return dbNames, err
}
