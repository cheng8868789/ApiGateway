package main

import (
	"fmt"
	"github.com/cheng8868789/ApiGateway/apiConfig"
	"github.com/cheng8868789/ApiGateway/apiHandler"
	"net/http"
)

func main() {

	err := Init()
	if err != nil {
		return
	}

	http.HandleFunc("/apiTest/", apiHandler.ServeHandler)

	//启动http server
	addr := apiConfig.GetAddr()
	http.ListenAndServe(addr, nil)

}

func Init() error {

	err := apiConfig.Init()
	if err != nil {
		fmt.Println("apiConfig.Init() failure")
		return err
	}

	fmt.Println("ApiGateway init success......")
	return nil
}
