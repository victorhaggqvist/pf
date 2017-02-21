package pf

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeToSlug(t *testing.T) {
	assert := assert.New(t)
	slug := CodeToSlug(http.StatusInternalServerError)

	assert.Equal(InternalServerErrorSlug, slug)
}
