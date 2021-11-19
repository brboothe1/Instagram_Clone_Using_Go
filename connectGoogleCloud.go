package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"time"
)

// createBucket() creates a new bucket in Google Cloud Storage to hold new files
func createBucket() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	ctx := context.Background()

	projectID := "sonorous-parsec-332420"

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("error while initiating client: %v", err)
	}
	defer client.Close()

	bucketName := "instagram-rebuild-bucket"

	bucket := client.Bucket(bucketName)

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	fmt.Printf("Bucket %v created.\n", bucketName)
}

//	uploadFileToBucket() uploads new files to the Google Bucket
func uploadFileToBucket(w io.Writer, bucket, object string) error {

	// using godotenv.Load() to pull google cloud credentials from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}
	defer client.Close()

	f, err := os.Open("FileName.text")
	if err != nil {
		log.Fatalf("os.Open error: %v", err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	fmt.Fprintf(w, "Blob %v uploaded.\n", object)
	return nil
}
