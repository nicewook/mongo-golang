package entity

type Release struct {
	Version int
	Date    string
}

type Review struct {
	Name      string
	Type      string
	UserID    string
	Comment   string
	CreatedAt string // RFC3339
}

type Product struct {
	Name string
	Type string
	Tags []string
	Release
	Reviews []Review
}
