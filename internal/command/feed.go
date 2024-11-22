package command

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "gator")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	xmlData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var rssFeed RSSFeed
	if err = xml.Unmarshal(xmlData, &rssFeed); err != nil {
		return nil, err
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i, item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i] = RSSItem{
			Title:       html.UnescapeString(item.Title),
			Description: html.UnescapeString(item.Description),
			Link:        item.Link,
			PubDate:     item.PubDate,
		}
	}
	return &rssFeed, nil

}
