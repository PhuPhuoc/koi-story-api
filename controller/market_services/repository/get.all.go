package marketrepository

import marketmodel "github.com/PhuPhuoc/koi-story-api/controller/market_services/model"

func (store *marketStore) GetAllPost() ([]marketmodel.PostMarket, error) {
	var posts []marketmodel.PostMarket

	query := `
        SELECT p.id as post_id, dm.product_name, dm.price, pi.file_path
        FROM post p
        JOIN detail_market dm ON p.id = dm.post_id
        JOIN post_image pi ON p.id = pi.post_id where pi.image_order=0
    `

	err := store.db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
