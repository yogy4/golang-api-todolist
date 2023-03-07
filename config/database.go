package config

import (
	"fmt"
	"golang-api-todolist/entities"
	"log"

	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConn() (*gorm.DB, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error reading config file %s", err)
	}

	v_user := viper.Get("MYSQL_USER")
	v_pass := viper.Get("MYSQL_PASSWORD")
	v_host := viper.Get("MYSQL_HOST")
	v_db := viper.Get("MYSQL_DBNAME")
	v_port := viper.Get("MYSQL_PORT")

	mysql_conn := fmt.Sprintf("%v:%v@(%v:%v)/%v?parseTime=true", v_user, v_pass, v_host, v_port, v_db)
	fmt.Println("connection \t\t", mysql_conn)

	// db, err := gorm.Open("mysql", mysql_conn)
	db, err := gorm.Open(mysql.Open(mysql_conn), &gorm.Config{})
	if err != nil {
		panic("failed to connect")
	}

	DB = db
	DB.Debug().AutoMigrate(&entities.Activity{}, &entities.Todo{})
	return db, nil
}
