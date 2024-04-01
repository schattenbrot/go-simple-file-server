package files

type File struct {
	ID        int    `json:"id"`
	FileName  string `json:"fileName"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}
