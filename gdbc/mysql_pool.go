package gdbc

import (
	"database/sql"
	"fmt"
	"github.com/daiguadaidai/easyq-api/logger"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	DB_CONFIG_USERNAME = "root"
	DB_CONFIG_PASSWORD = ""
	DB_CONFIG_HOST     = "127.0.0.1"
	DB_CONFIG_PORT     = 3306
	DB_CONFIG_CHARSET  = "utf8mb4"
	DB_CONFIG_TIMEOUT  = 10
)

type dbConfig struct {
	Username   string
	Password   string
	Host       string
	Port       int64
	DBName     string
	Charset    string
	Timeout    int64
	Autocommit bool
}

func newDBConfig(
	host string,
	port int64,
	username string,
	password string,
	dbName string,
	charset string,
	autoCommit bool,
	timeout int64,
) *dbConfig {
	cfg := new(dbConfig)
	// 初始化 host
	if len(strings.TrimSpace(host)) == 0 { // 没有指定 host
		logger.M.Warn("没有指定host. 使用默认 host:%s", DB_CONFIG_HOST)
		cfg.Host = DB_CONFIG_HOST
	} else {
		cfg.Host = host
	}

	// 初始化 port
	if port < 1 { // port 小于1
		logger.M.Warn("指定的Port:%d, 小于1, 使用默认 port:%d", port, DB_CONFIG_PORT)
		cfg.Port = DB_CONFIG_PORT
	} else {
		cfg.Port = port
	}

	// 初始化 username
	if len(strings.TrimSpace(username)) == 0 {
		logger.M.Warn("没有指定username, 使用默认 username:%s", DB_CONFIG_USERNAME)
		cfg.Username = DB_CONFIG_USERNAME
	} else {
		cfg.Username = username
	}

	// 初始化 password
	if len(strings.TrimSpace(password)) == 0 {
		logger.M.Warn("没有指定password, 使用默认 password:%s", DB_CONFIG_PASSWORD)
		cfg.Password = DB_CONFIG_PASSWORD
	} else {
		cfg.Password = password
	}

	// 初始化 charset
	if len(strings.TrimSpace(charset)) == 0 {
		logger.M.Warn("没有指定charset, 使用默认 charset:%s", DB_CONFIG_CHARSET)
		cfg.Charset = DB_CONFIG_CHARSET
	} else {
		cfg.Charset = charset
	}

	// 初始化 autotommit
	if !autoCommit {
		logger.M.Warn("指定 autocommit=0. (一般应用使用的是 autocommit=1, 请谨慎考虑.)")
	}
	cfg.Autocommit = autoCommit

	// 初始化timeout
	if timeout <= 0 {
		logger.M.Warn("没有指定timeout, 使用默认 timeout:%s", DB_CONFIG_TIMEOUT)
		cfg.Timeout = DB_CONFIG_TIMEOUT
	} else {
		cfg.Timeout = timeout
	}

	cfg.DBName = dbName

	logger.M.Debug("链接配置为: %s", cfg.String())
	return cfg
}

func (this *dbConfig) addr() string {
	return fmt.Sprintf("%s:%d", this.Host, this.Port)
}

func (this *dbConfig) String() string {
	return fmt.Sprintf("host: %s, port: %d, username: %s, password: ******, charset: %s, automommit: %t",
		this.Host, this.Port, this.Username, this.Charset, this.Autocommit)
}

func (this *dbConfig) DSN() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=%v&allowOldPasswords=1&timeout=%vs&autocommit=%v&parseTime=True&loc=Local",
		this.Username,
		this.Password,
		this.Host,
		this.Port,
		this.DBName,
		this.Charset,
		this.Timeout,
		this.Autocommit,
	)
}

const (
	MYSQL_POOL_MIN_OPEN       = 1
	MYSQL_POOL_MAX_OPEN       = 1
	MYSQL_POOL_MAX_OPEN_LIMIT = 1000
)

type MySQLPool struct {
	Role string
	sync.Mutex
	dbChan  chan *sql.DB
	Cfg     *dbConfig
	minOpen int64
	maxOpen int64
	numOpen int64
}

func Open(
	host string,
	port int64,
	username string,
	password string,
	dbName string,
	charset string,
	autoCommit bool,
	timeout int64,
	minOpen int64,
	maxOpen int64,
	role string,
) (*MySQLPool, error) {
	cfg := newDBConfig(host, port, username, password, dbName, charset, autoCommit, timeout)

	p := new(MySQLPool)
	p.Cfg = cfg
	p.dbChan = make(chan *sql.DB, MYSQL_POOL_MAX_OPEN_LIMIT)

	// 初始化 链接打开最小值
	if minOpen < 1 {
		p.minOpen = MYSQL_POOL_MIN_OPEN
		logger.M.Warn("最小使用链接数不能小于1, 默认值:%d", MYSQL_POOL_MIN_OPEN)
	} else {
		p.minOpen = minOpen
	}

	// 初始化 链接打开最大值
	if maxOpen < 1 {
		p.maxOpen = MYSQL_POOL_MAX_OPEN
		logger.M.Warn("最大使用链接数不能小于1, 默认值:%d", MYSQL_POOL_MIN_OPEN)
	} else {
		p.maxOpen = maxOpen
	}

	// 判断最大链接数是否大于系统允许的最大链接
	if maxOpen > MYSQL_POOL_MAX_OPEN_LIMIT {
		p.maxOpen = MYSQL_POOL_MAX_OPEN_LIMIT
		logger.M.Warn("指定最大链接数:%d, 大于系统允许最大连接数:%d, 默认设置最大链接数为:%d", maxOpen, MYSQL_POOL_MAX_OPEN_LIMIT, MYSQL_POOL_MAX_OPEN_LIMIT)
	}

	// 最小链接数不能大于最大链接数
	if p.minOpen > p.maxOpen {
		p.minOpen = p.maxOpen
		logger.M.Warn("最小链接数:%d 大于 最大链接数:%d. 设置最小链接为:%d", p.minOpen, p.maxOpen, p.maxOpen)
	}

	p.Role = role

	return p, nil
}

// 关闭连接池
func (this *MySQLPool) Close() {
	close(this.dbChan)
	for db := range this.dbChan {
		this.Lock()
		if err := this.closeConn(db); err != nil {
			logger.M.Error("链接关闭出错. %s", err.Error())
		}
		this.Unlock()
	}
}

func (this *MySQLPool) IncrNumOpen() {
	atomic.AddInt64(&this.numOpen, 1)
}

func (this *MySQLPool) DecrNumOpen() {
	atomic.AddInt64(&this.numOpen, -1)
}

// 关闭指定链接 // 序号在获取 mutex lock 使用, 不然会出现死锁
func (this *MySQLPool) closeConn(db *sql.DB) error {
	err := db.Close()
	this.DecrNumOpen()

	return err
}

// 获取链接
func (this *MySQLPool) Get() (*sql.DB, error) {
	// 先从chan中获取资源
	select {
	case db, ok := <-this.dbChan:
		if ok {
			return db, nil
		}
	default:
	}

	this.Lock()
	// 等待获取资源
	if this.NumOpen() >= this.maxOpen {
		this.Unlock()
		db := <-this.dbChan
		return db, nil
	}

	// 新键资源
	this.IncrNumOpen() // 添加已经使用资源
	// 新键链接
	db, err := sql.Open("mysql", this.Cfg.DSN())
	if err != nil {
		this.Unlock()
		this.DecrNumOpen() // 链接没有成功删除已经使用资源
		return nil, fmt.Errorf("链接数据库出错: %s", err.Error())
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	if err := db.Ping(); err != nil {
		this.Unlock()
		this.DecrNumOpen() // 链接没有成功删除已经使用资源
		return nil, fmt.Errorf("ping数据库出错: %s", err.Error())
	}

	this.Unlock()

	return db, nil
}

// 归还链接
func (this *MySQLPool) Release(db *sql.DB) error {
	this.Lock()

	if this.NumOpen() > this.maxOpen { // 关闭资源
		if err := this.closeConn(db); err != nil {
			logger.M.Error("归还链接, 当前打开的链接超过最大允许打开的阈值(%d). 关闭链接出错. %s", this.maxOpen, err.Error())
		}
		logger.M.Info("归还链接, 当前打开的链接超过最大允许打开的阈值(%d).", this.maxOpen)
		this.Unlock()
		return nil
	}

	this.Unlock()

	this.dbChan <- db
	return nil
}

// 获取允许最大打开数
func (this *MySQLPool) MaxOpen() int64 {
	return this.maxOpen
}

// 获取允许最小打开数
func (this *MySQLPool) MinOpen() int64 {
	return this.minOpen
}

// 当前已经打开的数量
func (this *MySQLPool) NumOpen() int64 {
	return atomic.LoadInt64(&this.numOpen)
}

// 设置最大允许的链接数
func (this *MySQLPool) SetMaxOpen(maxOpen int64) error {
	if maxOpen > MYSQL_POOL_MAX_OPEN_LIMIT {
		return fmt.Errorf("设置最大允许链接数:%d, 超过了系统限制:%d", maxOpen, MYSQL_POOL_MAX_OPEN_LIMIT)
	}

	atomic.StoreInt64(&this.maxOpen, maxOpen)
	return nil
}
