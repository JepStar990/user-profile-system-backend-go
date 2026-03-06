package storage

import (
    "bytes"
    "context"
    "fmt"
    "mime/multipart"
    "os"
    "time"

    "github.com/aws/aws-sdk-go-v2/aws"
    s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
    "github.com/aws/aws-sdk-go-v2/service/s3"

    "github.com/aws/aws-sdk-go-v2/config"
)

type S3Uploader struct {
    client     *s3.Client
    bucketName string
}

func NewS3Uploader() *S3Uploader {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        panic("cannot load AWS config: " + err.Error())
    }

    return &S3Uploader{
        client:     s3.NewFromConfig(cfg),
        bucketName: os.Getenv("AWS_S3_BUCKET"),
    }
}

func (u *S3Uploader) UploadAvatar(file *multipart.FileHeader, userID string) (string, error) {
    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    buf := new(bytes.Buffer)
    _, err = buf.ReadFrom(src)
    if err != nil {
        return "", err
    }

    key := fmt.Sprintf("avatars/%s_%d_%s", userID, time.Now().Unix(), file.Filename)

    _, err = u.client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket:      aws.String(u.bucketName),
        Key:         aws.String(key),
        Body:        bytes.NewReader(buf.Bytes()),
        ContentType: aws.String(file.Header.Get("Content-Type")),
        ACL:         s3types.ObjectCannedACLPublicRead,
    })
    if err != nil {
        return "", err
    }

    // Build public URL
    url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.bucketName, key)
    return url, nil
}
