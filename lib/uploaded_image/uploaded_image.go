package uploaded_image;

import (
  "github.com/nu7hatch/gouuid"
  "mime/multipart"
)

type UploadedImage struct {
  File *multipart.FileHeader
  uid *uuid.UUID
}

func New(file *multipart.FileHeader) (u_img *UploadedImage, err error)  {
  u, err := uuid.NewV4()
  if err != nil {
    return
  }
  u_img = &UploadedImage{
    File: file,
    uid: u,
  }
  return
}

func (u_img *UploadedImage) String() string {
  return u_img.uid.String()
}
