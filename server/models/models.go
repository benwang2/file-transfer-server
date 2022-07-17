package models

import "mime/multipart"

type FileType int64

const (
	File FileType = iota
	Redirect
)

type Resource struct {
	Type       string `json:"type"`
	URL        string `json:"url,omitempty"`
	URLType    string `json:"urlType,omitempty"`
	Name       string `json:"name,omitempty"`
	Expiration string `json:"expiration,omitempty"`
	Password   string `json:"password,omitempty"`
	Track      bool   `json:"track,omitempty"`

	// Files
	Path   string                `json:"-"`
	File   multipart.File        `json:"-"`
	Header *multipart.FileHeader `json:"-"`

	// URL
	Destination string `json:"-"`
}

type DBResource struct {
	FileName  string
	Path      string
	Type      int
	URL       string
	Created   string
	Track     string
	Size      int
	Expires   string
	AccessKey string
}

type Response struct {
	Message    string
	StatusCode int
	Data       string
}

type Secrets struct {
	DatabaseURL string `json:"databaseUrl"`
}
