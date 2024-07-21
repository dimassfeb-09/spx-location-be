package services

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"io"
	"time"
)

var app *firebase.App

func InitFirebaseApp(ctx context.Context, credentialsPath string) error {
	opt := option.WithCredentialsFile(credentialsPath)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	return err
}

func UploadImageURLOrder(ctx context.Context, file io.Reader, fileName string) (string, error) {
	client, err := app.Storage(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting Storage client: %v", err)
	}

	bucketName := "spx-location.appspot.com" // Replace with your bucket name
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", fmt.Errorf("error getting Storage client: %v", err)
	}

	timeNowUnix := time.Now().Unix()
	obj := bucket.Object(fmt.Sprintf("uploads/%d-%s", timeNowUnix, fileName))
	writer := obj.NewWriter(ctx)

	if _, err := io.Copy(writer, file); err != nil {
		return "", fmt.Errorf("error copying file to Firebase Storage: %v", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("error closing writer: %v", err)
	}

	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/", bucketName)
	url += "uploads%2F"
	url += fmt.Sprintf("%d-%s?alt=media&token=9442bfff-329c-4c4c-9aec-90af3565b757", timeNowUnix, fileName)
	return url, nil
}
