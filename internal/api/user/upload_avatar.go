package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UploadAvatar(ctx context.Context, req *desc.UploadAvatarRequest) (*desc.UploadAvatarResponse, error) {
	fileBytes := req.GetImage()
	if len(fileBytes) == 0 {
		return nil, status.Error(codes.InvalidArgument, "image is empty")
	}

	filename := req.GetFilename()
	if len(filename) == 0 {
		return nil, status.Error(codes.InvalidArgument, "filename is empty")
	}

	res, err := i.userService.UploadAvatar(ctx, fileBytes, filename)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.UploadAvatarResponse{
		Link: res,
	}, nil
}
