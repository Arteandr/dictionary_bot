package commands

import (
	"dictionary_bot/pkg/dictionary"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"strings"
	"sync"
)

func (h *CMDHandler) OnText(ctx tele.Context) error {
	words := strings.Split(ctx.Message().Text, "\n")

	if len(words) == 1 {
		if words[0] == "/start" {
			return nil
		}
	}

	str := ""
	var mut sync.Mutex

	wg := &sync.WaitGroup{}

	wg.Add(len(words))
	for _, word := range words {
		go func(w string) {
			defer wg.Done()
			r, err := dictionary.GetWordInfo(w)
			if err != nil {
				mut.Lock()
				str += fmt.Sprintf("%s: error\n", w)
				mut.Unlock()
				return
			}

			mut.Lock()
			str += fmt.Sprintf("<b>%s</b>: \n", w)
			for _, meaning := range r.Meanings {
				str += fmt.Sprintf("\t <u>%s</u> - %s\n", meaning.PartOfSpeech, meaning.Definitions[0].Definition)
			}
			mut.Unlock()
		}(strings.ToLower(strings.TrimSpace(word)))
	}

	wg.Wait()

	return ctx.Send(str, tele.ModeHTML)
}
