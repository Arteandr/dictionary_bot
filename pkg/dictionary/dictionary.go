package dictionary

import (
	"encoding/json"
	"net/http"
	"time"
)

const API_URL = "https://api.dictionaryapi.dev/api/v2/entries/en/"

func GetWordInfo(word string) (Response, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(API_URL + word)
	if err != nil {
		return Response{}, err
	}
	defer r.Body.Close()
	var unmarshalled []Response

	if err := json.NewDecoder(r.Body).Decode(&unmarshalled); err != nil {
		return Response{}, err
	}

	return unmarshalled[0], nil
}
