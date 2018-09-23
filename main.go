package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/marianina8/expression/azure"
	"gocv.io/x/gocv"
)

func check(msg string, e error) {
	if e != nil {
		panic(fmt.Errorf("%s: %s", msg, e.Error()))
	}
}

func main() {
	video := flag.String("video", "", "video for emotion analysis")
	flag.Parse()
	if *video == "" {
		flag.Usage()
		return
	}
	key := os.Getenv("emotion_key")
	host := os.Getenv("emotion_host")
	azure, err := azure.NewClient(host, key)
	check("new azure client", err)
	displayVideo(azure, *video)
}

func displayVideo(client *azure.Client, filename string) {
	stream, err := gocv.VideoCaptureFile(filename)
	if err != nil {
		check("capture video file", err)
	}
	defer stream.Close()

	// open display window
	window := gocv.NewWindow("Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()
	framenum := 0
	black := color.RGBA{0, 0, 0, 0}
	dominantEmotion := ""
	for {
		if ok := stream.Read(&img); !ok {
			return
		}
		if img.Empty() {
			continue
		}
		framenum++

		b, err := gocv.IMEncode(".jpg", img)
		check("gocv image encode", err)

		// check frame ever 120 frames / 1 sec
		if framenum%30 == 0 {
			emotionData := client.FaceAnalysis(bytes.NewReader(b))
			if (emotionData.FaceRectangle != azure.FaceRectangle{}) {
				dominantEmotion = emotionData.Dominant()
			}
		}
		gocv.PutText(&img, dominantEmotion, image.Pt(40, 200), gocv.FontHersheyPlain, 3, black, 2)

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
