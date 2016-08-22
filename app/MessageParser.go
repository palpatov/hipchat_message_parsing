package app

import (
	"io"

	"sync"

	"context"
	"time"

	"github.com/palpatov/hipchat_message_parsing/domain"
	"github.com/palpatov/hipchat_message_parsing/hcjson"
	"github.com/palpatov/hipchat_message_parsing/parsers"
)

const timeoutInSeconds = 1

//
// Parse the actual parsing logic goes here
//
func Parse(m string, w io.Writer) {
	ctx := context.Background()
	var cancel context.CancelFunc
	var result domain.ParseResult

	var wg sync.WaitGroup
	wg.Add(2)

	//
	// splitting this in two groutines due to the fact that ParseUrls parser fetches
	// remote resource (page).
	//
	go func() {
		defer wg.Done()

		//
		// context will timeout in timeoutInSeconds set above.
		//
		ctx, cancel = context.WithTimeout(ctx, time.Duration(timeoutInSeconds*time.Second))
		defer cancel()

		li := parsers.ParseUrls(ctx, m)
		if li != nil {
			result.Links = li.URLs
		}
	}()

	//
	// below goroutine is all local parsing, so no context is being passed into
	//
	go func() {
		defer wg.Done()
		me := parsers.ParseMentions(m)
		if me != nil {
			result.Mentions = me.MentionNames
		}
		em := parsers.ParseEmoticons(m)
		if em != nil {
			result.Emoticons = em.Emoticons
		}
	}()

	wg.Wait()
	hcjson.FormatJSONToResponse(result, w)
}
