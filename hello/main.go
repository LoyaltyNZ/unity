package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/awslabs/aws-lambda-go-api-proxy/core" // For converting APIGateway events to standard go http
	"github.com/nlopes/slack" // Slack api library
)

// Response - Basic json response
type Response struct {
	Message string `json:"message"`
}

// Handler - Handles Requests (Returns Echoed Message as Response)
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Build the APIGatewayProxy event into a http.request
	accessor := core.RequestAccessor{}
	r, nil := accessor.ProxyEventToHTTPRequest(request)

	// Parse the SlashCommand from Slack
	slashCommand, err := slack.SlashCommandParse(r)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Internal Server Error", StatusCode: 500}, nil
	}

	// Check the Verification token given from Slack matches our Bot Token
	if !slashCommand.ValidateToken(os.Getenv("VERIFICATION_TOKEN")) {
		return events.APIGatewayProxyResponse{Body: "Unauthorised", StatusCode: 401}, nil
	}

	// Made it through.
	return events.APIGatewayProxyResponse{Body: slashCommand.Text, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
