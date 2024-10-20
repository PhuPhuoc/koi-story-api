package marketrepository

import marketmodel "github.com/PhuPhuoc/koi-story-api/controller/market_services/model"

func (store *marketStore) GetAllPost() ([]marketmodel.PostMarket, error) {
	var posts []marketmodel.PostMarket

	query := `
        select p.id as post_id, dm.product_name, dm.price, coalesce(pi.file_path, '') as file_path
        from post p
        join detail_market dm ON p.id = dm.post_id
        LEFT JOIN
		(
			select post_id, file_path
			from post_image
			where deleted_at IS NULL
			order by image_order ASC
			limit 1
		)
		pi ON p.id = pi.post_id
		where p.deleted_at is null
		order by p.created_at desc
    `

	err := store.db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
