package subscribe

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"github.com/stretchr/testify/require"
	"testing"
	api_errors "user/pkg/api-errors"
	"user/tests"
	"user/tests/suite"
)

func TestSubscribe(t *testing.T) {
	ctx, st := suite.New(t)

	t.Run("success subscribe/unsubscribe", func(t *testing.T) {
		userId := tests.UserId
		subscId := tests.TestUserId

		_, err := st.UserClient.Subscribe(ctx, &desc.SubscribeRequest{
			UserId:         userId,
			SubscriptionId: subscId,
		})

		require.NoError(t, err)

		_, err = st.UserClient.UnSubscribe(ctx, &desc.SubscribeRequest{
			UserId:         userId,
			SubscriptionId: subscId,
		})
		require.NoError(t, err)
	})

	t.Run("same id user and subscriber", func(t *testing.T) {
		userId := tests.UserId
		subscId := tests.UserId

		_, err := st.UserClient.Subscribe(ctx, &desc.SubscribeRequest{
			UserId:         userId,
			SubscriptionId: subscId,
		})

		require.Error(t, err)
		require.Contains(t, err.Error(), api_errors.UserSubscribeCannotBeSame)
	})

	t.Run(" user id is empty", func(t *testing.T) {
		userId := ""
		subscId := tests.UserId

		_, err := st.UserClient.Subscribe(ctx, &desc.SubscribeRequest{
			UserId:         userId,
			SubscriptionId: subscId,
		})

		require.Error(t, err)
		require.Contains(t, err.Error(), api_errors.UserIdRequired)
	})

	t.Run(" user id is empty", func(t *testing.T) {
		userId := tests.UserId
		subscId := ""

		_, err := st.UserClient.Subscribe(ctx, &desc.SubscribeRequest{
			UserId:         userId,
			SubscriptionId: subscId,
		})

		require.Error(t, err)
		require.Contains(t, err.Error(), api_errors.UserIdRequired)
	})
}
