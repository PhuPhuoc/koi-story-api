package marketrepository

import (
	"fmt"

	marketmodel "github.com/PhuPhuoc/koi-story-api/controller/market_services/model"
)

func (store *marketStore) GetMarketDetailsByID(postID string) (*marketmodel.MarketDetails, error) {
	var md marketmodel.MarketDetails
	md.ID = postID

	tx, err := store.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	query_post := `
	select id, user_id, post_type, title, created_at
	from post
	where id = ?
	`
	if err := tx.Get(&md.Post, query_post, postID); err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	query_detail := `
	select post_id, product_name, product_type, price, seller_address, phone_number, description
	from detail_market
	where post_id = ?
	`
	if err := tx.Get(&md.DetailMarket, query_detail, postID); err != nil {
		return nil, fmt.Errorf("failed to get detail: %w", err)
	}
	fmt.Println("md.DetailMarket.ID: ", md.DetailMarket.ID)
	query_koi := `
	select k.detail_market_id, k.color, k.size, k.old, k.type
	from koi_info k
	join detail_market dm on dm.id = k.detail_market_id
	join post p on p.id = dm.post_id
	where p.id = ?
	`
	if err := tx.Get(&md.KoiInfo, query_koi, postID); err != nil {
		return nil, fmt.Errorf("failed to get koi: %w", err)
	}

	query_image := `
        SELECT id, post_id, file_path, image_order
        FROM post_image
        WHERE post_id = ?
        ORDER BY image_order ASC
    `
	if err := tx.Select(&md.ListImage, query_image, postID); err != nil {
		return nil, fmt.Errorf("failed to get images: %w", err)
	}

	return &md, nil
}
