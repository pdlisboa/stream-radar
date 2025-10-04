package user

import "github.com/gofiber/fiber/v2"

// CreateUser
// @Summary Create user
// @Description Create a new user
// @Tags user
// @Produce json
// @Router /user [POST]
// @Param body body CreateUserRequest true "body"
// @success 201 {object} dto.UserDTO
// @Failure 400  {object} any
// @Failure 401 {object} any
// @Failure 403 {object} any
// @Failure 422 {object} any
// @Failure 500 {object} any
func CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := service.Create(req)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user.ToDTO())
}

// GetUser
// @Summary Get user
// @DescriptionGet new user
// @Tags user
// @Produce json
// @Router /user [GET]
// @Param name query FindUserRequest false "query"
// @Param Authorization header string true "Bearer"
// @success 201 {object} dto.UserDTO
// @Failure 400  {object} any
// @Failure 401 {object} any
// @Failure 403 {object} any
// @Failure 422 {object} any
// @Failure 500 {object} any
func GetUser(c *fiber.Ctx) error {
	var req FindUserRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := service.Find(req)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user.ToDTO())
}
