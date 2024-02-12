package main

import (
	"log"

	"github.com/notedit/rtmp/format/rtmp"
)

func main() {
	cli := rtmp.NewClient()
	url := "rtmp://a.rtmp.youtube.com/live2?key=kvz4-m3a4-hu83-0m3d-dgz0"
	conn, c, err := cli.Dial(url, rtmp.PrepareWriting)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
}
