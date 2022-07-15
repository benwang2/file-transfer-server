package models

type Resource struct {
	Type       string   `json:"type"`
	Resource   string   `json:"resource,omitempty"`
	URL        string   `json:"url,omitempty"`
	Name       string   `json:"name,omitempty"`
	Expiration string   `json:"expiration,omitempty"`
	IPs        []string `json:"IPs,omitempty"`
	Password   string   `json:"password,omitempty"`
	Track      bool     `json:"track,omitempty"`
}

type UploadResponse struct {
	Message string
}
