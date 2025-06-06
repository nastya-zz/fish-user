package settings

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	api_errors "user/pkg/api-errors"
	_test "user/tests"
	"user/tests/suite"
)

func TestGetUserSettings(t *testing.T) {
	ctx, st := suite.New(t)

	t.Run("success get settings", func(t *testing.T) {
		res, err := st.UserClient.GetSettings(ctx, &desc.GetSettingsRequest{Id: _test.UserId})

		require.NoError(t, err)
		settings := res.GetSettings()

		require.NotNil(t, settings)
		assert.Equal(t, "PUBLIC", settings.Availability.String())
		assert.Equal(t, "RU", settings.Language.String())
	})

	t.Run("negative get settings userId empty", func(t *testing.T) {
		_, err := st.UserClient.GetSettings(ctx, &desc.GetSettingsRequest{Id: ""})

		require.Error(t, err)
		require.Contains(t, err.Error(), api_errors.UserIdRequired)
	})

	t.Run("negative get settings userId is not exist", func(t *testing.T) {
		_, err := st.UserClient.GetSettings(ctx, &desc.GetSettingsRequest{Id: "jk34jh5g3"})

		require.Error(t, err)
		require.Contains(t, err.Error(), api_errors.UserGetSettingsFailed)
	})

}
