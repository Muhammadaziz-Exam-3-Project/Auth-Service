package main

import (
	"Github.com/LocalEats/Api-get-way/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con1, err := grpc.NewClient(":8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	con2, err := grpc.NewClient(":8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	router := api.RouterApi(con1, con2)
	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
