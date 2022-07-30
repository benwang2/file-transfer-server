package storage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"server/database"
	"server/models"
	"strings"
)

func checksum(buf bytes.Buffer) (string, error) {
	hasher := sha256.New()
	if _, err := io.Copy(hasher, &buf); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func File(w http.ResponseWriter, r models.Resource, f multipart.File, h *multipart.FileHeader) (models.Response, error) {

	var resp *models.Response
	var buf bytes.Buffer

	defer f.Close()
	io.Copy(&buf, f)

	hash, _ := checksum(buf)

	pathSplit := strings.Split(h.Filename, ".")
	path := "./store/" + hash + "." + pathSplit[len(pathSplit)-1]
	r.Path = path
	r.File = f
	r.Header = h

	fmt.Printf("storage.go path=%s;\n", path)
	resp = database.SaveToDatabase(w, &r)

	if resp == nil {
		return models.Response{StatusCode: 400, Message: "Resource already exists."}, nil
	}

	resp.Message = "Status OK"
	resp.StatusCode = 200
	resp.Data = r.URL

	fmt.Printf("File uploaded with checksum:%s\n", hash)
	if _, err := os.Stat(path); err == nil {
		fmt.Println("storage.go File already exists - expected behavior.")
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

func URL(w http.ResponseWriter, r models.Resource) (models.Response, error) {
	fmt.Printf("link to \"%s\"\n", r.Destination)
	var resp *models.Response

	resp = database.SaveToDatabase(w, &r)

	resp.Message = "Status OK"
	resp.StatusCode = 200
	resp.Data = r.URL
	return *resp, nil
}
