package services

import (
	"context"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService(cloudinaryURL string) (*CloudinaryService, error) {
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return nil, err
	}
	return &CloudinaryService{cld: cld}, nil
}

func (s *CloudinaryService) UploadAvatar(ctx context.Context, file io.Reader, userID string) (string, string, error) {
	overwrite := true
	uploadParams := uploader.UploadParams{
		Folder:         "chatear/users/avatars",
		PublicID:       userID,
		Format:         "webp",
		Overwrite:      &overwrite,
		Transformation: "f_auto,q_auto",
	}

	result, err := s.cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", "", err
	}

	return result.SecureURL, result.PublicID, nil
}

func (s *CloudinaryService) DeleteAvatar(ctx context.Context, publicID string) error {
	_, err := s.cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicID})
	return err
}
