package marketrepository

import "github.com/PhuPhuoc/koi-story-api/utils"

func (store *marketStore) DeleteMarketPost(post_id string) error {
	query_post := `
    update post set deleted_at=?
    where id=?
    `
	_, err := store.db.Exec(query_post, utils.CreateDateTimeCurrentFormated(), post_id)
	if err != nil {
		return err
	}
	return nil
}
