[![Serverless](http://public.serverless.com/badges/v3.svg)](http://www.serverless.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/whithajess/unity)](https://goreportcard.com/report/github.com/whithajess/unity)
[![License](https://img.shields.io/github/license/whithajess/unity.svg)](LICENSE.md)

# Unity

A chatbot with the following goals in mind:

* Low cost to run, ideally no need for persistent running resources
* This essentially means we want AWS lambda + API gateway, and for only messages directed at the bot to hit our API endpoint
* Support for Slack - our chosen messaging platform
* Ideally also support for rich message formatting in slack.
* Easily extensible by others - it should be possible for other developers to add new functionality
* Access control - some features may need to be restricted to a subset of users.
* Friendly and helpful to interact with:
* Responses and interactions should never be rude towards users who are unfamiliar with the project or existence of the bot
* Format for humans not robots (prefer formatted tables, over raw JSON dumps for example)
* Help text should be included where possible to allow discovery of features
* Avoid over technical jargon. Clear and accurate is desirable, but a technically correct but not understandable message has missed the mark.
