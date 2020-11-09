package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg)
	output, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String("my-file.csv"),
	})
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(output.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", b)
}
