package main

import "strings"

func badWordReplacement(body string) string {
	profanities := make(map[string]bool)
	profanities["kerfuffle"] = true
	profanities["sharbert"] = true
	profanities["fornax"] = true

	if len(body) <= 0 {
		return ""
	}

	words := strings.Split(body, " ")
	for i, word := range words {
		word = strings.ToLower(word)
		if profanities[word] {
			words[i] = "****"
		}
	}

	result := strings.Join(words, " ")
	return result
}
