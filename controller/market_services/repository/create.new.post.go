package marketrepository

import (
	"fmt"

	marketmodel "github.com/PhuPhuoc/koi-story-api/controller/market_services/model"
	"github.com/PhuPhuoc/koi-story-api/utils"
	"github.com/google/uuid"
)

func (store *marketStore) CreateNewMarketPost(post marketmodel.NewPostMarket) error {

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

	post.Post.ID = uuid.New().String()
	post.Post.CreatedAt = utils.CreateDateTimeCurrentFormated()
	query_post := `
    insert into post (id, user_id, post_type, title, created_at)
    values (:id, :user_id, :post_type, :title, :created_at)
    `
	_, err = tx.NamedExec(query_post, post.Post)
	if err != nil {
		return fmt.Errorf("failed to insert new market (post): %w", err)
	}

	post.DetailMarket.ID = uuid.New().String()
	post.DetailMarket.PostID = post.Post.ID
	post.DetailMarket.CreatedAt = utils.CreateDateTimeCurrentFormated()
	query_detail := `
    insert into detail_market (id, post_id, product_name, product_type, price, seller_address, phone_number, description, created_at)
    values (:id, :post_id, :product_name, :product_type, :price, :seller_address, :phone_number, :description, :created_at)
    `
	_, err = tx.NamedExec(query_detail, post.DetailMarket)
	if err != nil {
		return fmt.Errorf("failed to insert new market (detail): %w", err)
	}

	post.KoiInfo.ID = uuid.New().String()
	post.KoiInfo.DetailMarketID = post.DetailMarket.ID
	post.KoiInfo.CreatedAt = utils.CreateDateTimeCurrentFormated()
	query_info := `
    insert into koi_info (id, detail_market_id, color, size, old, type, created_at)
    values (:id, :detail_market_id, :color, :size, :old, :type, :created_at)
    `
	_, err = tx.NamedExec(query_info, post.KoiInfo)
	if err != nil {
		return fmt.Errorf("failed to insert new market (koi_info): %w", err)
	}

	query_image := `
	insert into post_image (id, post_id, file_path, image_order)
	values (?, ?, ?, ?)
	`
	for i, url := range post.ListImage {
		img_id := uuid.New().String()
		_, err := tx.Exec(query_image, img_id, post.Post.ID, url, i)
		if err != nil {
			return fmt.Errorf("failed to insert new market (post_image): %w", err)
		}
	}

	return tx.Commit()
}
