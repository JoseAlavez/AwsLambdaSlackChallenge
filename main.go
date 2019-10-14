package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Challenge struct {
	Challenge string `json:"challenge"`
}

func main() {
	lambda.Start(func(request Challenge) (Challenge, error) {
		return request, nil
	})
}
