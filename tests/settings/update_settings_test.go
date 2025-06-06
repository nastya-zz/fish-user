package settings

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	apierrors "user/pkg/api-errors"
	_test "user/tests"
	"user/tests/suite"
)

func TestUpdateSettings(t *testing.T) {
	ctx, st := suite.New(t)

	t.Run("success update settings", func(t *testing.T) {
		res, err := st.UserClient.GetSettings(ctx, &desc.GetSettingsRequest{Id: _test.UserId})

		require.NoError(t, err)
		settings := res.GetSettings()

		oppositeValues := getOppositeValues(settings)

		resUpd, err := st.UserClient.UpdateSettings(ctx, &desc.UpdateSettingsRequest{
			UserId:       _test.UserId,
			SettingsInfo: oppositeValues,
		})

		require.NoError(t, err)
		require.NotNil(t, resUpd.Settings)

		// проверяем сохранились ли указаные значения
		assert.Equal(t, oppositeValues.Availability.String(), resUpd.Settings.Availability.String())
		assert.Equal(t, oppositeValues.Language.String(), resUpd.Settings.Language.String())

		// проверяем что изменились данные сравнивая исходное значение
		assert.NotEqual(t, settings.Availability.String(), resUpd.Settings.Availability.String())
		assert.NotEqual(t, settings.Language.String(), resUpd.Settings.Language.String())
	})

	t.Run("negative update settings user id empty", func(t *testing.T) {
		_, err := st.UserClient.UpdateSettings(ctx, &desc.UpdateSettingsRequest{
			UserId:       "",
			SettingsInfo: &desc.AccountSettings{},
		})

		require.Error(t, err)
		require.Contains(t, err.Error(), apierrors.UserIdRequired)
	})

	t.Run("negative update settings user id not found", func(t *testing.T) {
		_, err := st.UserClient.UpdateSettings(ctx, &desc.UpdateSettingsRequest{
			UserId: "29784237xx",
			SettingsInfo: &desc.AccountSettings{
				Language:     desc.Language_RU,
				Availability: desc.Availability_PRIVATE,
			},
		})

		require.Error(t, err)
		require.Contains(t, err.Error(), apierrors.UserUpdateSettingsFailed)
	})

	t.Run("success reset settings user ", func(t *testing.T) {
		_, err := st.UserClient.ResetSettings(ctx, &desc.ResetSettingsRequest{
			UserId: _test.UserId,
		})

		require.NoError(t, err)
	})

	t.Run("negative reset settings user ", func(t *testing.T) {
		_, err := st.UserClient.ResetSettings(ctx, &desc.ResetSettingsRequest{
			UserId: "ldjkfsljfi",
		})

		require.Error(t, err)
		require.Contains(t, err.Error(), apierrors.UserUpdateSettingsFailed)
	})
}

func getOppositeValues(settings *desc.AccountSettings) *desc.AccountSettings {
	result := &desc.AccountSettings{}

	if settings.Language == desc.Language_ENG {
		result.Language = desc.Language_RU
	} else {
		result.Language = desc.Language_ENG
	}

	if settings.Availability == desc.Availability_PRIVATE {
		result.Availability = desc.Availability_PUBLIC
	} else {
		result.Availability = desc.Availability_PRIVATE
	}

	return result
}
