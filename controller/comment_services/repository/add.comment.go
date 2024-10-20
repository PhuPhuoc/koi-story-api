package commmentrepository

import (
	commentmodel "github.com/PhuPhuoc/koi-story-api/controller/comment_services/model"
	"github.com/PhuPhuoc/koi-story-api/utils"
	"github.com/google/uuid"
)

func (store *commentStore) AddCommentToPost(post_id string, com commentmodel.CreateCommentModel) error {
	comment := commentmodel.Comment{
		ID:        uuid.New().String(),
		UserID:    com.UserID,
		PostID:    post_id,
		Content:   com.Content,
		CreatedAt: utils.CreateDateTimeCurrentFormated(),
	}
	_, err := store.db.NamedExec(`
		insert into comment (id, user_id, post_id, content, created_at)
		values (:id, :user_id, :post_id, :content, :created_at)
	`, comment)
	if err != nil {
		return err
	}
	return nil
}
