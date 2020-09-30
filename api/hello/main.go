package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{"Content-Type": "text/plain", "Access-Control-Allow-Origin": "michaelshiplet.com"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body: "Yes! I'm here! But it's time to go!",
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}

