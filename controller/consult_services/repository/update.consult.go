package consultrepository

import (
	"fmt"

	consultmodel "github.com/PhuPhuoc/koi-story-api/controller/consult_services/model"
)

func (store *consultStore) UpdateConsult(post_id string, data consultmodel.UpdateConsultInfo) error {
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

	data.PostUpdate.ID = post_id
	query_post := `
    update post set post_type=:post_type, title=:title
    where id=:id
    `
	_, err = tx.NamedExec(query_post, data.PostUpdate)
	if err != nil {
		return fmt.Errorf("failed to insert new consult (post): %w", err)
	}

	data.Consult.ID = post_id
	query_consutl := `
    update detail_consult set content=:content
    where id=:id
    `
	_, err = tx.NamedExec(query_consutl, data.Consult)
	if err != nil {
		return fmt.Errorf("failed to insert new consult (consult): %w", err)
	}

	query_image := `
    update post_image set file_path=?
    where post_id=?
    `
	_, err = tx.Exec(query_image, data.FilePath, post_id)
	if err != nil {
		return fmt.Errorf("failed to insert new consult (image): %w", err)
	}
	return tx.Commit()
}
