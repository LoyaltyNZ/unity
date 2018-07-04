package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/nlopes/slack"
	"github.com/whithajess/slack/slackevents" // Need to use my fork until https://github.com/nlopes/slack/pull/326 is merged
)

// Response - Basic json response
type Response struct {
	Message string `json:"message"`
}

// Initialise Slack API with the Bot Token
var api = slack.New(os.Getenv("OAUTH_ACCESS_TOKEN"))

// Listener - Listens for Slack Events and sends legitimate ones to Amazon Lex for natural language parsing
// based on the results Lex will call another Lambda
func Listener(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	eventsAPIEvent, error := slackevents.ParseEvent(json.RawMessage(request.Body), slackevents.OptionVerifyToken(slackevents.TokenComparator{os.Getenv("VERIFICATION_TOKEN")}))

	if error != nil {
		// TODO: Proper responses for Unauthorised
		fmt.Println("Error:", error)
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
						// TODO: This is where I am going to integrate with Amazon Lex
						// Need to move this into a shared method
						api.PostMessage(ev.Channel, "Yes, slackbot.", postParams)
					}
				}
			case *slackevents.AppMentionEvent:
				{
					// TODO: This is where I am going to integrate with Amazon Lex
					// Need to move this into a shared method
					api.PostMessage(ev.Channel, "Yes, slackbot.", postParams)
				}
			default:
				{
					fmt.Println("Bad Inner Event Type:", event)
					return events.APIGatewayProxyResponse{Body: "Bad Request", StatusCode: 400}, nil
				}
			}
			return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
		}
	default:
		{
			fmt.Println("Bad Outer Event Type:", event)
			return events.APIGatewayProxyResponse{Body: "Bad Request", StatusCode: 400}, nil
		}
	}
}

func main() {
	lambda.Start(Listener)
}
