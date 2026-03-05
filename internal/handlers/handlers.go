package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"Morse-Code-Converter/internal/service"
)

// IndexHandler return form is file index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler it processes the file upload, converts the data, and saves the result.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	result, err := service.AutoConvert(string(data))
	if err != nil {
		http.Error(w, fmt.Sprintf("filed to process file: %v", err), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	newFileName := time.Now().UTC().String() + ext

	dst, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "Failed to create local file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = dst.WriteString(result)
	if err != nil {
		http.Error(w, "Failed to write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
