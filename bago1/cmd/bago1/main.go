package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/balionis-arvydas/balionis-go/bago1/pkg/bago1"
	"log"
)

func main() {
	log.Printf("main: +")
	defer log.Printf("main: -")

	lambda.Start(bago1.HandleRequest)
}
