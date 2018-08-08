package mysql

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type MysqlOptions struct {
	Hostname    string
	Port        string
	User        string
	Password    string
	DbName      string
	TablePrefix string
}

// MysqlService implements mysql service
type MysqlService struct {
	options MysqlOptions
	DB      *sql.DB
}

// NewMysqlService create new mysql service
func NewMysqlService(opts MysqlOptions) *MysqlService {
	service := &MysqlService{}
	url := opts.User + ":" + opts.Password + "@tcp(" + opts.Hostname + ":" + opts.Port + ")/" + opts.DbName + "?charset=utf8&parseTime=True"
	db, err := sql.Open("mysql", url)
	if err != nil {
		// todo : add log
		return nil
	}
	service.DB = db
	service.options = opts
	return service
}
