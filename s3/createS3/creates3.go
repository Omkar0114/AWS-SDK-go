package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main(){

    cfg , err := config.LoadDefaultConfig(context.TODO())
	if err != nil{
		fmt.Println("Failed to load aws credentials")
		return
	}

	// create a new s3 client

	s3client := s3.NewFromConfig(cfg)

	// create s3 bucket

	bucketName := "bucketcreatewith-golang"

	_, err = s3client.CreateBucket(context.Background(), &s3.CreateBucketInput{
        Bucket: &bucketName,
    })
    if err != nil {
        fmt.Println("Error creating S3 bucket:", err)
        return
    }

	fmt.Println("S3 bucket:" , bucketName , "created successfully")

	fmt.Println("Do you want to get list of s3 buckets for your account ? Y/N ")

	var resp string
	fmt.Scanf("%s", &resp)

	if resp == "Y" {
		getList(cfg)
	} else{
		fmt.Println("Thank you for your response!")
	}
}


func getList(cfg config.Config) {
	
	client := s3.NewFromConfig(aws.Config{Region: "us-east-1"})

    result, err := client.ListBuckets(context.TODO(), nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("S3 Buckets list: ")
    for _, bucket := range result.Buckets {
        fmt.Println(*bucket.Name)
    }
}