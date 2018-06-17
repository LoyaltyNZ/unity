[![Serverless](http://public.serverless.com/badges/v3.svg)](http://www.serverless.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/whithajess/unity)](https://goreportcard.com/report/github.com/whithajess/unity)
[![License](https://img.shields.io/github/license/whithajess/unity.svg)](LICENSE.md)

***Current Status: Will respond to messages with "Yes, Slackbot."***

# Unity

An AWS SlackBot with the following goals in mind:

* Low cost to run, ideally no need for persistent running resources
* Easily extensible by others - it should be possible for other developers to add new functionality
* Access control - some features may need to be restricted to a subset of users.
* Friendly and helpful to interact with:
    * Responses and interactions should never be rude towards users who are unfamiliar with the project or existence of the bot
    * Format for humans not robots (prefer formatted tables, over raw JSON dumps for example)
    * Help text should be included where possible to allow discovery of features
    * Avoid over technical jargon. Clear and accurate is desirable, a technically correct but not understandable message has missed the mark.

#### Technologies

* [AWS Lambda](https://aws.amazon.com/lambda/)
* [AWS API Gateway](https://aws.amazon.com/api-gateway/)
* [Slack](https://slack.com/)

#### Prerequisites

* [AWS Account](https://portal.aws.amazon.com/gp/aws/developer/registration/index.html)
* [Serverless](https://serverless.com/framework/docs/getting-started/)
* [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/installing.html)

#### Setup

***1.*** Login to https://api.slack.com/apps and Create New App
    * The App name you use will be what users need to mention to communicate with Unity bot i.e @unity
    * We are using the Events api so you will need to enable events
    * The App will need to be installed for you to get an OAuth token.

***2.*** After creating the application in Slack we need to set tokens for the App to use:
    * These are region specific if you haven't set `AWS_PROFILE` you may want to use the flag for setting the region `--region ap-southeast-2`
```bash
  # The verification token
  # can be found under Basic Information in the App on https://api.slack.com/apps
  # This gives us the ability to check the messages sent to the App are actually coming from Slack
  aws ssm put-parameter --name unityBotVerificationToken --type String --value SecretToken

  # The OAuth token
  # can be found under OAuth & Permissions in the App on https://api.slack.com/apps
  # This gives the App the ability to send messages to Slack
  aws ssm put-parameter --name unityBotOAuthAccessToken --type String --value xoxb-SecretToken
```

***3.*** After this is all set you will need to deploy the application
```bash
  make
  sls deploy
```

***4.*** The output of deployment will have an `endpoint`, you need to set this on the App on https://api.slack.com/apps under Event Subscriptions Request URL.

***5.*** After this you should subscribe the App to Bot Events (Still under Event Subscriptions)
    * I suggest for this you use:
        * `app_mention` - Subscribe to only the message events that mention your app or bot
        * `message.app_home` - A user sent a message to your Slack app
        * `message.im` - A message was posted in a direct message channel


<!-- TODO: Write Instructions for Collaboration -->
