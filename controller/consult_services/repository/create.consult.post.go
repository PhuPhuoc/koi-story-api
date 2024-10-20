package consultrepository

import (
	"fmt"

	consultmodel "github.com/PhuPhuoc/koi-story-api/controller/consult_services/model"
	"github.com/PhuPhuoc/koi-story-api/utils"
	"github.com/google/uuid"
)

func (store *consultStore) CreateNewConsult(data consultmodel.CreateConsultInfo) error {
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

	data.Post.ID = uuid.New().String()
	data.Post.CreatedAt = utils.CreateDateTimeCurrentFormated()
	query_post := `
    insert into post (id, user_id, post_type, title, created_at)
    values (:id, :user_id, :post_type, :title, :created_at)
    `
	_, err = tx.NamedExec(query_post, data.Post)
	if err != nil {
		return fmt.Errorf("failed to insert new consult (post): %w", err)
	}

	data.Consult.ID = uuid.New().String()
	data.Consult.PostID = data.Post.ID
	query_consutl := `
    insert into detail_consult (id, post_id, content)
    values (:id, :post_id, :content)
    `
	_, err = tx.NamedExec(query_consutl, data.Consult)
	if err != nil {
		return fmt.Errorf("failed to insert new consult (post): %w", err)
	}

	img_id := uuid.New().String()
	query_image := `
	insert into post_image (id, post_id, file_path, image_order)
	values (?, ?, ?, ?)
	`
	_, err = tx.Exec(query_image, img_id, data.Post.ID, data.FilePath, 0)
	if err != nil {
		return fmt.Errorf("failed to insert new market (post_image): %w", err)
	}

	return tx.Commit()
}
