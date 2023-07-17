package gdbc

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"sync"
	"testing"
)

func Test_GetAndReleaseMySQLDB(t *testing.T) {
	host := "172.22.72.136"
	port := 3306
	username := "root"
	password := "123456"
	database := "easydb"
	autocommit := true
	timeout := 10
	charset := "utf8mb4"
	minOpen := 8
	maxOpen := 8

	pool, err := Open(host, int64(port), username, password, database, charset, autocommit, int64(timeout), int64(minOpen), int64(maxOpen))
	if err != nil {
		t.Fatal(err.Error())
	}
	defer pool.Close()

	wg := new(sync.WaitGroup)
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, tag int) {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				db, err := pool.Get()
				if err != nil {
					logs.Error(err.Error())
					return
				}
				var threadId int64
				if err := db.QueryRow("SELECT CONNECTION_ID()").Scan(&threadId); err != nil {
					logs.Error(err.Error())
					return
				}
				fmt.Printf("%d. tag: %d, ThreadId: %d\n", j, tag, threadId)
				if err := pool.Release(db); err != nil {
					logs.Error("归还链接出错. %s", err.Error())
				}
			}
		}(wg, i)
	}

	wg.Wait()
}
