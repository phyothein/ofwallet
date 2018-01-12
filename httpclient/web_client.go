package httpclient

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
)

const (
	POST = 1
	GET  = 2
)

type Response struct {
	StatusCode int
	Body       string
}

type HttpRequest struct {
	request  *http.Request
	Response chan Response
}

func NewHttpRequest(url string, method int, rpc string) *HttpRequest {


	httpRequest  := new(HttpRequest)

    fmt.Println(rpc)
	if method == POST {
		request, err:= http.NewRequest("POST", url, strings.NewReader(rpc))
		if err != nil {
			return nil
		}
		httpRequest.request = request

	} else if method == GET {
		request, err := http.NewRequest("GET", url, strings.NewReader(rpc))
		if err != nil {
			return nil
		}
		httpRequest.request = request
	}

	httpRequest.request.Header.Set("Content-Type", "application/json")

	httpRequest.Response = make(chan Response)

	return httpRequest
}

func (hs *HttpRequest) Do() {

	go func() {

		respon :=new(Response)
		resp, err := http.DefaultClient.Do(hs.request)

		if err!=nil||resp==nil{
			respon.StatusCode=404
			respon.Body = ""
			hs.Response<-*respon
			return
		}

		defer func() {
			if resp!=nil{
				resp.Body.Close()
			}
		}()

		if resp.StatusCode==200{
			respon.StatusCode=resp.StatusCode
			b, _ := ioutil.ReadAll(resp.Body)
			respon.Body = string(b)
			hs.Response<-*respon
		}
	}()
}
