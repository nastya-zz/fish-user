package user

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"user/tests"
	"user/tests/suite"
)

func TestGetUserProfile_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	res, err := st.UserClient.GetProfile(ctx, &desc.GetProfileRequest{Id: tests.UserId})

	require.NoError(t, err)

	assert.NotEmpty(t, res.GetProfile())
	assert.Equal(t, res.GetProfile().GetId(), tests.UserId)
}

func TestGetUserProfile_Negative(t *testing.T) {
	ctx, st := suite.New(t)

	testSlice := []struct {
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

	for _, tt := range testSlice {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.UserClient.GetProfile(ctx, &desc.GetProfileRequest{Id: tt.userId})

			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
