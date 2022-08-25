package dto

import (
	"bytes"
	"encoding/gob"
)

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

type ErrorResp struct {
	Code    string
	Message string
}

// Utility functions
func DeepCopy(src, dist interface{}) (err error) {
	buf := bytes.Buffer{}
	if err = gob.NewEncoder(&buf).Encode(src); err != nil {
		return
	}
	return gob.NewDecoder(&buf).Decode(dist)
}
