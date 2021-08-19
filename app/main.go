package main

import (
	routes "jomblo/app/routes"

	_MatchesUsecase "jomblo/business/matches"
	_MatchesController "jomblo/controllers/matches"
	_MatchesRepository "jomblo/drivers/databases/matches"

	_UserUsecase "jomblo/business/users"
	_UserController "jomblo/controllers/users"
	_UserRepository "jomblo/drivers/databases/users"

	_ChannelsUsecase "jomblo/business/channels"
	_ChannelsController "jomblo/controllers/channels"
	_ChannelsRepository "jomblo/drivers/databases/channels"

	_ChatsUsecase "jomblo/business/chats"
	_ChatsController "jomblo/controllers/chats"
	_ChatsRepository "jomblo/drivers/databases/chats"

	"jomblo/drivers/mysql"
	"log"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_UserRepository.Users{},
		&_MatchesRepository.Matches{},
		&_ChannelsRepository.Channels{},
		&_ChatsRepository.Chats{},
	)
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
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

	MatchesRepository := _MatchesRepository.NewMysqlMatchesRepository(db)
	MatchesUsecase := _MatchesUsecase.UserMatchesUsecase(MatchesRepository)
	MatchesController := _MatchesController.NewMatchesController(MatchesUsecase)

	ChannelsRepository := _ChannelsRepository.NewMysqlChannelsRepository(db)
	ChannelsUsecase := _ChannelsUsecase.UserChannelsUsecase(ChannelsRepository)
	ChannelsController := _ChannelsController.NewChannelsController(ChannelsUsecase)

	ChatsRepository := _ChatsRepository.NewMysqlChatsRepository(db)
	ChatsUsecase := _ChatsUsecase.UserChatsUsecase(ChatsRepository)
	ChatsController := _ChatsController.NewUserController(ChatsUsecase)

	routeInit := routes.ControllerList{
		UserController:     *UserController,
		MatchesController:  *MatchesController,
		ChannelsController: *ChannelsController,
		ChatsController:    *ChatsController,
	}

	routeInit.RouteRegister(e)
	e.Start(
		":3000",
	)
}
