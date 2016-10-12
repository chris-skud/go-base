package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/chris-skud/go-base/correlation"
	"github.com/chris-skud/go-base/foo"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

func main() {
	// Initialize context
	ctx := context.Background()

	// Register handler wrapping domain handler in correlation middleware
	http.Handle("/foo", correlation.InitID(ctx, foo.New().AddHandler()))

	// Start server
	port := "9030"
	log.Infof("listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
