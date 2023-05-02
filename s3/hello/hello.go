package main

// we are going to print list of buckets in our account

import(
	"fmt"
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

)

func main(){
	// Load the aws shared credentials first
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil{
		fmt.Printf("ARE YOU CONFIGURED YOUR AWS CRENDENTIALS? Here is the error: %v", err)
		return
	}

	s3client := s3.NewFromConfig(cfg)

	// define the count upto list the buckets
	count := 10

	fmt.Printf("Let's list upto %v buckets for your account \n", count)

	result, err := s3client.ListBuckets(context.TODO(), nil)
	if err != nil{
		fmt.Printf("Couldn't load any buckets, error : %v", err)
		return
	}

	if len(result.Buckets) == 0 {
		fmt.Println("No buckets are present in the account.")
	} else {
		if count > len(result.Buckets) {
			count = len(result.Buckets)
		}
		for _, bucket := range result.Buckets[:count]{
			fmt.Printf("\t%v\n", *bucket.Name)
		}
	}
}