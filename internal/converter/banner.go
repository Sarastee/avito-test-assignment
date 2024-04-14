package converter

import (
	"github.com/sarastee/avito-test-assignment/internal/model"
)

// UpdateBannerToUpdateBannerSQL converts model.UpdateBanner to model.UpdateBannerSQL
func UpdateBannerToUpdateBannerSQL(bnr *model.UpdateBanner) model.UpdateBannerSQL {
	return model.UpdateBannerSQL{
		IsActive:  BoolPointerToSQLNullBool(bnr.IsActive),
		FeatureID: Int64PointerToSQLNullInt64(bnr.FeatureID),
		TagsIDs:   Int64SlicePointerToSQLNullInt64Slice(bnr.TagIDs),
		Content:   RawMessagePointerToSQLNullRawMessage(bnr.Content),
	}
}
