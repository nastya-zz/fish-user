package converter

import (
	"encoding/json"
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

	return model.Profile{
		ID:         model.UserId(payload.ID),
		Name:       payload.Name,
		Email:      payload.Email,
		IsVerified: payload.IsVerified,
		CreatedAt:  payload.CreatedAt,
	}
}
