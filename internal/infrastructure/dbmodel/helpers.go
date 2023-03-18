package dbmodel

import (
	"booking/internal/domain"
	"encoding/json"
	"fmt"
)

// scan from data into dst, dst must be a pointer.
func scan(data, dst interface{}) error {
	if data == nil {
		return domain.NewError(domain.ErrCodeDatabaseError, "Scanned data is null")
	}

	j, ok := data.([]uint8)
	if !ok {
		return domain.NewError(domain.ErrCodeDatabaseError, "Scanned data is not []uint8")
	}

	err := json.Unmarshal(j, dst)
	if err != nil {
		return domain.NewErrorWrap(err, domain.ErrCodeDatabaseError, fmt.Sprintf("Scanned data cannot be unmarshalled (%v)", err))
	}

	return nil
}
