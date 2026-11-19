package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "ошибка при чтении index.html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при загрузке файла - myFile", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "ошибка при конвертации содержимого файла", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(fileHeader.Filename)
	timestamp := time.Now().Format("20060102-150405.000000000")
	name := "converted" + timestamp + ext

	err = os.WriteFile(name, []byte(result), 0644)
	if err != nil {
		http.Error(w, "ошибка при сохранении результата", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result))
}
