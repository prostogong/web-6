package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var count1 int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(count1)))
		return
	} else if r.Method == "POST" {
		r.ParseForm()
		s := r.FormValue("count")
		if s == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))

			return
		}
		number, err := strconv.Atoi(s)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))

			return
		}
		count1 += number
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Метод не поддерживается"))
		return
	}

}

func main() {
	http.HandleFunc("/count", handler)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера!")
	}
}
