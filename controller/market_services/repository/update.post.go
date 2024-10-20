package marketrepository

import (
	"fmt"

	marketmodel "github.com/PhuPhuoc/koi-story-api/controller/market_services/model"
)

func (store *marketStore) UpdateMarketDetails(post_id string, post marketmodel.UpdatePostMarket) error {

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

	post.PostUpdate.ID = post_id
	query_post := `
    update post set post_type=:post_type, title=:title
    where id=:id
    `
	_, err = tx.NamedExec(query_post, post.PostUpdate)
	if err != nil {
		return fmt.Errorf("failed to update market (post): %w", err)
	}

	post.DetailMarket.PostID = post_id
	query_detail := `
    update detail_market set product_name=:product_name, product_type=:product_type, price=:price,
    seller_address=:seller_address, phone_number=:phone_number, description=:description
    where post_id=:post_id
    `
	_, err = tx.NamedExec(query_detail, post.DetailMarket)
	if err != nil {
		return fmt.Errorf("failed to update market (detail_market): %w", err)
	}

	// Lấy detail_market_id dựa trên post_id
	var detailMarketID string
	query_get_detail_id := `
	 select id from detail_market where post_id = ?
	 `
	err = tx.Get(&detailMarketID, query_get_detail_id, post_id)
	if err != nil {
		return fmt.Errorf("failed to get detail_market_id: %w", err)
	}

	post.KoiInfo.DetailMarketID = detailMarketID
	query_info := `
    update koi_info set color=:color, size=:size, old=:old, type=:type
    where detail_market_id=:detail_market_id
    `
	_, err = tx.NamedExec(query_info, post.KoiInfo)
	if err != nil {
		return fmt.Errorf("failed to update market (koi_info): %w", err)
	}

	return tx.Commit()
}
