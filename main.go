package main

import(
	"context"
	"log"
	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main(){
	cfg , err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Println("Error loading AWS config:", err)
		return
	}

	client := s3.NewFromConfig(cfg)

	resp, err := client.ListBuckets(context.TODO(), nil)
	if err != nil {
		log.Println("Error listing S3 buckets:", err)
		return
	}


	log.Println("S3 Buckets list: ")
	for _, bucket := range resp.Buckets{
		log.Println(*bucket.Name)
	}
}