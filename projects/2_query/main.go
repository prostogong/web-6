package main

import (
	"fmt"
	"net/http" // пакет для поддержки HTTP протокола
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("name")
	w.Write([]byte("Hello, " + s + "!"))
}

// и тут тоже (если очень надо)
func main() {
	http.HandleFunc("/api/user", handler)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
