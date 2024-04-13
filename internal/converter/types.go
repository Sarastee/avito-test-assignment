package converter

import "database/sql"

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
