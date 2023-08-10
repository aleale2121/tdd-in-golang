package db_test

import (
	"testing"

	"github.com/aleale2121/golang-tdd/bookswap/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	ps := db.NewPostingService()
	b := db.Book{
		ID: uuid.New().String(),
	}
	err := ps.NewOrder(b)
	assert.Nil(t, err)
}
