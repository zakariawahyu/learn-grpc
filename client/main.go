package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"learn-grpc/common/config"
	"learn-grpc/model"
	"log"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "n001",
		Name:     "Zakaria Wahyu",
		Password: "hl12/3m,a",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	user2 := model.User{
		Id:       "n002",
		Name:     "Nur Utomo",
		Password: "sj/.za",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	garage1 := model.Garage{
		Id:   "q001",
		Name: "Quel'thalas",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	garage2 := model.Garage{
		Id:   "q002",
		Name: "Amsterdam",
		Coordinate: &model.GarageCoordinate{
			Latitude:  77.123123123,
			Longitude: 99.1231313123,
		},
	}

	user := serviceUser()
	fmt.Println("\n", "===========> user test")
	// register user1
	user.Register(context.Background(), &user1)
	// register user2
	user.Register(context.Background(), &user2)

	// show all user register
	resUser, err := user.List(context.Background(), new(model.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}

	resUserString, _ := json.Marshal(resUser.List)
	log.Println(string(resUserString))

	garage := serviceGarage()
	fmt.Println("\n", "===========> garage test")
	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})
	// add garage2 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage2,
	})
	// add garage2 to user2
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user2.Id,
		Garage: &garage2,
	})

	// show all garages of user1
	resGarage, err := garage.List(context.Background(), &model.GarageUserId{UserId: user2.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	resStringGarage, _ := json.Marshal(resGarage.List)
	log.Println(string(resStringGarage))
}
