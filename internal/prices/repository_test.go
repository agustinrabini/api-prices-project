package prices

import (
	"testing"

	"github.com/agustinrabini/api-prices-project/internal/platform/storage"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestRepository_NewRepository(t *testing.T) {
	db := storage.NewMock()
	defer db.Close()
	repository := NewRepository(nil)
	assert.NotNil(t, repository)
}
