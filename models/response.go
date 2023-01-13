package models

import "github.com/gofiber/fiber/v2"

type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
	Type    string `json:"type"`
}
