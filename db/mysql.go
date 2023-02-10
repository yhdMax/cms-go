package db

import (
	"github.com/fatih/color"
	"github.com/jmoiron/sqlx"
	"time"
	"vue-next-admin-go/modles"

	_ "github.com/go-sql-driver/mysql"
)

// DB 声明全局变量
var (
	DB *sqlx.DB
)

func InitMysql(config modles.Configuration) {
	newDb(config.DB.URL, config.DB.MaxOpenConns, config.DB.MaxIdleConns)
}

func newDb(url string, maxOpenConns, maxIdleConns int) {
	DB = sqlx.MustConnect("mysql", url)
	// 设置最大的连接数 默认是无限制 如果超出限制了 就会排队等待
	DB.SetMaxOpenConns(maxOpenConns)
	// 设置最大的空闲连接数 默认是无限制 业务量小的时候 可以把多余的连接释放掉，只保留一定数量的连接数
	DB.SetMaxIdleConns(maxIdleConns)

	// 检查mysql是否连接成功
	ping(DB)
}

func ping(db *sqlx.DB) {
	// 格式化 print
	colorPrint := color.New()
	colorPrint.Add(color.Italic)
	colorPrint.Add(color.FgGreen)
	colorPrint.Add(color.Bold)

	// 测试mysql是否连通
	if err := db.Ping(); err != nil {

		colorPrint.Add(color.BgRed)
		colorPrint.Add(color.FgWhite)

		// 出错时关闭数据库，防止新的数据操作开始
		err = db.Close()

		// time.Now().UnixNano() / 1e6	当前程序执行时间戳
		// time.Now().Format("2006-01-02 15:04:05")	当前日期格式化显示	format参数为格式化模板

		_, err := colorPrint.Println("Init Mysql Server Error:", err.Error(), time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			return
		}
		return
	}

	_, err := colorPrint.Println("Init Mysql Server Successfully:", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return
	}
}
