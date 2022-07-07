package scylladb

import (
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

func ToCQLUUID(uuid uuid.UUID) gocql.UUID {
	result, _ := gocql.ParseUUID(uuid.String())
	return result
}

func FromCQLUUID(gocqlUUID gocql.UUID) *uuid.UUID {
	result, _ := uuid.Parse(gocqlUUID.String())
	return &result
}
