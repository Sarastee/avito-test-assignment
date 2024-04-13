package model

import (
	"database/sql"
	"encoding/json"
	"time"

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

// BannerID model struct
type BannerID struct {
	ID int64 `json:"banner_id"`
}

// Content model struct
type Content struct {
	Content   json.RawMessage `json:"content"`
	Revision  int64           `json:"revision_id"`
	CreatedAt time.Time       `json:"created_at"`
}

// Banner model struct
type Banner struct {
	ID        int64     `json:"banner_id"`
	IsActive  *bool     `json:"is_active"`
	FeatureID *int64    `json:"feature_id"`
	TagIDs    []int64   `json:"tag_ids"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Revisions []Content `json:"revisions"`
}

// BannerInfo model struct
type BannerInfo struct {
	FeatureID sql.NullInt64
	TagID     sql.NullInt64
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
