package module

type Contents struct {
	Title    string    `json:"title,omitempty"`
	Contents []Content `json:"contents,omitempty"`
}

type Content struct {
	ContentType string                 `json:"content_type,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
}
