package filestorage

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"server/models"
)

func SaveFileToStorage(r *models.Resource, f multipart.File, h *multipart.FileHeader) (bool, error) {

	return true, nil
}

func File(r *models.Resource, f multipart.File, h *multipart.FileHeader) (models.UploadResponse, error) {
	defer f.Close()
	var resp models.UploadResponse
	var buf bytes.Buffer
	io.Copy(&buf, f)

	out, _ := os.Create(h.Filename)
	out.Write([]byte(buf.String()))
	out.Close()

	resp.Message = "Hello"
	return resp, nil
}

func URL(r *models.Resource) (models.UploadResponse, error) {
	fmt.Println("Processing url")
	var resp models.UploadResponse

	resp.Message = "Hello"
	return resp, nil
}
