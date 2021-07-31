package repository

// func (repo *repository) BulkUpsertCabang(request []*model.Cabang) ([]*model.Cabang, error) {
// 	addsQuery := make([]string, 0, len(request))

// 	for _, cabang := range request {
// 		addsQuery = append(addsQuery, fmt.Sprintf("(SELECT * FROM cabang WHERE cabang = %s)", cabang.Cabang))
// 	}

// 	if len(request) > 0 {
// 		query := `INSERT INTO embrio4.cabang (id, cabang, description, created_at, updated_at)` + strings.Join(addsQuery, ",") +
// 			`ON DUPLICATE KEY UPDATE SET
// 			cabang = :cabang,
// 			description = :description,
// 			updated_at = NOW()
// 		`

// 		_, err := repo.db.Exec(query, request)
// 		if err != nil {
// 			log.Error(err)
// 			return nil, err
// 		}
// 	}

// 	return request, nil
// }
