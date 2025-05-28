package converter

import (
	"encoding/json"
	"github.com/google/uuid"
	"user/internal/logger"
	"user/internal/model"
)

// todo refactor
func UserFromPayload(bs []byte) model.Profile {
	const op = "converter.UserFromPayload"
	var payload model.UserPayload

	if err := json.Unmarshal(bs, &payload); err != nil {
		logger.Warn(op, "err", err)

		return model.Profile{}
	}

	id, _ := uuid.Parse(payload.ID)
	return model.Profile{
		ID:         model.UserId(id.String()),
		Name:       payload.Name,
		Email:      payload.Email,
		IsVerified: payload.IsVerified,
		CreatedAt:  payload.CreatedAt,
	}
}
