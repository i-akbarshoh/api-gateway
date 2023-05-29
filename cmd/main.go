package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/i-akbarshoh/api-gateway/internal/controller/user"
	"github.com/i-akbarshoh/api-gateway/internal/pkg/config"
	"github.com/i-akbarshoh/api-gateway/internal/pkg/proto"
	"github.com/i-akbarshoh/api-gateway/internal/pkg/router"
	user1 "github.com/i-akbarshoh/api-gateway/internal/usecase/user"
	"google.golang.org/grpc"
)

func main() {
	c := config.C.AuthClient
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.Host, c.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewAuthClient(conn)
	uc := user1.NewUsecase(&client)
	controller := user.NewController(uc)
	r := router.New(controller)

	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		if err := http.ListenAndServe("localhost:8080", r); err != nil {
			log.Println(err)
		}
	}()
	fmt.Println("Program initialized...")
	wg.Wait()
}