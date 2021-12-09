package repository

import (
	"os"
	"testing"

	"github.com/Jailtons7/imersao-gateway/adapter/repository/fixture"
	"github.com/Jailtons7/imersao-gateway/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDBInsert(t *testing.T) {
	migrationDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)
	repository := NewTransactionRepositoryDB(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}
