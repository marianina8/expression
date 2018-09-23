# Expression
Expression (emotion) analysis on faces found in video files using Microsofts Emotion API.

## Installing

Before you begin, you'll of course need the Go programming language installed.

Install:
[GoCV](https://gocv.io/)

Gain access to a [Microsofts Emotion API](https://azure.microsoft.com/en-us/services/cognitive-services/emotion/) and grab an API key and host from their Azure Portal.

Export `emotion_key` and `emotion_host`:

`export emotion_key={API Key}`

`export emotion_host={Host}`

1. Clone this repo.
2. Run `dep ensure`
3. Build the application with `go build`.
5. Run the program according the usage below.

## Usage

`expression -video {video file path}`

video is a required parameter.

## Future To Do

-[ ] Combine with body posture detection to evaluate based on body language general sentiment of human
-[ ] Combine with NLP to see if their words match their facial expression and general body posture, or if there's a disconnect.
