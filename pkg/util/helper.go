package util

import (
	"github.com/google/uuid"
	"strconv"
)

func GetUniqueID(field *string) error {
	uid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	*field = strconv.Itoa(int(uid.ID()))
	return nil
}
