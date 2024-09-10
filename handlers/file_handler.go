package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"106HW/db"
	"106HW/models"

	"github.com/gorilla/mux"
)

const MAX_UPLOAD_SIZE = 10 * 1024 * 1024

var allowedFileTypes = map[string]bool{
	".txt": true,
	".jpg": true,
	".png": true,
	".pdf": true,
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "Fayl hajmi juda katta.", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Faylni yuklashda xatolik.", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	ext := filepath.Ext(handler.Filename)
	if !allowedFileTypes[ext] {
		http.Error(w, "Ruxsat etilmagan fayl turi.", http.StatusBadRequest)
		return
	}

	uploadPath := "./uploads/" + handler.Filename
	dst, err := os.Create(uploadPath)
	if err != nil {
		http.Error(w, "Faylni saqlashda xatolik.", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Faylni yozishda xatolik.", http.StatusInternalServerError)
		return
	}
	fileMetadata := models.FileMetadata{
		Filename: handler.Filename,
		Size:     handler.Size,
		UploadAt: time.Now(),
	}
	db.SaveFileMetadata(fileMetadata)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Fayl muvaffaqiyatli yuklandi.",
		"filename": handler.Filename,
		"filepath": uploadPath,
	})
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	filePath := "./uploads/" + filename
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "Fayl topilmadi.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, filePath)
}

func GetFiles(w http.ResponseWriter, r *http.Request) {
	files := db.GetFileMetadata()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}
