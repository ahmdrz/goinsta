package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TheForgotten69/goinsta/v2"
)

func main() {
	insta := goinsta.New(
		os.Getenv("INSTAGRAM_USERNAME"),
		os.Getenv("INSTAGRAM_PASSWORD"),
	)
	if err := insta.Login(); err != nil {
		log.Fatal(err)
	}

	defer insta.Logout()

	log.Println("Download random photo")
	var client http.Client
	request, err := http.NewRequest("GET", "https://picsum.photos/800/800", nil)
	if err != nil {
		log.Fatal(err)
	}
	thumbnail, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer thumbnail.Body.Close()

	log.Println("Download random video")
	request, err = http.NewRequest("GET", "https://www.w3toys.com/download.php?url=https%3A%2F%2Fscontent-cdt1-1.cdninstagram.com%2Fv%2Ft50.2886-16%2F142537180_690913001575458_6043626278285741826_n.mp4%3F_nc_ht%3Dscontent-cdt1-1.cdninstagram.com%26_nc_cat%3D106%26_nc_ohc%3DtngasMRSAgsAX-kTy42%26oe%3D60197438%26oh%3De81483acbdfa4570a433c74544c6edc0%26dl%3D1", nil)
	if err != nil {
		log.Fatal(err)
	}
	video, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer video.Body.Close()

	postedVideo, err := insta.UploadVideo(video.Body, "awesomeVID", "awesome! :)", thumbnail.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Success upload video %s", postedVideo.ID)
}
