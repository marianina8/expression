package azure

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// NewClient is the default constructor for creating a vision client using the api key.
func NewClient(emotionAPI, emotionKey string) (*Client, error) {
	if emotionKey == "" {
		return nil, errors.New("missing emotion_key")
	}
	if emotionAPI == "" {
		return nil, errors.New("missing emotion_host")
	}
	return &Client{
		API:    emotionAPI,
		Key:    emotionKey,
		client: *http.DefaultClient,
	}, nil
}

// FaceAnalysis runs sentiment analysis on an image
func (c *Client) FaceAnalysis(body io.Reader) EmotionData {
	result := []EmotionData{}
	contentType := "application/octet-stream"

	var URL *url.URL
	URL, err := url.Parse(c.API)
	if err != nil {
		check("url parsing", err)
	}
	URL.Path += "/face/v1.0/detect"
	parameters := url.Values{}
	parameters.Add("returnFaceAttributes", "emotion")
	URL.RawQuery = parameters.Encode()

	b, err := ioutil.ReadAll(body)
	if err != nil {
		check("reading file", err)
	}

	req, err := http.NewRequest(http.MethodPost, URL.String(), bytes.NewReader(b))
	if err != nil {
		check("create new request", err)
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", c.Key)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Content-Length", strconv.Itoa(len(b)))

	resp, err := c.client.Do(req)
	if err != nil {
		check("request", err)
	}
	defer resp.Body.Close()

	var errResp struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	var resBody io.Reader
	peeker := bufio.NewReader(resp.Body)

	// when emotion API succeeds, a JSON array is returned instead of an object, so we need to peek
	if head, err := peeker.Peek(1); err != nil {
		check("peeking", err)
	} else if head[0] == '{' {
		var body bytes.Buffer
		err = json.NewDecoder(io.TeeReader(peeker, &body)).Decode(&errResp)
		if err != nil {
			check("decoding error response", err)
		}
		resBody = &body
	} else {
		resBody = peeker
	}
	err = fmt.Errorf("status: %s:%s", errResp.Code, errResp.Message)
	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusTooManyRequests:
		return EmotionData{}
	case http.StatusUnauthorized:
		check("unauthorized", err)
	case http.StatusBadRequest:
		check("bad request", err)
	default:
		check("unrecognized", err)
	}
	json.NewDecoder(resBody).Decode(&result)
	return result[0]
}

func check(msg string, e error) {
	if e != nil {
		panic(fmt.Errorf("%s: %s", msg, e.Error()))
	}
}
