package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"learn-grpc/model"
	"os"
	"strings"
)

func main() {
	user1 := &model.User{
		Id:       "u001",
		Name:     "Johns Tom",
		Password: "f0rTheP4s",
		Gender:   model.UserGender_FEMALE,
	}

	userList := &model.UserList{
		List: []*model.User{
			user1,
		},
	}

	garage1 := &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.23423,
			Longitude: 24.23412,
		},
	}

	garageList := &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	garageListByUser := &model.GarageListBuyer{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	// Print Proto Object
	// pada print dibawah ini, object akan tercetak apa adanya. Generate struct memiliki property lain
	// ==== original
	fmt.Printf("# ==== Original\n       %#v \n", user1)
	fmt.Printf("# ==== Original\n       %#v \n", garageListByUser)
	// ==== as string, bisa juga langsung diubah ke dalam bentuk string
	fmt.Printf("# ==== As String\n       %v \n", user1.String())
	fmt.Printf("# ==== As String\n       %v \n", garageListByUser.String())

	// Print Proto Object ke JSON String
	var buf bytes.Buffer
	if err := (&jsonpb.Marshaler{}).Marshal(&buf, garageList); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	jsonString := buf.String()
	fmt.Printf("# ==== As JSON String\n       %v \n", jsonString)

	// Print JSON String ke Proto Object
	buf2 := strings.NewReader(jsonString)
	protoObject := new(model.GarageList)
	if err := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As JSON String to Proto\n       %v \n", protoObject.String())

	protoObject2 := new(model.GarageList)
	if err := jsonpb.UnmarshalString(jsonString, protoObject2); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As JSON String to Proto\n       %v \n", protoObject2.String())

	// Bisa juga menggunakan package proto
	userBytes, err := proto.Marshal(userList)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("# ==== As Bytes Data\n       %v \n", userBytes)

	listUser := &model.UserList{}
	if err := proto.Unmarshal(userBytes, listUser); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("# ==== As JSON\n       %v \n", listUser)
}
