package apiHandler

import (
	"encoding/json"
	"fmt"
	"github.com/cheng8868789/ApiGateway/apiConfig"
	"github.com/cheng8868789/Common/models"
	"io/ioutil"
	"net/http"
	"strings"
)

func ServeHandler(rw http.ResponseWriter, req *http.Request) {

	urlPath := req.URL.Path[8:]

	config := apiConfig.GetConfig()
	urls := config.UrlData.Urls
	for _, url := range urls {
		if strings.EqualFold(url.Url, urlPath) {
			fmt.Println("get serviceName")
		} else {
			continue
		}
	}

	msg, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.Write([]byte(err.Error()))
	}

	requestModel := new(models.ApiModel)
	err = json.Unmarshal(msg, requestModel)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}
	fmt.Println(requestModel)

}
