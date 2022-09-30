package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	response := Ok("ok")
	assert.Equal(t, 200, response.StatusCode)
}

func TestBadRequest(t *testing.T) {
	response := BadRequest("invalid", "error fake")
	assert.Equal(t, 400, response.StatusCode)
}

func TestInternalServerError(t *testing.T) {
	response := InternalServerError("internal server error", "error fake")
	assert.Equal(t, 500, response.StatusCode)
}
