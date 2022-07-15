package middleware

import (
	"encoding/json"
	"net/http"
	"server/filestorage"
	"server/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm((1024 * 8) << 20)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var resource models.Resource
	s := string(r.FormValue("data"))
	err = json.Unmarshal([]byte(s), &resource)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var resp models.UploadResponse

	if resource.Type == "url" {
		resp, err = filestorage.URL(&resource)
	} else if resource.Type == "file" {
		file, header, err := r.FormFile("file")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err = filestorage.File(&resource, file, header)
	} else {
		http.Error(w, "server: invalid file type", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
