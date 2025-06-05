package subscribe

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"github.com/stretchr/testify/require"
	"testing"
	"user/tests"
	"user/tests/suite"
)

func TestUnsubscribe(t *testing.T) {
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

	t.Run("negative subscribe/unsubscribe", func(t *testing.T) {
		userId := tests.TestUserId
		subscId := tests.UserId

		_, err := st.UserClient.UnSubscribe(ctx, &desc.SubscribeRequest{
			UserId:         userId,
			SubscriptionId: subscId,
		})

		require.Error(t, err)
		require.Contains(t, err.Error(), "Не удалось удалить подписку")
	})
}
