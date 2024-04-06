package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    // 读取 HTML 文件
    tmpl := template.Must(template.ParseFiles("main.html"))
    // 渲染 HTML 模板
    tmpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    // 获取搜索框中的输入内容
    query := r.FormValue("query")
    fmt.Println("搜索关键词:", query)
    // 这里可以根据搜索关键词进行相关操作，比如调用 Google 搜索 API，然后展示相关内容
    // 在这个示例中，我们直接输出搜索框输入的内容
    fmt.Fprintf(w, "搜索框中的内容为：%s", query)
}

func main() {
    // 设置路由
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/search", searchHandler)

    // 启动服务器
    fmt.Println("服务器启动在 http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
