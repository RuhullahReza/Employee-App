package utils

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BodyResponse struct {
	Code       string      `json:"code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	ServerTime int64       `json:"serverTime"`
}

func NewResponseBody(code string, msg string, data interface{}, t ...int64) *BodyResponse {
	m := new(BodyResponse)
	m.Code = code
	m.Message = msg
	m.Data = data
	if len(t) > 0 {
		m.ServerTime = t[0]
	} else {
		m.ServerTime = time.Now().UTC().UnixNano() / 1000000
	}

	return m
}

func JSONWithCode(ctx *fiber.Ctx, code int, msg string, data interface{}, t ...int64) error {
	return ctx.Status(code).JSON(NewResponseBody(http.StatusText(code), msg, data, t...))
}

func ResponseOK(ctx *fiber.Ctx, msg string, data interface{}) error {
	return JSONWithCode(ctx, fiber.StatusOK, msg, data)
}

func ResponseCreated(ctx *fiber.Ctx, msg string, data interface{}) error {
	return JSONWithCode(ctx, fiber.StatusCreated, msg, data)
}

func ResponseNotFound(ctx *fiber.Ctx, msg string) error {
	return JSONWithCode(ctx, fiber.StatusNotFound, msg, nil)
}

func ResponseBadRequest(ctx *fiber.Ctx, msg string) error {
	return JSONWithCode(ctx, fiber.StatusBadRequest, msg, nil)
}

func ResponseInternalServerError(ctx *fiber.Ctx, msg string) error {
	return JSONWithCode(ctx, fiber.StatusInternalServerError, msg, nil)
}