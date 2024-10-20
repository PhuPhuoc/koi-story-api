package consultmodel

type PostConsult struct {
	ID         string `db:"id" json:"id"`
	UserID     string `db:"user_id" json:"user_id"`
	UserName   string `db:"user_name" json:"user_name"`
	UserAvatar string `db:"user_avatar" json:"user_avatar"`
	PostType   string `db:"post_type" json:"post_type"`
	Title      string `db:"title" json:"title"`
	Content    string `db:"content" json:"content"`
	FilePath   string `db:"file_path" json:"file_path"`
	CreatedAt  string `db:"created_at" json:"created_at"`
}

type CreateConsultInfo struct {
	Post
	Consult
	FilePath string `db:"file_path" json:"file_path"`
}

type UpdateConsultInfo struct {
	PostUpdate
	Consult
	FilePath string `db:"file_path" json:"file_path"`
}

type PostUpdate struct {
	ID       string `db:"id" json:"-"`
	PostType string `db:"post_type" json:"post_type"`
	Title    string `db:"title" json:"title"`
}

type Post struct {
	ID        string `db:"id" json:"-"`
	UserID    string `db:"user_id" json:"user_id"`
	PostType  string `db:"post_type" json:"post_type"`
	Title     string `db:"title" json:"title"`
	CreatedAt string `db:"created_at" json:"-"`
}

type Consult struct {
	ID      string `db:"id" json:"-"`
	PostID  string `db:"post_id" json:"-"`
	Content string `db:"content" json:"content"`
}
