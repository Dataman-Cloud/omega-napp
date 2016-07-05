package store

import (
	"fmt"
	"github.com/Dataman-Cloud/omega-napp/config"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
	"github.com/mattes/migrate/migrate"
	"sync"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db != nil {
		return db
	}
	mutex := sync.Mutex{}
	mutex.Lock()
	InitDB()
	defer mutex.Unlock()

	return db
}

func InitDB() {
	var err error
	conf := config.Pairs()
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		conf.Db.User,
		conf.Db.Password,
		conf.Db.Host,
		conf.Db.Port,
		conf.Db.Name)
	log.Infof("mysql connection uri: %s", uri)
	db, err = gorm.Open("mysql", uri)
	if err != nil {
		log.Fatalf("init mysql error: %v", err)
	}
	db.DB().SetMaxIdleConns(conf.Db.MaxIdleConns)
	db.DB().SetMaxOpenConns(conf.Db.MaxOpenConns)
}

func UpgradeDB() {
	conf := config.Pairs()
	driver := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s",
		conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port, conf.Db.Name)
	log.Info("upgrading DB", driver)
	errors, ok := migrate.UpSync(driver, "./sql")
	log.Debug(ok)
	log.Debug(errors)
	if errors != nil && len(errors) > 0 {
		for _, err := range errors {
			log.Error("db err ", err)
		}
		log.Fatalf("can't upgrade db %v", errors)
	}
	if !ok {
		log.Fatal("can't upgrade db")
	}
	log.Info("DB upgraded")
}
