package user

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"user/tests/suite"
)

func TestGetUserProfile_HappyPath(t *testing.T) {
	const (
		userId = "72a9cbbc-0175-4e6c-92dd-98629b6aac41"
	)

	ctx, st := suite.New(t)

	res, err := st.UserClient.GetProfile(ctx, &desc.GetProfileRequest{Id: userId})

	require.NoError(t, err)

	assert.NotEmpty(t, res.GetProfile())
	assert.Equal(t, res.GetProfile().GetId(), userId)
}

func TestGetUserProfile_Negative(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		name        string
		userId      string
		expectedErr string
	}{
		{
			name:        "Id empty",
			userId:      "",
			expectedErr: "Не указан id пользователя",
		},
		{
			name:        "User not found",
			userId:      "not-exist-id",
			expectedErr: "Пользователь с таким Id не найден",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.UserClient.GetProfile(ctx, &desc.GetProfileRequest{Id: tt.userId})

			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
