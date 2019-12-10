package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const (
	grafanaUser = "admin" //Пользователь под которым мы будем авторизовываться
	grafanaHost = "grafana:3000" //Адрес расположения grafana
	grafanaHeader = "X-GRAFANA-AUTH" //Header наличие которого определяет успешную авторизацию
	)

func SendJSONError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, `{"error":{"msg":%q}}`, msg)
}


func handlerProxy(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Host)
	if strings.HasPrefix(r.URL.String(), "/api") {
		//Проверка прав в сервисе авторизации
	}

	url, err := url.Parse(fmt.Sprintf("http://%s/", grafanaHost))
	if err != nil {
		SendJSONError(w, err.Error())
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(url)

	fmt.Println(r.URL.Host)
	r.Header.Set(grafanaHeader, grafanaUser)

	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handlerProxy)
	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}