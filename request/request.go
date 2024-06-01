package request

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slyfluffy/bible-bot/scripture"
	"strings"
)

func cleanupText(scripture scripture.Scripture) {
	replacer := strings.NewReplacer("\n", "")
	scripture.Text = replacer.Replace(scripture.Text)
}

func RandomVerse() scripture.Scripture {
	fmt.Println("Getting the random verse from bible-api.com...")
	response, err := http.Get("https://bible-api.com/?translation=kjv&random=verse")
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := io.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println("API Response as a string:\n" + bodyString)

	var scripture scripture.Scripture
	json.Unmarshal(bodyBytes, &scripture)
	fmt.Printf("API Response as a struct %+v\n", scripture)

	cleanupText(scripture)

	return scripture
}


