package main

import(
    "net/http"
    "html/template"
    "log"
    "path"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

func setRoutes() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/ws", wsHandler)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fp := path.Join("templates", "index.html")
    tmpl, err := template.ParseFiles(fp)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func reader(conn *websocket.Conn) {
    for {
        messageType, p, err := conn.ReadMessage()

        if err != nil {
            log.Println(err)
            return
        }

        log.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }
    }
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool {
        return true
    }

    ws, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        log.Println(err)
    }

    reader(ws)

}

func main() {
    setRoutes()
    log.Fatal(http.ListenAndServe(":8081", nil))
}
