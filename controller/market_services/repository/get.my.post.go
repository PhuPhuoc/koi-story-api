package marketrepository

import marketmodel "github.com/PhuPhuoc/koi-story-api/controller/market_services/model"

func (store *marketStore) GetMyPost(my_id string) ([]marketmodel.PostMarket, error) {
	var posts []marketmodel.PostMarket

	query := `
        select p.id as post_id, dm.product_name, dm.price, coalesce(pi.file_path, '') as file_path
        from post p
        join detail_market dm ON p.id = dm.post_id
        LEFT JOIN
		(
    		select post_id, file_path, ROW_NUMBER() OVER (PARTITION BY post_id ORDER BY image_order ASC) as rn
    		from post_image
    		where deleted_at IS NULL
		) pi ON p.id = pi.post_id AND pi.rn = 1
		where p.deleted_at is null and p.user_id=?
		order by p.created_at desc
    `
	err := store.db.Select(&posts, query, my_id)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
