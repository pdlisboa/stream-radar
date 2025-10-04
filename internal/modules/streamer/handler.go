package streamer

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"stream-radar/domain/dto"
	"stream-radar/internal/logger"
	"stream-radar/internal/modules/user"
	"stream-radar/internal/utils"
)

var userService user.IUserService

func init() {
	userService = &user.UserService{}
}

// RegisterStreamers
// @Summary Register Streamers
// @DescriptionGet Register Streamers
// @Tags streamers
// @Produce json
// @Router /streamer [post]
// @Param body body RegisterStreamersRequest true "payload"
// @Param Authorization header string true "Bearer"
// @success 201 {object} []dto.StreamerDTO
// @Failure 400  {object} any
// @Failure 401 {object} any
// @Failure 403 {object} any
// @Failure 422 {object} any
// @Failure 500 {object} any
func RegisterStreamers(c *fiber.Ctx) error {
	id := utils.GetUserIdInRequest(c)
	user, err := userService.Get(id)

	log := logger.GetInstance()

	if err != nil {
		log.Debug("User not found", zap.Uint("id", id))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	var req RegisterStreamersRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error("Error parsing body RegisterStreamersRequest", zap.Error(err))
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	streamers, err := service.RegisterStreamers(req.Streamers, user)
	if err != nil {
		return err
	}

	var response []dto.StreamerDTO

	for _, streamer := range streamers {
		response = append(response, streamer.ToDto())
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
