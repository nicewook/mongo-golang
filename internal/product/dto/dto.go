package dto

type Release struct {
	Version int    `json:"version,omitempty"`
	Date    string `json:"date,omitempty"`
}

type Review struct {
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	Comment   string `json:"comment,omitempty"`
	CreatedAt string `json:"created_at,omitempty"` // RFC3339
}

type Product struct {
	Name    string   `json:"name,omitempty"`
	Type    string   `json:"type,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Release `json:"release,omitempty"`
	Reviews []Review `json:"reviews,omitempty"`
}
