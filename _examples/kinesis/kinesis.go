package main

import (
	"os"
	"time"

	"github.com/yewno/log"
	"github.com/yewno/log/handlers/kinesis"
	"github.com/yewno/log/handlers/multi"
	"github.com/yewno/log/handlers/text"
)

func main() {
	log.SetHandler(multi.New(
		text.New(os.Stderr),
		kinesis.New("logs"),
	))

	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "tobi",
	})

	for range time.Tick(time.Millisecond * 100) {
		ctx.Info("upload")
		ctx.Info("upload complete")
	}
}
