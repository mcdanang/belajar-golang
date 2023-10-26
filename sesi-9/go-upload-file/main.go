package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Couldn't get file from form data", http.StatusBadRequest)
			return
		}
		defer file.Close()

		uploadDir := "./uploads"
		os.MkdirAll(uploadDir, os.ModePerm)
		filePath := filepath.Join(uploadDir, handler.Filename)

		// open new file to store data from uploaded file
		newFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Couldn't create file", http.StatusBadRequest)
			return
		}
		defer newFile.Close()

		// copy data from uploaded file to new file
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, "Couldn't copy file data", http.StatusBadRequest)
			return
		}

		fmt.Println(w, "File uploaded succesfully")

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server started at :1010")
	http.ListenAndServe(":1010", nil)
}
