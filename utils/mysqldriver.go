package utils

import (
	"cv_app/configs"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)
	fmt.Println(connectionString)
	// "root:@tcp(127.0.0.1:3306)/be5db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	// db.Migrator().DropTable(&entities.Ability{})
	// db.Migrator().DropTable(&entities.Achievement{})
	// db.Migrator().DropTable(&entities.Education{})
	// db.Migrator().DropTable(&entities.Language{})
	// db.Migrator().DropTable(&entities.WorkExperience{})
	// db.Migrator().DropTable(&entities.User{})

	// db.AutoMigrate(&entities.User{})
	// db.AutoMigrate(&entities.Ability{})
	// db.AutoMigrate(&entities.Achievement{})
	// db.AutoMigrate(&entities.Education{})
	// db.AutoMigrate(&entities.Language{})
	// db.AutoMigrate(&entities.WorkExperience{})

	// var userUid []string

	// for i := 0; i < 50; i++ {

	// 	userUid := shortuuid.New()
	// 	password, _ := middlewares.HashPassword("xyz")

	// 	db.Create(&entities.User{
	// 		UserUid:  userUid,
	// 		Name:     faker.Name(),
	// 		Email:    faker.Email(),
	// 		Password: password,
	// 		Address:  "jl.dramaga no.22",
	// 		Gender:   "female",
	// 		About:    "testingtesting",
	// 	})

	// }

}
