package main

import (
	"testing"
)

func TestCopyFileFromS3(t *testing.T) {
	t.Run("doesnt work", func(t *testing.T) {
		bucket := "aws-go-sdk-v2-883"
		// this key works as expected with v0.24.0!
		file := "private/us-east-1:bc324f78-8b99-4359-a9a7-b1df7ecea360/customers_2.csv"
		err := CopyFileFromS3(bucket, file)
		if err != nil {
			t.Fatalf("oh no: %s", err)
		}
	})
	t.Run("ok", func(t *testing.T) {
		bucket := "aws-go-sdk-v2-883"
		file := "my-file.txt"
		err := CopyFileFromS3(bucket, file)
		if err != nil {
			t.Fatalf("oh no: %s", err)
		}
	})
}
