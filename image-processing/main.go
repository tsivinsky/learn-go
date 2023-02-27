package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path"
	"strings"
)

var styleToFunc = map[string]func(m image.Image) image.Image{
	"bw": func(m image.Image) image.Image {
		return toBW(m)
	},
	"resize": func(m image.Image) image.Image {
		// random values lol
		x1 := 720
		y1 := 80
		x2 := 1230
		y2 := 600
		return resize(m, image.Rect(x1, y1, x2, y2))
	},
}

var (
	style string
	out   string
)

func main() {
	flag.StringVar(&style, "style", "", "-style bw")
	flag.StringVar(&out, "out", "./image.jpeg", "-out ./image.jpeg")

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No path to .jpeg image provided")
		os.Exit(1)
	}

	var fp = flag.Args()[0]

	f, err := os.OpenFile(fp, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	res := m

	cases := strings.Split(style, ",")
	for _, c := range cases {
		if f, ok := styleToFunc[c]; ok {
			res = f(res)
		}
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outPath := path.Join(cwd, out)

	nf, err := os.OpenFile(outPath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer nf.Close()

	err = jpeg.Encode(nf, res, &jpeg.Options{Quality: 90})
	if err != nil {
		log.Fatal(err)
	}
}
