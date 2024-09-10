package main

import (
    "log"
    "net/http"
    "os"

    "106HW/handlers"
    "106HW/db"
    
    "github.com/gorilla/mux"
)

func main() {
    os.MkdirAll("./uploads", os.ModePerm)

    db.InitDB()

    r := mux.NewRouter()

    r.HandleFunc("/upload", handlers.UploadFile).Methods("POST")
    r.HandleFunc("/download/{filename}", handlers.DownloadFile).Methods("GET")
    r.HandleFunc("/files", handlers.GetFiles).Methods("GET")

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
