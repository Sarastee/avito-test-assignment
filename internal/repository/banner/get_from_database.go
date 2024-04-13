package banner

// GetBannerFromDatabase function which get banner instance from database
//func (r Repo) GetBannerFromDatabase(ctx context.Context, tagID int64, featureID int64, isRoleAdmin bool) (string, error) {
//	queryFormat := `
//	SELECT br.%s
//	FROM  %s br
//	JOIN %s brt ON br.%s = brt.%s
//	WHERE ((br.%s = TRUE) OR (@%s = TRUE))
//		AND br.%s = @%s
//		AND brt.%s = @%s
//		AND br.%s IN (
//	    SELECT %s
//		FROM %s
//		WHERE %s = br.%s
//	)
//	`
//
//	query := fmt.Sprintf(
//		queryFormat,
//		contentColumn,
//		bannerRevisionsTable,
//		bannerRevisionTagsTable, revisionIDColumn, revisionIDColumn,
//		isActiveColumn, isRoleAdminColumn,
//		featureIDColumn, featureIDColumn,
//		tagIDColumn, tagIDColumn,
//		bannerIDColumn,
//		bannerIDColumn,
//		bannersTable,
//		selectedRevisionIDColumn,
//		revisionIDColumn,
//	)
//	q := db.Query{
//		Name:     "banner_repository.GetBannerFromDatabase",
//		QueryRaw: query,
//	}
//
//	args := pgx.NamedArgs{
//		featureIDColumn:   featureID,
//		tagIDColumn:       tagID,
//		isRoleAdminColumn: isRoleAdmin,
//	}
//
//	rows, err := r.db.DB().QueryContext(ctx, q, args)
//	if err != nil {
//		return "", err
//	}
//
//	banner, err := pgx.CollectOneRow(rows, pgx.RowTo[string])
//	if err != nil {
//		if errors.Is(err, pgx.ErrNoRows) {
//			return "", repository.ErrBannerNotFound
//		}
//
//		return "", err
//	}
//
//	return banner, nil
//}
