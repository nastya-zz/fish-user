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
