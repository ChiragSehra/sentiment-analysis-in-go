package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cdipaolo/sentiment"
)

func main() {
	// Open the file
	csvfile, err := os.Open("data/IMDBDataset.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	// Parse the file
	r := csv.NewReader(csvfile)

	// Creating Sentiment Analysis Model
	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}
	var analysis *sentiment.Analysis

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		analysis = model.SentimentAnalysis(record[0], sentiment.English)
		var sentiment string
		if analysis.Score == 1 {
			sentiment = "positive"
		} else {
			sentiment = "negative"
		}
		fmt.Printf("Review: %s \n and Sentiment:%s\n", record[0], sentiment)

	}
}
