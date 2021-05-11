package db

import (
	"backend/pkg/tools"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type DataBase struct {
	Storage *gorm.DB
}

func New(config *MySQLConfig) *DataBase {
	return &DataBase{Storage: open(config)}
}

func open(c *MySQLConfig) *gorm.DB {
	DNS := tools.JoinStrings(
		c.User, ":",
		c.Password, "@tcp(",
		c.Host, ":",
		c.Port, ")/",
		c.DbName, "?charset=utf8&parseTime=true&loc=Local")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		DSN:                       DNS,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NoLowerCase:   false,
		},
		Logger:                 logger.Default.LogMode(logger.Warn),
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		PrepareStmt:                              true,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
	})
	if err != nil {
		log.Println("Please Check MySQL Connection...")
		os.Exit(-2)
	}
	log.Println()
	log.Printf("\033[1;32;32m MySQL RUNNING [ %s:%s ] \033[0m", c.Host, c.Port)
	log.Println()
	return db
}

func SetUp(db *gorm.DB, models ...interface{}) {
	database, _ := db.DB()
	database.SetMaxIdleConns(10)
	database.SetMaxOpenConns(128)
	if err := db.AutoMigrate(models...); err != nil {
		log.Println("Create Table Error: ", err.Error())
	}
}

func Close(db *gorm.DB) {
	database, _ := db.DB()
	database.Close()
}
