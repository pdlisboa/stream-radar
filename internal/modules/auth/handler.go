package auth

import "github.com/gofiber/fiber/v2"

// Login
// @Summary Login
// @Description Login
// @Tags auth
// @Produce json
// @Router /login [POST]
// @Param body body LoginRequest true "body"
// @success 201 {object} string
// @Failure 400  {object} any
// @Failure 401 {object} any
// @Failure 403 {object} any
// @Failure 422 {object} any
// @Failure 500 {object} any
func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := service.Login(req)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(token)

}
