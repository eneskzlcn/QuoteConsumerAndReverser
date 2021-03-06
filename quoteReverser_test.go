package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIFetchQuotesSuccessfully(t *testing.T){
	quotes := make([]Quote,0)
	resp, err := http.Get(GetEnv("API_BASE_URL"))
	if err != nil {
		log.Println(err.Error())
	}
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body,&quotes); err != nil{
		log.Println("Can not parse body into quotes struct ")
	}
	for index, quote := range quotes {
		log.Println(fmt.Sprintf("%d-%s-%s",index,quote.Author,quote.Text))
	}
	quotesMap := AuthorsQuotesMap{}
	quotesMap.ConstructWithGivenSlice(quotes)
	quotesMap.PrintAsWantedJSON()
}
//test grouping by authors
func TestIGroupQuotesByAuthorsSuccessfully(t *testing.T){

	testData := []Quote{
		{
			Author: "Einstein",
			Text:   "Physics",
		},
		{
			Author: "Einstein",
			Text:   "Music",
		},
		{
			Author: "Einstein",
			Text:   "Physics",
		},
		{
			Author: "Einstein",
			Text:   "Silence",
		},
		{
			Author: "Edward",
			Text:   "Physics",
		},
		{
			Author: "Edward",
			Text:   "Music",
		},
	}
	quotes := AuthorsQuotesMap{}
	quotes.ConstructWithGivenSlice(testData)

	assert.Equalf(t,2,len(quotes["Edward"]),"Edward has 2 sentences if the grouping operation is done in truth")
	assert.Equalf(t,4,len(quotes["Einstein"]),"Einstein has 3 sentences if the grouping operation is done in truth")

}
//test reverse operation

func TestIReverseSentencesSuccessfully(t *testing.T){

	textToReverse := "Physics"
	reversedText := Reverse(textToReverse)
	assert.Equalf(t, "scisyhP",reversedText,"If the reverse operation done , Physics should be scisyhP")
}