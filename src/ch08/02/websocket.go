package _2

import (
	"net/http"
	"log"
	"github.com/golang/net/websocket"
	"fmt"
)

func Echo(ws *websocket.Conn){
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't recieve")
			break
		}
		fmt.Println("Recieved back from client: " + reply)
		msg := "Received: " + reply
		fmt.Println("Sending to client: " + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main()  {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil);err != nil {
		log.Fatal("ListenAndServer:", err)
	}

}
