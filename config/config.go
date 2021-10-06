package config

import (
	"fmt"
	"log"
	"pasarwarga/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var (
		Article   = models.Article{}
		Category  = models.Category{}
		Users     = models.Users{}
		Otps      = models.Otps{}
		Company   = models.Company{}
		Location  = models.Locations{}
		Position  = models.Position{}
		Candidate = models.Candidate{}
	//	PositionCategory = models.PositionCategory{} dihilangi dulu ( bisa cmn gatau error knp saat di up kke 2)
	// perhatikan collationnya kalau beda bisa error
		Filepdf = models.Filepdf{}
	)

	dsn := "root:@tcp(127.0.0.1:3306)/pasarwarga?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false}) //set true kalau mau dimatiin

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&Article, &Category, &Users, &Otps, &Company, &Location, &Position, &Candidate, &Filepdf)

	fmt.Println("Connecting To Database...")

	DB = db

}
