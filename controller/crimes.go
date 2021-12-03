package crimes

import (
	"fmt"
	"net/http"
	"reflect"

	crimeModels "github.com/bsdpunk/goApp/models"
	"github.com/gin-gonic/gin"
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

//CreateUser ... Create Detroit
func CreateDetroit(c *gin.Context) {
	var detroit crimeModels.Detroit
	c.BindJSON(&detroit)
	err := crimeModels.CreateDetroit(&detroit)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, detroit)
	}
}

func LoadDetroit() {
	detroit, detroits := crimeModels.GetDetroit()
	//detroits := crimeModels.GetDetroits()
	fmt.Println(reflect.TypeOf(detroit))
	fmt.Println(reflect.TypeOf(DB))
	//	fmt.Println(*detroit{})
	DB.AutoMigrate(&detroit)
	DB.Migrator().CreateConstraint(detroit, "fk_id")
	DB.Migrator().HasConstraint(detroit, "fk_id")
	DB.Migrator().CreateIndex(detroit, "id")
	DB.CreateInBatches(detroits, 100)
}

//func LoadOffense() {
//detroit, detroits := crimeModels.GetDetroit()
//detroits := crimeModels.GetDetroits()
//fmt.Println(reflect.TypeOf(detroit))
//fmt.Println(reflect.TypeOf(DB))
//	fmt.Println(*detroit{})

//DB.Migrator().CreateConstraint(detroit, "fk_id")
//DB.Migrator().HasConstraint(detroit, "fk_id")
//DB.Migrator().CreateIndex(detroit, "id")
//DB.CreateInBatches(detroits, 100)

//}

func GetDB() *gorm.DB {
	return DB

}

//func (c *gin.Context) {
//    var detroit crimeModels.Detroit
//    c.BindJSON(&todo)
// DB call to create a todo
// Config.DB.Create(todo).Error;
//    err := Models.CreateADetroit(&detroit)
//    if err != nil {
//	c.AbortWithStatus(http.StatusNotFound)
//    } else {
//	c.JSON(http.StatusOK, todo)
//    }
//}
