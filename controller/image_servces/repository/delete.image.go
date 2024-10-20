package imagerepository

import "github.com/PhuPhuoc/koi-story-api/utils"

func (store *imageStore) DeleteImagePost(image_id string) error {
	query_post := `
    update post_image set deleted_at=?
    where id=?
    `
	_, err := store.db.Exec(query_post, utils.CreateDateTimeCurrentFormated(), image_id)
	if err != nil {
		return err
	}
	return nil
}
