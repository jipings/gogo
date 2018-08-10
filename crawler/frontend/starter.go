package main

import (
	"net/http"

	"imooc.com/tutorial/crawler/frontend/controller"
)

func main() {
	// 开启FileServer之后服务器会自动去找index.html文件
	// TODO: 重新定义template.html文件和index.html
	http.Handle("/", http.FileServer(http.Dir("view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		"view/index.html",
	))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
