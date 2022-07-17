package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/keygen"
	"server/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func setupTables() {
	db.Exec(
		`CREATE TABLE IF NOT EXISTS resources (
			resourcePath VARCHAR(2083) NOT NULL,
            resourceType INT NOT NULL,
			url VARCHAR(2000) NOT NULL,
			created DATETIME NOT NULL,
			track TINYINT(1) NOT NULL DEFAULT 0,
			size INT NOT NULL,
			expires DATETIME NULL,
			accessKey VARCHAR(255) NULL,
			UNIQUE INDEX resourcePath_UNIQUE (resourcePath ASC) VISIBLE,
			UNIQUE INDEX url_UNIQUE (url ASC) VISIBLE,
			PRIMARY KEY (url)
		) 	ENGINE = InnoDB
	`)
}

func SaveFileToDatabase(w http.ResponseWriter, resource *models.Resource) *models.Response {
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	fmt.Printf("Saving to database;\n")

	var resp models.Response
	if resource.URL == "" {
		resource.URL = generateUniqueURL(resource.URLType)
		resp.StatusCode = 200
	} else if !isURLUnique(resource.URL) {
		resp.StatusCode = 400
		http.Error(w, "URL provided already in use.", http.StatusBadRequest)
		return nil
	}

	today := time.Now()
	created := today.Format(time.RFC3339)
	expires := today.AddDate(0, 0, 0).Format(time.RFC3339)

	shouldTrack := 0
	if resource.Track {
		shouldTrack = 1
	}

	castedType := 0
	// fmt.Printf("Resource type = %s\n", resource.Type)
	if resource.Type == "file" {
		castedType = 1
	}

	var query string = "INSERT INTO resources ( `resourcePath`, `url`, `resourceType`, `created`, `track`, `size`, `expires`, `accessKey` ) " +
		`VALUES ( "%s", "%s", %d, "%s", %d, %d, "%s", "%s" )`

	query = fmt.Sprintf(query, resource.Path, resource.URL, castedType, string(created), shouldTrack, int64(resource.Header.Size), string(expires), resource.Password)
	if _, err := tx.Exec(query); err != nil {
		return nil
	}

	err = tx.Commit()

	return &resp
}

func isURLUnique(url string) bool {
	if url == "" {
		return false
	}

	row := db.QueryRow("SELECT url FROM resources WHERE url=\"" + url + "\"")
	var res string
	err := row.Scan(&res)

	if err == sql.ErrNoRows {
		return true
	} else if err != nil {
		log.Fatal(err)
	}

	return false
}

func generateUniqueURL(urlType string) string {
	var url string
	for !isURLUnique(url) {
		if urlType == "word" {
			url = keygen.Word()
		} else {
			url = keygen.Chars(3)
		}
	}

	return url
}

func GetResourceByKey(w http.ResponseWriter, key string) (*models.DBResource, error) {
	row := db.QueryRow(
		"SELECT resourcePath, resourceType, created, track, size, expires, accessKey " +
			"FROM resources " +
			"WHERE `url`=\"" + key + "\"")
	var resource models.DBResource

	err := row.Scan(&resource.Path, &resource.Type, &resource.Created, &resource.Track, &resource.Size, &resource.Expires, &resource.AccessKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			panic(err)
		}
	}

	return &resource, nil
}

func Start() {
	fmt.Println("Connecting to database")
	databaseURL := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
	fmt.Println(databaseURL)
	db, _ = sql.Open("mysql", databaseURL)

	// if err != nil {
	// 	panic(err)
	// }

	if err := db.Ping(); err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func Stop() {
	db.Close()
}
