package converter

import (
	"database/sql"
	"encoding/json"
)

// Int64PointerToSQLNullInt64 converts Int64 Pointer to SQL Nullable Int64
func Int64PointerToSQLNullInt64(intPtr *int64) sql.NullInt64 {
	var out sql.NullInt64
	if intPtr != nil {
		out = sql.NullInt64{Int64: *intPtr, Valid: true}
	} else {
		out = sql.NullInt64{Valid: false}
	}

	return out
}

// BoolPointerToSQLNullBool converts Bool Pointer to SQL Nullable Bool
func BoolPointerToSQLNullBool(boolPtr *bool) sql.NullBool {
	var out sql.NullBool
	if boolPtr != nil {
		out = sql.NullBool{Bool: *boolPtr, Valid: true}
	} else {
		out = sql.NullBool{Valid: false}
	}

	return out
}

// Int64SlicePointerToSQLNullInt64Slice converts Int64 Slice Pointer to SQL Null Int64 Slice
func Int64SlicePointerToSQLNullInt64Slice(intSlicePtr *[]int64) sql.Null[[]int64] {
	var out sql.Null[[]int64]
	if intSlicePtr != nil {
		if len(*intSlicePtr) > 0 {
			out = sql.Null[[]int64]{V: *intSlicePtr, Valid: true}
		} else {
			out = sql.Null[[]int64]{Valid: false}
		}
	} else {
		out = sql.Null[[]int64]{Valid: false}
	}

	return out
}

// RawMessagePointerToSQLNullRawMessage converts json.RawMessage pointer to SQL Null json.RawMessage
func RawMessagePointerToSQLNullRawMessage(rawPtr *json.RawMessage) sql.Null[json.RawMessage] {
	var out sql.Null[json.RawMessage]
	if rawPtr != nil {
		out = sql.Null[json.RawMessage]{V: *rawPtr, Valid: true}
	} else {
		out = sql.Null[json.RawMessage]{Valid: false}
	}

	return out
}
