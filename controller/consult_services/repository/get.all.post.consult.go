package consultrepository

import (
	"fmt"

	consultmodel "github.com/PhuPhuoc/koi-story-api/controller/consult_services/model"
)

func (store *consultStore) GetListConsult() ([]consultmodel.PostConsult, error) {
	var result []consultmodel.PostConsult

	query := `
	SELECT
		p.id,
		p.user_id,
		u.display_name AS user_name,
		u.profile_picture_url AS user_avatar,
		p.post_type,
		p.title,
		c.content,
		pim.file_path,
		p.created_at
	FROM
		post p
	JOIN
		detail_consult c ON p.id = c.post_id
	JOIN
		user u ON p.user_id = u.id
	LEFT JOIN
		post_image pim ON p.id = pim.post_id
	WHERE
		p.deleted_at is null
	ORDER BY
		p.created_at DESC
	`

	err := store.db.Select(&result, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get list of consults: %w", err)
	}

	return result, nil
}
