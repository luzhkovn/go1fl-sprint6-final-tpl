package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/luzhkovn/go1fl-sprint6-final-tpl/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	convertedStr, err := service.AutoConvert(string(data))
	if err != nil {
		http.Error(w, "Ошибка конвертации файла", http.StatusInternalServerError)
		return
	}
	filename := filepath.Ext(header.Filename)
	result := time.Now().UTC().String() + filename

	newfile, err := os.Create(result)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer newfile.Close()
	newfile.WriteString(convertedStr)
	w.Write([]byte(convertedStr))
}
