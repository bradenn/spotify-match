package service

import (
	"spotify-match/schemas"
)

type ComparedUsers struct {
	Names []string `json:"names"`
}

func Compare(res schemas.CompareInput) (ComparedUsers, error) {

	return ComparedUsers{}, nil
}
