package http

import (
	"HomeAI/auth"
	"HomeAI/wol"
	"fmt"
	"net/http"
)

const HeaderAuthToken string = "X-AuthToken"

// 启动HTTP服务
func ServerRun(listenAddr string) error {

	http.HandleFunc("/wol", func(response http.ResponseWriter, request *http.Request) {
		userToken := request.Header.Get(HeaderAuthToken)
		err := auth.Auth(userToken)
		if err != nil {
			forbidden(response)
			return
		}
		request.ParseForm()
		mac := request.Form.Get("mac")
		err = wol.Wol(mac)
		if err != nil {
			response.Write([]byte("error\n"))
			response.Write([]byte(err.Error()))
		} else {
			response.Write([]byte("ok"))
		}
	})

	fmt.Printf("listening at %v\n", listenAddr)
	return http.ListenAndServe(listenAddr, nil)
}

func forbidden(response http.ResponseWriter) error {
	response.WriteHeader(403)
	_, err := response.Write([]byte("Forbidden"))
	return err
}
