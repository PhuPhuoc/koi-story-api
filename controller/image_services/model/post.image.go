package imagemodel

type PostImage struct {
	ID         string `db:"id" json:"id"`
	PostID     string `db:"post_id" json:"-"`
	FilePath   string `db:"file_path" json:"file_path"`
	ImageOrder int    `db:"image_order" json:"image_order"`
	DeletedAt  string `db:"deleted_at" json:"-"`
}
