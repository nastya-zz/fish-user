package settings

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"user/tests/suite"
)

func TestGetUserSettings(t *testing.T) {
	const (
		userId = "72a9cbbc-0175-4e6c-92dd-98629b6aac41"
	)

	ctx, st := suite.New(t)

	t.Run("success get settings", func(t *testing.T) {
		res, err := st.UserClient.GetSettings(ctx, &desc.GetSettingsRequest{Id: userId})

		require.NoError(t, err)
		settings := res.GetSettings()

		require.NotNil(t, settings)
		assert.Equal(t, "PUBLIC", settings.Availability.String())
		assert.Equal(t, "RU", settings.Language.String())
	})

	t.Run("negative get settings userId empty", func(t *testing.T) {
		_, err := st.UserClient.GetSettings(ctx, &desc.GetSettingsRequest{Id: ""})

		require.Error(t, err)
		require.Contains(t, err.Error(), "missing user id")
	})

	t.Run("negative get settings userId is not exist", func(t *testing.T) {
		_, err := st.UserClient.GetSettings(ctx, &desc.GetSettingsRequest{Id: "jk34jh5g3"})

		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to get user settings")
	})

}
