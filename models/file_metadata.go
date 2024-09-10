package models

import "time"

type FileMetadata struct {
    ID        int       `json:"id"`
    Filename  string    `json:"filename"`
    Size      int64     `json:"size"`
    UploadAt  time.Time `json:"upload_at"`
}
