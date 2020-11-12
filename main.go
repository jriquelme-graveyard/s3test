package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	var bucket, key string
	flag.StringVar(&bucket, "bucket", "", "bucket name")
	flag.StringVar(&key, "key", "", "key to get")
	flag.Parse()
	switch {
	case bucket == "":
		log.Fatal("missing bucket")
	case key == "":
		log.Fatal("missing key")
	default:
		err := CopyFileFromS3(bucket, key)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CopyFileFromS3(bucket, key string) error {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		return err
	}
	client := s3.NewFromConfig(cfg)
	output, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(output.Body)
	if err != nil {
		return err
	}
	log.Printf("%s", b)
	return nil
}
