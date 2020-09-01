package main

import (
	"log"
	"net/http"
)

//301 永久重定向
func Redirect301Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.baidu.com", http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/baidu1", Redirect301Handler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))
}
