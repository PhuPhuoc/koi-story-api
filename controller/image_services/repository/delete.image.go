package imagerepository

import "github.com/PhuPhuoc/koi-story-api/utils"

func (store *imageStore) DeleteImagePost(image_id string) error {
	tx, err := store.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	query_post := `
    update post_image set post_id=?, deleted_at=?
    where id=?
    `
	_, err = tx.Exec(query_post, "", utils.CreateDateTimeCurrentFormated(), image_id)
	if err != nil {
		return err
	}

	return nil
}
