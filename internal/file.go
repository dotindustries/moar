package internal

type File struct {
	Name     string `json:"name,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	Data     []byte `json:"data,omitempty"`
	Uri      string `json:"uri,omitempty"`
}
