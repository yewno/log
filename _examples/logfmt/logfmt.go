package main

import (
	"errors"
	"os"
	"time"

	"github.com/yewno/log"
	"github.com/yewno/log/handlers/logfmt"
)

func main() {
	log.SetHandler(logfmt.New(os.Stderr))

	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "tobi",
	})

	for range time.Tick(time.Millisecond * 200) {
		ctx.Info("upload")
		ctx.Info("upload complete")
		ctx.Warn("upload retry")
		ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	}
}
