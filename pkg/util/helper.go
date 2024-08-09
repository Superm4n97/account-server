package util

import (
	"github.com/google/uuid"
	"strconv"
)

func GetUniqueID() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(uid.ID())), nil
}
