package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"server/database"
	"server/filestorage"
	"server/keygen"
	"server/models"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func RandKey(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, keygen.Chars(3)+"\n"+keygen.Word())
}

func Access(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resource, err := database.GetResourceByKey(w, vars["key"])
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Found no resource by key=%s\n", vars["key"])
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		} else {
			panic(err)
		}
	} else {
		// 1 : file
		// 2 : redirect
		if resource.Type == 1 {
			w.Header().Set("Content-Disposition", "attachment; filename="+resource.FileName)
			w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

			http.ServeFile(w, r, resource.Path)
		} else if resource.Type == 2 {

		}
	}
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

	var resp models.Response

	if resource.Type == "url" {
		resp, err = filestorage.URL(resource)
	} else if resource.Type == "file" {
		file, header, err := r.FormFile("file")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("middleware.go url=%s;\n", resource.URL)
		resp, err = filestorage.File(w, resource, file, header)

		if err != nil {
			panic(err)
		}

	} else {
		http.Error(w, "server: invalid file type", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
