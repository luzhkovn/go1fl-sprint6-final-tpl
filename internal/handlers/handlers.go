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

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	var formKey string
	if r.MultipartForm != nil && len(r.MultipartForm.File) > 0 {
		for key := range r.MultipartForm.File {
			formKey = key
			break
		}
	} else {
		http.Error(w, "Файл не найден", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile(formKey)
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
	timeStr := time.Now().UTC().Format("20060102150405")
	result := timeStr + filepath.Ext(header.Filename)

	newfile, err := os.Create(result)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer newfile.Close()
	newfile.WriteString(convertedStr)
	w.Write([]byte(convertedStr))
}
