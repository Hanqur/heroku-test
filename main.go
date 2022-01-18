package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //анализ аргументов,
    fmt.Fprintf(w, "Hello!") // отправляем данные на клиентскую сторону
}

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")

    http.HandleFunc("/", HomeRouterHandler) // установим роутер
    err := http.ListenAndServe(":" + port, nil) // задаем слушать порт
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}