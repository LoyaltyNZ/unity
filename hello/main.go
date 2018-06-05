package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/nlopes/slack"
	"github.com/whithajess/slack/slackevents"
)

// Response - Basic json response
type Response struct {
	Message string `json:"message"`
}

// Initialise Slack API with the Bot Token
var api = slack.New(os.Getenv("OAUTH_ACCESS_TOKEN"))

// Handler - Handles Requests (Returns Echoed Message as Response)
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("request:", request)
	eventsAPIEvent, error := slackevents.ParseEvent(json.RawMessage(request.Body), slackevents.OptionVerifyToken(slackevents.TokenComparator{os.Getenv("VERIFICATION_TOKEN")}))

	if error != nil {
		// TODO: Proper responses for Unauthorised
		fmt.Println("Error:", error) // TODO: Worth a PR to nlopes/slack to change the error from "No" I mean come on
		return events.APIGatewayProxyResponse{Body: "Internal Server Error", StatusCode: 500}, nil
	}

	switch event := eventsAPIEvent.Data.(type) {
	// In the case of a events Verification always respond with the Challenge
	// Makes for easy setup
	case *slackevents.EventsAPIURLVerificationEvent:
		{
			return events.APIGatewayProxyResponse{Body: event.Challenge, StatusCode: 200}, nil
		}
	case *slackevents.EventsAPICallbackEvent:
		{
			postParams := slack.PostMessageParameters{}
			innerEvent := eventsAPIEvent.InnerEvent
			switch ev := innerEvent.Data.(type) {
			case *slackevents.MessageEvent:
				{
					// Don't do anything if Unity is the one sending the message
					// We don't want an infinite recursion situation.
					if ev.Username != "Unity" {
				      // TODO: This is where to add the bot logic
					  api.PostMessage(ev.Channel, "Yes, hello.", postParams)
					}
				}
			}
			return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
		}
	default:
		{
			fmt.Println("Bad Request Event:", event)
			return events.APIGatewayProxyResponse{Body: "Bad Request", StatusCode: 400}, nil
		}
	}
}

func main() {
	lambda.Start(Handler)
}
