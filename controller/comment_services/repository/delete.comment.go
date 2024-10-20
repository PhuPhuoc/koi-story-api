package commmentrepository

import "github.com/PhuPhuoc/koi-story-api/utils"

func (store *commentStore) DeleteComment(comment_id string) error {
	query_post := `
    update comment set deleted_at=?
    where id=?
    `
	_, err := store.db.Exec(query_post, utils.CreateDateTimeCurrentFormated(), comment_id)
	if err != nil {
		return err
	}
	return nil
}
