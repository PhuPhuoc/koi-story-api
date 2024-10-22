package imagerepository

import (
	imagemodel "github.com/PhuPhuoc/koi-story-api/controller/image_services/model"
	"github.com/google/uuid"
)

func (store *imageStore) AddNewImageInPost(postID string, url imagemodel.UpdateImage) error {
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

	var exists bool
	err = tx.Get(&exists, "select exists(select 1 from post_image where post_id = ?)", postID)
	if err != nil {
		return err
	}

	imageOrder := 0
	if exists {
		err = tx.Get(&imageOrder, "select max(image_order) from post_image where post_id = ?", postID)
		if err != nil {
			return err
		}
		imageOrder++
	}
	img_id := uuid.New().String()
	_, err = tx.Exec(`
		insert into post_image (id, post_id, file_path, image_order)
		values (?, ?, ?, ?)
	`, img_id, postID, url.FilePath, imageOrder)
	if err != nil {
		return err
	}

	return tx.Commit()
}
