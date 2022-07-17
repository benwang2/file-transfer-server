package filestorage

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"server/database"
	"server/models"
	"strings"
)

func md5sum(file multipart.File) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func File(w http.ResponseWriter, r models.Resource, f multipart.File, h *multipart.FileHeader) (models.Response, error) {

	var resp *models.Response
	var buf bytes.Buffer
	io.Copy(&buf, f)
	hash, _ := md5sum(f)

	defer f.Close()

	pathSplit := strings.Split(h.Filename, ".")
	path := "./store/" + hash + "." + pathSplit[len(pathSplit)-1]
	r.Path = path
	r.File = f
	r.Header = h

	fmt.Printf("filestorage.go path=%s;\n", path)
	resp = database.SaveFileToDatabase(w, &r)

	if resp == nil {
		return models.Response{StatusCode: 400, Message: "Resource already exists."}, nil
	}

	resp.Message = "Status OK"
	resp.StatusCode = 200
	resp.Data = r.URL

	if _, err := os.Stat(path); err == nil {
		fmt.Println("filestorage.go File already exists - expected behavior.")
	} else if errors.Is(err, os.ErrNotExist) {
		out, _ := os.Create(path)
		out.Write([]byte(buf.String()))
		out.Close()
	} else {
		resp.Message = "Unknown error occurred."
		resp.StatusCode = 500
	}

	return *resp, nil
}

func URL(r models.Resource) (models.Response, error) {
	fmt.Println("Processing url")
	var resp models.Response

	resp.Message = "Hello"
	return resp, nil
}
