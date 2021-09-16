package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"petcard/internal/migrations"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	connection, err := gorm.Open(mysql.Open("root:root@/petcard"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	err = migrations.DatabaseConfig(connection)
	if err != nil {
		panic("could not run databaseConfig...")
	}

	return DB
}
