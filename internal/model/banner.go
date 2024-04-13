package model

import (
	"encoding/json"

	"github.com/miladibra10/vjson"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// CreateBanner model struct
type CreateBanner struct {
	IsActive  bool            `json:"is_active"`
	FeatureID int64           `json:"feature_id"`
	TagsIDs   []int64         `json:"tag_ids"`
	Content   json.RawMessage `json:"content"`
}

// Banner model struct
type Banner struct {
	TagIDs    []int64         `json:"tag_ids"`
	FeatureID *int64          `json:"feature_id"`
	Content   json.RawMessage `json:"content"`
	IsActive  *bool           `json:"is_active"`
}

// BannerID model struct
type BannerID struct {
	ID int64 `json:"banner_id"`
}

// ValidateCreateBanner function which validates CreateBanner struct.
func ValidateCreateBanner(data []byte) error {
	schema := validator.NewSchema(
		vjson.Boolean("is_active").Required(),
		vjson.Integer("feature_id").Positive().Required(),
		vjson.Array("tag_ids", vjson.Integer("id").Positive().Required()),
		vjson.Object("content", vjson.NewSchema()).Required(),
	)

	return schema.ValidateBytes(data)
}
