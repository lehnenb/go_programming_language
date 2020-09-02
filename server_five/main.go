// Server2 is a minimal "echo" and counter server
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(cycles int64, out io.Writer) {
	var palette = []color.Color{color.White, color.Black}

	const (
		whiteIndex = 0 // first color in palette
		blackIndex = 1 // next color in palette
	)

	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func handler(w http.ResponseWriter, r *http.Request) {
	cyclesParam, cyclesOk := r.URL.Query()["cycles"]

	if !cyclesOk || len(cyclesParam) == 0 {
		log.Println("Url Param 'cycles' is missing")
		return
	}

	fmt.Printf("Request - %s: %s\n", r.URL.Path, time.Now().Format(time.RFC850))
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Printf("Header[%q] = %q\n", k, v)
	}

	fmt.Printf("Host = %q\n", r.Host)
	fmt.Printf("RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Printf("URL.Path = %q\n", r.URL.Path)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Printf("Form[%q] = %q\n", k, v)
	}

	cycles, err := strconv.ParseInt(cyclesParam[0], 10, 64)

	if err != nil {
		log.Println("Invalid cycles param")
	}

	lissajous(cycles, w)
}
