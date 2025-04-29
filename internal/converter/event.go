package converter

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"user/internal/model"
)

// todo refactor
func UserFromPayload(bs []byte) model.Profile {
	var payload model.UserPayload

	if err := json.Unmarshal(bs, &payload); err != nil {
		log.Println(err)

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
