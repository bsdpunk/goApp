package crimes

import (
	"fmt"
	"reflect"

	crimeModels "github.com/bsdpunk/goApp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func con() *gorm.DB {
	dsn := "host=192.168.1.94 user=dusty password=Qapla1999 dbname=dusty port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db
}

func init() {
	fmt.Println("vim-go")
	DB = con()
}

func LoadDetroit() {
	detroit, detroits := crimeModels.GetDetroit()
	//detroits := crimeModels.GetDetroits()
	fmt.Println(reflect.TypeOf(detroit))
	fmt.Println(reflect.TypeOf(DB))
	//	fmt.Println(*detroit{})
	DB.AutoMigrate(&detroit)
	DB.CreateInBatches(detroits, 100)
}

func GetDB() *gorm.DB {
	return DB

}
