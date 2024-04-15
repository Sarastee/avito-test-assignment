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
	Content   json.RawMessage `json:"content" swaggertype:"object" additionalProperties:"true"`
}

// UpdateBanner model struct
type UpdateBanner struct {
	IsActive  *bool            `json:"is_active,omitempty"`
	FeatureID *int64           `json:"feature_id,omitempty"`
	TagIDs    *[]int64         `json:"tag_ids,omitempty"`
	Content   *json.RawMessage `json:"content,omitempty" swaggertype:"object" additionalProperties:"true"`
}

// UpdateBannerSQL model struct
type UpdateBannerSQL struct {
	IsActive  sql.NullBool              `json:"is_active,omitempty"`
	FeatureID sql.NullInt64             `json:"feature_id,omitempty"`
	TagsIDs   sql.Null[[]int64]         `json:"tag_ids,omitempty"`
	Content   sql.Null[json.RawMessage] `json:"content,omitempty"`
}

// BannerID model struct
type BannerID struct {
	ID int64 `json:"banner_id"`
}

// Content model struct
type Content struct {
	Content   json.RawMessage `json:"content" swaggertype:"object" additionalProperties:"true"`
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
	Revisions []Content `json:"revisions" swaggertype:"object" additionalProperties:"true"`
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

// ValidateUpdateBanner function which validates UpdateBanner struct.
func ValidateUpdateBanner(data []byte) error {
	schema := validator.NewSchema(
		vjson.Boolean("is_active"),
		vjson.Integer("feature_id").Positive(),
		vjson.Array("tag_ids", vjson.Integer("id").Positive()),
		vjson.Object("content", vjson.NewSchema()),
	)

	return schema.ValidateBytes(data)
}
