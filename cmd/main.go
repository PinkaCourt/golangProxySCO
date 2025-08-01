package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
	wsClient *websocket.Conn
)

func main() {
	// Запускаем WebSocket сервер
	go startWebSocketServer()

	// Запускаем HTTP сервер
	http.HandleFunc("/message/text", handleHTTPMessage)
	log.Println("HTTP сервер запущен на порту 8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("Ошибка при запуске HTTP сервера: %v", err)
	}
}

func startWebSocketServer() {
	http.HandleFunc("/websocket", wsHandler)
	log.Println("WebSocket сервер запущен на порту 6696")
	if err := http.ListenAndServe(":6696", nil); err != nil {
		log.Fatalf("Ошибка при запуске WebSocket сервера: %v", err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	wsClient, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Ошибка при Upgrade: ", err)
		return
	}
	defer wsClient.Close()

	log.Println("WebSocket соединение открыто")

	// Слушаем сообщения от клиента
	for {
		_, msg, err := wsClient.ReadMessage()
		if err != nil {
			log.Println("Ошибка при чтении сообщения: ", err)
			break
		}
		log.Printf("Получено сообщение от клиента: %s", msg)
	}
}

func handleHTTPMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var message string
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Ошибка при декодировании данных", http.StatusBadRequest)
		return
	}

	message = r.FormValue("data")

	if wsClient != nil {
		if err := wsClient.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			http.Error(w, "Ошибка при отправке сообщения по WebSocket", http.StatusInternalServerError)
			return
		}
		log.Printf("Отправлено сообщение по WebSocket: %s", message)
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Данные успешно получены")
}
