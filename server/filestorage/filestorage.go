package filestorage

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"server/database"
	"server/models"
)

func File(w http.ResponseWriter, r models.Resource, f multipart.File, h *multipart.FileHeader) (models.Response, error) {
	defer f.Close()
	var resp *models.Response
	var buf bytes.Buffer
	io.Copy(&buf, f)

	path := "./store/" + h.Filename
	r.Path = path
	r.File = f
	r.Header = h

	fmt.Printf("filestorage.go path=%s;\n", path)
	resp = database.SaveFileToDatabase(w, &r)

	if resp == nil {
		return models.Response{StatusCode: 400, Message: "Resource already exists."}, nil
	}

	out, _ := os.Create(path)
	out.Write([]byte(buf.String()))
	out.Close()

	resp.Message = "Status OK"
	resp.StatusCode = 200
	resp.Data = r.URL
	return *resp, nil
}

func URL(r models.Resource) (models.Response, error) {
	fmt.Println("Processing url")
	var resp models.Response

	resp.Message = "Hello"
	return resp, nil
}
