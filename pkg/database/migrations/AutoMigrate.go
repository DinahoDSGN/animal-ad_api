package migrations

import (
	"gorm.io/gorm"
	"log"
	"petcard/pkg/models"
)

func DatabaseConfig(connection *gorm.DB) error {
	//Drop(connection)

	Migrate(connection)

	//SQLQueries(connection)

	return nil
}

func Migrate(connection *gorm.DB) error {
	err := connection.AutoMigrate(&models.Ad{}, &models.Animal{}, &models.User{}, &models.Breed{}, &models.AdLocation{})
	if err != nil {
		return err
	}

	log.Println("Migrations finished")

	return nil
}

func Drop(connection *gorm.DB) error {
	err := connection.Migrator().DropTable("users")
	if err != nil {
		return err
	}
	err = connection.Migrator().DropTable("ads")
	if err != nil {
		return err
	}
	err = connection.Migrator().DropTable("specifies")
	if err != nil {
		return err
	}
	err = connection.Migrator().DropTable("breeds")
	if err != nil {
		return err
	}
	return nil
}

//func SQLQueries(connection *gorm.DB) {
//	user := telegramModels.User{
//		Name:     "Niet",
//		Lastname: "Nazhimedenov",
//		Username: "dinahosl",
//		Email:    "dinahodsgn@gmail.com",
//		Password: "123",
//	}
//
//	connection.Create(&user)
//
//	user = telegramModels.User{
//		Name:     "Angelina",
//		Lastname: "Fast",
//		Username: "fastik",
//		Email:    "faaast@gmail.com",
//		Password: "456",
//	}
//
//	connection.Create(&user)
//
//	user = telegramModels.User{
//		Name:     "Abylai",
//		Lastname: "Abdreym",
//		Username: "abo_sya",
//		Email:    "abo@gmail.com",
//		Password: "789",
//	}
//
//	connection.Create(&user)
//
//	spec := telegramModels.Animal{
//		Name:       "Snezhka",
//		Breed:      "Unknown",
//		Color:      "White",
//		Gender:     false,
//		Vaccinated: false,
//		Spayed:     false,
//		Passport:   false,
//	}
//
//	connection.Create(&spec)
//
//	spec = telegramModels.Animal{
//		Name:       "Simba",
//		Breed:      "Unknown",
//		Color:      "White",
//		Gender:     true,
//		Vaccinated: false,
//		Spayed:     false,
//		Passport:   false,
//	}
//
//	connection.Create(&spec)
//
//	spec = telegramModels.Animal{
//		Name:       "Asya",
//		Breed:      "Human",
//		Color:      "Black",
//		Gender:     false,
//		Vaccinated: false,
//		Spayed:     true,
//		Passport:   true,
//	}
//
//	connection.Create(&spec)
//
//	ad := telegramModels.Ad{
//		Title:       "My little Snezhka",
//		AdLocation:    "Esik City",
//		Description: "at it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now ",
//		SpecifyId:   1,
//		UserId:      1,
//	}
//
//	connection.Create(&ad)
//
//	ad = telegramModels.Ad{
//		Title:       "My little Simba",
//		AdLocation:    "Esik City",
//		Description: "at it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now ",
//		SpecifyId:   2,
//		UserId:      2,
//	}
//
//	connection.Create(&ad)
//
//	ad = telegramModels.Ad{
//		Title:       "My little Asya",
//		AdLocation:    "Esik City",
//		Description: "at it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now ",
//		SpecifyId:   3,
//		UserId:      3,
//	}
//
//	connection.Create(&ad)
//}
