package mysqldao

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MysqlConifg implements config
type MysqlConifg struct {
	Hostname    string
	Port        string
	User        string
	Password    string
	DbName      string
	TablePrefix string
	Debug       bool
}

// RdsService
type RdsService struct {
	config MysqlConifg
	DB     *gorm.DB
}

// NewRdsService create new mysql dao service
func NewRdsService(config MysqlConifg) (*RdsService, error) {
	impl := &RdsService{}
	impl.config = config

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.TablePrefix + defaultTableName
	}

	url := config.User + ":" + config.Password + "@tcp(" + config.Hostname + ":" + config.Port + ")/" + config.DbName + "?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Fatalf("mysql connection error:%s", err.Error())
		return nil, err
	}

	db.LogMode(config.Debug)

	impl.DB = db

	return impl, nil
}

// AddTable add
func (s *RdsService) AddTable(t interface{}) {
	if ok := s.DB.HasTable(t); !ok {
		if err := s.DB.CreateTable(t).Error; err != nil {
			log.Fatalf("create mysql table error:%s", err.Error())
		}
	}
	var tab []interface{}
	s.DB.AutoMigrate(append(tab, t))
}

// AddTables create tables for given object
func (s *RdsService) AddTables(tables []interface{}) {
	for _, t := range tables {
		if ok := s.DB.HasTable(t); !ok {
			if err := s.DB.CreateTable(t).Error; err != nil {
				log.Fatalf("create mysql table error:%s", err.Error())
			}
		}
	}

	// auto migrate to keep schema update to date
	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON'T change existing column's type or delete unused columns to protect your data
	s.DB.AutoMigrate(tables...)
}

// Close db
func (s *RdsService) Close() error {
	return s.DB.Close()
}

// Add single item
func (s *RdsService) Add(item interface{}) error {
	return s.DB.Create(item).Error
}

// Del single item
func (s *RdsService) Del(item interface{}) error {
	return s.DB.Delete(item).Error
}

// Save single item
func (s *RdsService) Save(item interface{}) error {
	return s.DB.Save(item).Error
}
