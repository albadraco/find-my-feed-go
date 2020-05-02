package parser

import (
	"context"
	"github.com/mmcdole/gofeed"
)

// Header headers to insert into call
type Header struct {
	Name  string
	Value string
}


func (f *Parser) ParseURLWithHeaders(feedURL string, context.Background(), headers ...Header) {
	client := f.httpClient()

	req, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if h := req.Header.Get("User-Agent"); h == "" {
		req.Header.Set("User-Agent", "Gofeed/1.0")
	}
	if headers != nil {
		for _, header := range headers {
			req.Header.Set(header.Name, header.Value)
		}
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer func() {
			ce := resp.Body.Close()
			if ce != nil {
				err = ce
			}
		}()
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
		}
	}

	return f.Parse(resp.Body)
}