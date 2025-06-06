package user

import (
	"github.com/brianvoe/gofakeit/v6"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
	api_errors "user/pkg/api-errors"
	_test "user/tests"
	"user/tests/suite"
)

func TestUpdateUserProfile_Bio(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		name        string
		bio         string
		expectedBio string
	}{
		{
			name:        "Bio empty text",
			bio:         "",
			expectedBio: "",
		},
		{
			name:        "Bio with text",
			bio:         "new bio",
			expectedBio: "new bio",
		},
	}

	resGetUser, errGet := st.UserClient.GetProfile(ctx, &desc.GetProfileRequest{Id: _test.UserId})
	require.NoError(t, errGet)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resUpdUser, err := st.UserClient.UpdateProfile(ctx, &desc.UpdateProfileRequest{Info: &desc.UpdateProfile{
				Id:         _test.UserId,
				Bio:        &wrapperspb.StringValue{Value: tt.bio},
				AvatarPath: &wrapperspb.StringValue{Value: resGetUser.Profile.AvatarPath},
				Email:      &wrapperspb.StringValue{Value: resGetUser.Profile.Email},
				Name:       &wrapperspb.StringValue{Value: resGetUser.Profile.Name},
				IsPublic:   &wrapperspb.BoolValue{Value: resGetUser.Profile.IsPublic},
			}})

			require.NoError(t, err)

			// Проверяем, что bio обновилось как ожидалось
			assert.Equal(t, tt.expectedBio, resUpdUser.Profile.Bio)

			// Проверяем, что остальные поля не изменились
			assert.Equal(t, resGetUser.Profile.AvatarPath, resUpdUser.Profile.AvatarPath)
			assert.Equal(t, resGetUser.Profile.Email, resUpdUser.Profile.Email)
			assert.Equal(t, resGetUser.Profile.Name, resUpdUser.Profile.Name)
			assert.Equal(t, resGetUser.Profile.IsPublic, resUpdUser.Profile.IsPublic)

		})
	}
}

func TestUpdateUserProfile_Name(t *testing.T) {
	ctx, st := suite.New(t)

	fakeName := gofakeit.Username()

	tests := []struct {
		name         string
		username     string
		expectedName string
	}{
		{
			name:         "username empty text",
			username:     "",
			expectedName: "",
		},
		{
			name:         "username with text",
			username:     fakeName,
			expectedName: fakeName,
		},
	}

	resGetUser, errGet := st.UserClient.GetProfile(ctx, &desc.GetProfileRequest{Id: _test.UserId})
	require.NoError(t, errGet)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resUpdUser, err := st.UserClient.UpdateProfile(ctx, &desc.UpdateProfileRequest{Info: &desc.UpdateProfile{
				Id:         _test.UserId,
				Bio:        &wrapperspb.StringValue{Value: resGetUser.Profile.Bio},
				AvatarPath: &wrapperspb.StringValue{Value: resGetUser.Profile.AvatarPath},
				Email:      &wrapperspb.StringValue{Value: resGetUser.Profile.Email},
				Name:       &wrapperspb.StringValue{Value: tt.username},
				IsPublic:   &wrapperspb.BoolValue{Value: resGetUser.Profile.IsPublic},
			}})

			require.NoError(t, err)

			// Проверяем, что name обновилось как ожидалось
			assert.Equal(t, tt.username, resUpdUser.Profile.Name)

			// Проверяем, что остальные поля не изменились
			assert.Equal(t, resGetUser.Profile.AvatarPath, resUpdUser.Profile.AvatarPath)
			assert.Equal(t, resGetUser.Profile.Email, resUpdUser.Profile.Email)
			assert.Equal(t, resGetUser.Profile.Bio, resUpdUser.Profile.Bio)
			assert.Equal(t, resGetUser.Profile.IsPublic, resUpdUser.Profile.IsPublic)
		})
	}

}

func TestUpdateUserProfile_Email_Negative(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		name        string
		email       string
		expectedErr string
	}{
		{
			name:        "email empty text",
			email:       "",
			expectedErr: api_errors.UserEmailEmpty,
		},
		{
			name:        "not match pattern",
			email:       "w453 345l mkh",
			expectedErr: api_errors.UserEmailNotMatchPattern,
		},
	}

	resGetUser, errGet := st.UserClient.GetProfile(ctx, &desc.GetProfileRequest{Id: _test.UserId})
	require.NoError(t, errGet)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.UserClient.UpdateProfile(ctx, &desc.UpdateProfileRequest{Info: &desc.UpdateProfile{
				Id:         _test.UserId,
				Bio:        &wrapperspb.StringValue{Value: resGetUser.Profile.Bio},
				AvatarPath: &wrapperspb.StringValue{Value: resGetUser.Profile.AvatarPath},
				Name:       &wrapperspb.StringValue{Value: resGetUser.Profile.Name},
				Email:      &wrapperspb.StringValue{Value: tt.email},
				IsPublic:   &wrapperspb.BoolValue{Value: resGetUser.Profile.IsPublic},
			}})

			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}

}
