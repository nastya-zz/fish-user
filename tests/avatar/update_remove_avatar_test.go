package avatar

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
	_test "user/tests"
	"user/tests/suite"
)

func TestUpdateRemoveAvatar_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	resAva, err := st.UserClient.UploadAvatar(ctx, &desc.UploadAvatarRequest{
		UserId:   _test.UserId,
		Image:    []byte("test"),
		Filename: "test.jpg",
	})

	require.NoError(t, err)
	assert.NotEmpty(t, resAva.GetLink())

	resUser, err := st.UserClient.GetProfile(ctx, &desc.GetProfileRequest{
		Id: _test.UserId,
	})
	if err != nil {
		st.T.Fatal(err)
	}

	_, err = st.UserClient.UpdateProfile(ctx, &desc.UpdateProfileRequest{
		Info: &desc.UpdateProfile{
			Id:         _test.UserId,
			Bio:        &wrapperspb.StringValue{Value: resUser.Profile.Bio},
			AvatarPath: &wrapperspb.StringValue{Value: resAva.GetLink()},
			Email:      &wrapperspb.StringValue{Value: resUser.Profile.Email},
			Name:       &wrapperspb.StringValue{Value: resUser.Profile.Name},
			IsPublic:   &wrapperspb.BoolValue{Value: resUser.Profile.IsPublic},
		},
	})
	if err != nil {
		st.T.Fatal(err)
	}

	remRes, err := st.UserClient.RemoveAvatar(ctx, &desc.RemoveAvatarRequest{
		UserId:   _test.UserId,
		Filename: resAva.GetLink(),
	})

	require.NoError(t, err)
	assert.Empty(t, remRes.String())
}

func TestRemoveAvatar_Negative(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		name        string
		userId      string
		filename    string
		expectedErr string
	}{
		{
			name:        "user id is required",
			userId:      "",
			filename:    "345345",
			expectedErr: "user id is required",
		},
		{
			name:        "filename empty text",
			userId:      _test.UserId,
			expectedErr: "filename is required",
		},
		{
			name:        "filename not exist",
			userId:      _test.UserId,
			filename:    "w453 345l mkh",
			expectedErr: "cannot remove avatar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.UserClient.RemoveAvatar(ctx, &desc.RemoveAvatarRequest{
				UserId:   tt.userId,
				Filename: tt.filename,
			})

			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
