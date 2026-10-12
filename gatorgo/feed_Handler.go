package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func fetchFeed(context context.Context, feedURL string) (*config.RSSFeed, error) {

	// declaring variables
	var emptyFeed *config.RSSFeed
	var client http.Client
	var data config.RSSFeed

	// creating and doing the http.Request
	req, errRequest := http.NewRequestWithContext(context, "GET", feedURL, nil)
	if errRequest != nil {
		return emptyFeed, errRequest
	}
	req.Header.Set("user-agent", "gator")

	res, errResponse := client.Do(req)
	if errResponse != nil {
		return emptyFeed, errResponse
	}
	defer res.Body.Close()

	// reading the data from the request
	resData, errReading := io.ReadAll(res.Body)
	if errReading != nil {
		return emptyFeed, errReading
	}

	errUnmarshal := xml.Unmarshal(resData, &data)
	if errUnmarshal != nil {
		return emptyFeed, errUnmarshal
	}

	// unescaping html
	data.Channel.Title = html.UnescapeString(data.Channel.Title)
	data.Channel.Description = html.UnescapeString(data.Channel.Description)
	for _, item := range data.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &data, nil
}
