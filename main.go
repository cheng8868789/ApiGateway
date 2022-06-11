package main

import (
	"fmt"
	"github.com/cheng8868789/ApiGateway/apiConfig"
	"github.com/cheng8868789/ApiGateway/apiHandler"
	"net/http"
)

func main() {

	Init()

	http.HandleFunc("/apiTest",apiHandler.ServeHandler)

	//启动http server
	http.ListenAndServe("0.0.0.0:7080",nil)

}


func Init()  {

	apiConfig.Init()

	fmt.Println("ApiGateway init success......")

}
