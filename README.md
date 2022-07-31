# File Transfer Server
A micro-service providing file transfer, storage, and redirects - best suited for individual use.

## About The Project
This project was made using Golang and [gorilla/mux](https://github.com/gorilla/mux). 

The service stores files on the server's file system and records the path to this file in a MySQL database. Redirects are also stored in the database. To prevent duplicate file uploads, files are identified by their hashes before being committed to the file system.

## Getting Started
This project requires an installation of Golang, [gorilla/mux](https://github.com/gorilla/mux), and a MySQL database.

The MySQL database can be configured using the query in [database/database.go](https://github.com/benwang2/file-transfer-server/blob/main/src/database/database.go).

To run the project, use the command `go run main.go` while in the `src` directory.

## Planned Features
- Improved error handling for end user
- Basic authentication to ensure only admin can use service
- Add tracking and password protection to uploaded resources