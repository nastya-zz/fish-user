package converter

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"user/internal/model"
)

func ToDescProfileFromProfile(profile *model.Profile) *desc.UserProfile {
	return &desc.UserProfile{
		Id:              string(profile.ID),
		Email:           profile.Email,
		Name:            profile.Name,
		IsVerified:      profile.IsVerified,
		IsPublic:        profile.IsPublic,
		CreatedAt:       timestamppb.New(profile.CreatedAt),
		Bio:             profile.Bio,
		ExperienceLevel: int32(profile.ExperienceLevel),
		AvatarPath:      profile.AvatarPath,
	}
}
