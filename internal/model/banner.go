package model

import "encoding/json"

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
