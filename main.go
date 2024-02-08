package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "главная страница")
}

func login_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Функция, реализующая метод /login")
}

func verify_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Функция, реализующая метод /verify")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/login/", login_page)
	http.HandleFunc("/verify/", verify_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	go func() {
		handleRequest()
	}()

	// Ждем сигналов SIGTERM или Interrupt
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	// Получен сигнал, останавливаем сервер
	fmt.Println("Получен сигнал для остановки сервера...")

	// Создаем HTTP-сервер
	server := &http.Server{
		Addr: ":8080",
	}

	err := server.Shutdown(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Сервер остановлен.")
}
