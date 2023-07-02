package mongodb

import "github.com/google/uuid"

type Model interface {
	GetID() *uuid.UUID
	SetID(id *uuid.UUID)
}
