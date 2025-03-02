package main

import (
	"context"
	"log"
	"os"

	"github.com/FlowingSPDG/streamdeck"
	sdresolume "github.com/FlowingSPDG/streamdeck-resolume/Source/code"
)

func main() {
	ctx := context.Background()
	params, err := streamdeck.ParseRegistrationParams(os.Args)
	if err != nil {
		log.Fatalf("Failed to parse registration params: %v", err)
	}

	plugin := sdresolume.NewPlugin(streamdeck.NewClient(ctx, params))

	if err := plugin.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
