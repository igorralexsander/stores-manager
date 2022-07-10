package model

import (
	"github.com/google/uuid"
	"time"
)

type Store struct {
	ID         *uuid.UUID    `json:"id,omitempty"`
	Name       string        `json:"name"`
	Url        string        `json:"url"`
	MaxTimeout time.Duration `json:"maxTimeout"`
	Group      string        `json:"group"`
}
