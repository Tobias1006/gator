package main

import (
	"context"
	"fmt"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func handlerAggregator(s *config.State, cmd config.Cmd) error {
	ctxt := context.Background()
	feedUrl := "https://www.wagslane.dev/index.xml"
	feed, errFetch := fetchFeed(ctxt, feedUrl)
	if errFetch != nil {
		return errFetch
	}
	fmt.Print(feed)
	return nil
}
