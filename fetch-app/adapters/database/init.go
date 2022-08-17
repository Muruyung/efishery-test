package database

import (
	"MyAPI/config"

	"xorm.io/xorm"
)

// Database is wrapper for xorm engine
type Database struct {
	*xorm.Engine
}

// Init create new database engine
func Init() (*Database, error) {
	db, err := xorm.NewEngine(config.DBENGINE(), config.DBCONFIG())
	if err != nil {
		return nil, err
	}

	return &Database{
		db,
	}, nil
}
