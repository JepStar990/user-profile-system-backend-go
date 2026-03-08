package storage

import (
    "context"
    "os"
    "time"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    s3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Presigner generates short-lived presigned URLs for S3 objects.
// NOTE: The presign client lives in the s3 package (s3.NewPresignClient),
// not in a separate "service/s3/presign" import path.
type S3Presigner struct {
    Client    *s3.Client
    Presigner *s3.PresignClient
    Bucket    string
}

// NewS3Presigner initializes an S3 client and presign client using the AWS default config chain.
func NewS3Presigner() *S3Presigner {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        panic("cannot load AWS config: " + err.Error())
    }

    client := s3.NewFromConfig(cfg)
    presigner := s3.NewPresignClient(client)

    return &S3Presigner{
        Client:    client,
        Presigner: presigner,
        Bucket:    os.Getenv("AWS_S3_BUCKET"),
    }
}

// GenerateDownloadURL returns a temporary presigned URL for downloading an S3 object key.
// Default expiry: 15 minutes (adjustable).
func (p *S3Presigner) GenerateDownloadURL(key string) (string, error) {
    input := &s3.GetObjectInput{
        Bucket: aws.String(p.Bucket),
        Key:    aws.String(key),
    }

    // Use PresignGetObject with an expiry option.
    out, err := p.Presigner.PresignGetObject(
        context.TODO(),
        input,
        func(opts *s3.PresignOptions) {
            opts.Expires = 15 * time.Minute
        },
    )
    if err != nil {
        return "", err
    }

    return out.URL, nil
}
