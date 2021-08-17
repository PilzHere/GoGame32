package entity

import (
	"github.com/google/uuid"
)

type Entity struct {
	Id        uuid.UUID
	ToDestroy bool
}
