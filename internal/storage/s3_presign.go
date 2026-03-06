package storage

import (
    "context"
    "time"
    "os"

    "github.com/aws/aws-sdk-go-v2/aws"
    s3 "github.com/aws/aws-sdk-go-v2/service/s3"
    s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
    "github.com/aws/aws-sdk-go-v2/service/s3/presign"
    "github.com/aws/aws-sdk-go-v2/config"
)

// S3Presigner creates temporary URLs for direct client downloads.
type S3Presigner struct {
    Client    *s3.Client
    Presigner *presign.PresignClient
    Bucket    string
}

func NewS3Presigner() *S3Presigner {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        panic("cannot load AWS config: " + err.Error())
    }

    client := s3.NewFromConfig(cfg)

    return &S3Presigner{
        Client:    client,
        Presigner: presign.NewPresignClient(client),
        Bucket:    os.Getenv("AWS_S3_BUCKET"),
    }
}

// GenerateDownloadURL returns a temporary signed URL for S3 downloads.
func (p *S3Presigner) GenerateDownloadURL(key string) (string, error) {

    input := &s3.GetObjectInput{
        Bucket: aws.String(p.Bucket),
        Key:    aws.String(key),
    }

    // Default: 15 minutes expiry
    presigned, err := p.Presigner.PresignGetObject(
        context.TODO(),
        input,
        presign.WithExpires(15*time.Minute),
    )
    if err != nil {
        return "", err
    }

    return presigned.URL, nil
}
