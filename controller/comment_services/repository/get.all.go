package commmentrepository

import commentmodel "github.com/PhuPhuoc/koi-story-api/controller/comment_services/model"

func (store *commentStore) GetAllCommentInPost(post_id string) ([]commentmodel.PostComment, error) {
	query := `
	select c.id, c.user_id, c.content, c.created_at,
	u.display_name as username, u.profile_picture_url as user_avatar_url
	from comment c join user u on c.user_id=u.id
	where c.deleted_at is null and c.post_id=? order by c.created_at desc 
	`
	var comments []commentmodel.PostComment
	if err := store.db.Select(&comments, query, post_id); err != nil {
		return nil, err
	}
	return comments, nil
}
