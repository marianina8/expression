package azure

import "net/http"

// Client is the client used to make API calls to the Azure vision service(s)
type Client struct {
	Key    string
	client http.Client
	API    string
}

// EmotionData contains face details from MS Emotion API
type EmotionData struct {
	FaceID         string         `json:"faceId"`
	FaceRectangle  FaceRectangle  `json:"faceRectangle"`
	FaceLandmarks  interface{}    `json:"faceLandmarks"`
	FaceAttributes FaceAttributes `json:"faceAttributes"`
}

// Dominant returns string description of strongest emotion detected
func (e EmotionData) Dominant() string {
	maxScore := 0.0
	emotion := "unknown"
	if e.FaceAttributes.Emotion.Anger > maxScore {
		maxScore = e.FaceAttributes.Emotion.Anger
		emotion = "anger"
	}
	if e.FaceAttributes.Emotion.Contempt > maxScore {
		maxScore = e.FaceAttributes.Emotion.Contempt
		emotion = "contempt"
	}
	if e.FaceAttributes.Emotion.Disgust > maxScore {
		maxScore = e.FaceAttributes.Emotion.Disgust
		emotion = "disgust"
	}
	if e.FaceAttributes.Emotion.Fear > maxScore {
		maxScore = e.FaceAttributes.Emotion.Fear
		emotion = "fear"
	}
	if e.FaceAttributes.Emotion.Happiness > maxScore {
		maxScore = e.FaceAttributes.Emotion.Happiness
		emotion = "happiness"
	}
	if e.FaceAttributes.Emotion.Neutral > maxScore {
		maxScore = e.FaceAttributes.Emotion.Neutral
		emotion = "neutral"
	}
	if e.FaceAttributes.Emotion.Sadness > maxScore {
		maxScore = e.FaceAttributes.Emotion.Sadness
		emotion = "sadness"
	}
	if e.FaceAttributes.Emotion.Surprise > maxScore {
		maxScore = e.FaceAttributes.Emotion.Surprise
		emotion = "surprise"
	}
	return emotion
}

// FaceAttributes contains face attributes and emotion scores
type FaceAttributes struct {
	Age        float64 `json:"age"`
	Gender     string  `json:"gender"`
	Smile      float64 `json:"smile"`
	FacialHair struct {
		Moustache float64 `json:"moustache"`
		Beard     float64 `json:"beard"`
		Sideburns float64 `json:"sideburns"`
	} `json:"facialHair"`
	Glasses  string `json:"glasses"`
	HeadPose struct {
		Roll  float64 `json:"roll"`
		Yaw   int     `json:"yaw"`
		Pitch int     `json:"pitch"`
	} `json:"headPose"`
	Emotion Scores `json:"emotion"`
	Hair    struct {
		Bald      float64 `json:"bald"`
		Invisible bool    `json:"invisible"`
		HairColor []struct {
			Color      string  `json:"color"`
			Confidence float64 `json:"confidence"`
		} `json:"hairColor"`
	} `json:"hair"`
	Makeup struct {
		EyeMakeup bool `json:"eyeMakeup"`
		LipMakeup bool `json:"lipMakeup"`
	} `json:"makeup"`
	Occlusion struct {
		ForeheadOccluded bool `json:"foreheadOccluded"`
		EyeOccluded      bool `json:"eyeOccluded"`
		MouthOccluded    bool `json:"mouthOccluded"`
	} `json:"occlusion"`
	Accessories []struct {
		Type       string  `json:"type"`
		Confidence float64 `json:"confidence,omitempty"`
	} `json:"accessories"`
	Blur struct {
		BlurLevel string  `json:"blurLevel"`
		Value     float64 `json:"value"`
	} `json:"blur"`
	Exposure struct {
		ExposureLevel string  `json:"exposureLevel"`
		Value         float64 `json:"value"`
	} `json:"exposure"`
	Noise struct {
		NoiseLevel string  `json:"noiseLevel"`
		Value      float64 `json:"value"`
	} `json:"noise"`
}

// FaceRectangle defines the rectangle around analyzed face
type FaceRectangle struct {
	Left   int `json:"left"`
	Top    int `json:"top"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// Scores contains face emotion scores
type Scores struct {
	Anger     float64 `json:"anger"`
	Contempt  float64 `json:"contempt"`
	Disgust   float64 `json:"disgust"`
	Fear      float64 `json:"fear"`
	Happiness float64 `json:"happiness"`
	Neutral   float64 `json:"neutral"`
	Sadness   float64 `json:"sadness"`
	Surprise  float64 `json:"surprise"`
}
