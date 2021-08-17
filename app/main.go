package main

import (
	routes "jomblo/app/routes"
	_UserUsecase "jomblo/business/users"
	_UserController "jomblo/controllers/users"
	_UserRepository "jomblo/drivers/databases/users"
	"jomblo/drivers/mysql"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_UserRepository.Users{},
	)
}

func main() {
	configDB := mysql.ConfigDB{
		DB_Username: "root",
		DB_Password: "Gaktauaku123!",
		DB_Host:     "localhost",
		DB_Port:     "3306",
		DB_Database: "jomblo",
	}

	db := configDB.InitialDB()
	dbMigrate(db)

	e := echo.New()
	UserRepository := _UserRepository.NewMysqlUserRepository(db)
	UserUsecase := _UserUsecase.NewUserUseCase(UserRepository)
	UserController := _UserController.NewUserController(UserUsecase)

	routeInit := routes.ControllerList{
		UserController: *UserController,
	}

	routeInit.RouteRegister(e)
	e.Start(
		":3000",
	)
}
