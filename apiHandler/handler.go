package apiHandler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ServeHandler(rw http.ResponseWriter, req *http.Request){

	msg ,err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.Write([]byte(err.Error()))
	}
	if len(msg) != 0 {
		fmt.Println(msg)
	}



}