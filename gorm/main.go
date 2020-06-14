package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:redli9600,.0@@(112.74.36.53)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.WithFields(log.Fields{
			"db": "127.0.0.1",
		}).Info("connect mysql error")
		return
	}
	defer db.Close()

	//if !db.HasTable(&Like{}) {
	//	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
	//		panic(err)
	//	}
	//}
	//
	//ip := "127.0.0.1"
	//ua := "ua"
	//title := "title"
	//
	//like := &Like{
	//	Ip:        ip,
	//	Ua:        ua,
	//	Title:     title,
	//	Hash:      murmur3.Sum64([]byte(strings.Join([]string{ip, ua, title}, "-"))) >> 1,
	//	CreatedAt: time.Now(),
	//}
	//
	//if err := db.Create(like).Error; err != nil {
	//	return
	//}
	var count int
	err = db.Model(&Like{}).Where(&Like{Ip: "127.0.0.1"}).Count(&count).Error
	if err != nil {
		return
	}
	log.WithFields(log.Fields{"data count": count}).Info("query data")

	like := Like{}
	err = db.First(&like).Error
	log.Info(like)

	for range time.Tick(time.Millisecond * 100) {
		log.Info(time.Now().Unix())
	}
}
