package image

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	// "github.com/disintegration/imaging"

	"github.com/disintegration/imaging"
	"golang.org/x/net/html"
)

// Resize :
func Resize(src string) {
	// srcImg, err := LoadImage(src)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	srcImg, err := imaging.Open(src)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	// Resize srcImage to width = 800px preserving the aspect ratio.
	// dst := imaging.Resize(srcImg, 800, 0, imaging.Lanczos)

	if srcImg.Bounds().Max.X > 900 {
		srcImg = imaging.Resize(srcImg, 900, 0, imaging.Lanczos)
	}
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, srcImg); err != nil {
		log.Fatal(err)
	}

	fmt.Println("image buf => ", buf)

	// err = imaging.Save(srcImg, "/home/ray/img_test/sz-airport-resized.jpg")
	// if err != nil {
	// 	log.Fatalf("failed to save image: %v", err)
	// }

	elasped := time.Now().Sub(start)

	fmt.Printf("time elasped: %f\n", elasped.Seconds())
}

// LoadImage :
func LoadImage(filePath string) (image.Image, error) {
	// f, err := os.Open(filePath)
	// if err != nil {
	// 	return nil, err
	// }
	// defer f.Close()

	// img, _, err := image.Decode(f)
	// if err != nil {
	// 	return nil, err
	// }

	img, err := imaging.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return img, nil
}

// ResizeImageByURL :
func ResizeImageByURL(url string, width int) string {
	if strings.Contains(url, "https://ahezime.com/file") || strings.Contains(url, "http://localhost") {
		log.Println("resizing image...")
		url = url + "?r=true&w=" + strconv.Itoa(width)
	}
	return url
}

// ResizeImagesByURLInHTML :
func ResizeImagesByURLInHTML(content string, width int) (string, error) {
	if width <= 0 {
		width = 1000
	}
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		log.Error(err)
		return "", err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for i, v := range n.Attr {
				if v.Key == "src" {
					if strings.Contains(v.Val, "https://ahezime.com/file") || strings.Contains(v.Val, "http://localhost") {
						v.Val = v.Val + "?r=true&w=" + strconv.Itoa(width)
						n.Attr[i].Val = v.Val // !important if you wanna to change the value
						log.Println("resizing image => ", v.Val)
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	buf := bytes.NewBuffer([]byte{})
	if err = html.Render(buf, doc); err != nil {
		return "", err
	}
	content = buf.String()

	return content, nil
}
