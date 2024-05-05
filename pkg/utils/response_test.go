package utils

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)


func TestResponse(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

	t.Run("JSONWithCode", func(t *testing.T) {
		result := JSONWithCode(ctx, 200, "success", nil, 1)
		assert.Equal(t, result, nil)
	})

	t.Run("ResponseOk", func(t *testing.T) {
		result := ResponseOK(ctx, "success", nil)
		assert.Equal(t, result, nil)
	})

	t.Run("ResponseCreated", func(t *testing.T) {
		result := ResponseCreated(ctx, "created", nil)
		assert.Equal(t, result, nil)
	})

	t.Run("ResponseNotFound", func(t *testing.T) {
		result := ResponseNotFound(ctx, "bad request")
		assert.Equal(t, result, nil)
	})

	t.Run("ResponseBadRequest", func(t *testing.T) {
		result := ResponseBadRequest(ctx, "bad request")
		assert.Equal(t, result, nil)
	})
	
	t.Run("ResponseInternalServerError", func(t *testing.T) {
		result := ResponseInternalServerError(ctx, "bad request")
		assert.Equal(t, result, nil)
	})

}