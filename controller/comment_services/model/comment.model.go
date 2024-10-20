package commentmodel

type Comment struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	PostID    string `db:"post_id"`
	Content   string `db:"content"`
	CreatedAt string `db:"created_at"`
}

type PostComment struct {
	ID            string `db:"id" json:"id"`
	UserID        string `db:"user_id" json:"-"`
	UserName      string `db:"username" json:"username"`
	UserAvatarUrl string `db:"user_avatar_url" json:"user_avatar_url"`
	Content       string `db:"content" json:"content"`
	CreatedAt     string `db:"created_at" json:"created_at"`
}

type CreateCommentModel struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
