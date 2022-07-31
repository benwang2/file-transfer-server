package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"src/keygen"
	"src/models"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func setupTables() {
	tx, _ := db.Begin()
	_, err := tx.Exec(
		"CREATE TABLE IF NOT EXISTS `resources` (" +
			"`resourcePath` VARCHAR(2083) NOT NULL," +
			"`resourceType` INT NOT NULL," +
			"`fileName` VARCHAR(260) NULL," +
			"`url` VARCHAR(64) NOT NULL," +
			"`created` DATETIME NOT NULL," +
			"`track` TINYINT(1) NOT NULL DEFAULT 0," +
			"`size` INT NULL," +
			"`expires` DATETIME NULL," +
			"`accessKey` VARCHAR(255) NULL," +
			"UNIQUE INDEX `url_UNIQUE` (url ASC) VISIBLE," +
			"PRIMARY KEY (url)" +
			") 	ENGINE = InnoDB")
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}

func SaveToDatabase(w http.ResponseWriter, resource *models.Resource) *models.Response {
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	fmt.Printf("Saving to database;\n")

	var resp models.Response
	resource.URL = strings.TrimSpace(resource.URL)

	if resource.URL == "" {
		resource.URL = generateUniqueURL(resource.URLType)
		resp.StatusCode = 200
	} else if !isURLUnique(resource.URL) {
		resp.StatusCode = 400
		http.Error(w, "URL provided already in use.", http.StatusBadRequest)
		return nil
	}

	fmt.Printf("database.go url=%s\n", resource.URL)

	today := time.Now()
	created := today.Format(time.RFC3339)
	expires := today.AddDate(0, 0, 0).Format(time.RFC3339)

	shouldTrack := 0
	if resource.Track {
		shouldTrack = 1
	}

	castedType := 0
	if resource.Type == "file" {
		castedType = 1
	} else if resource.Type == "link" {
		castedType = 2
	}

	var query string
	switch castedType {
	case 1:
		query = "INSERT INTO resources ( `fileName`, `resourcePath`, `url`, `resourceType`, `created`, `track`, `size`, `expires`, `accessKey` ) " +
			`VALUES ( "%s", "%s", "%s", %d, "%s", %d, %d, "%s", "%s" )`

		query = fmt.Sprintf(query, resource.Header.Filename, resource.Path, resource.URL, castedType, string(created), shouldTrack, int64(resource.Header.Size), string(expires), resource.Password)
		break
	case 2:
		query = "INSERT INTO resources ( `fileName`, `resourcePath`, `url`, `resourceType`, `created`, `track`, `size`, `expires`, `accessKey` ) " +
			`VALUES ( "", "%s", "%s", "%d", "%s", "%d", "0", "%s", "%s" )`

		query = fmt.Sprintf(query, resource.Destination, resource.URL, castedType, string(created), shouldTrack, string(expires), resource.Password)
		break
	}

	if _, err := tx.Exec(query); err != nil {
		log.Fatal(err)
		return nil
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

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

		if url == "api" {
			url = ""
		}
	}

	return url
}

func GetResourceByKey(w http.ResponseWriter, key string) (*models.DBResource, error) {
	row := db.QueryRow(
		"SELECT resourcePath, resourceType, fileName, created, track, size, expires, accessKey " +
			"FROM resources " +
			"WHERE `url`=\"" + key + "\"")
	var resource models.DBResource

	err := row.Scan(
		&resource.Path,
		&resource.Type,
		&resource.FileName,
		&resource.Created,
		&resource.Track,
		&resource.Size,
		&resource.Expires,
		&resource.AccessKey)
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
	fmt.Println("database url=" + databaseURL)
	db, _ = sql.Open("mysql", databaseURL)
	setupTables()

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
