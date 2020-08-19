package genderzine_test

import (
	"context"
	"log"

	"github.com/alexeyco/genderzine"
)

func Example() {
	client := genderzine.New()

	res, err := client.Check(context.Background(), "Alice")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Name", res.Name)               // Name Alice
	log.Println("Gender", res.Gender)           // Gender female
	log.Println("Probability", res.Probability) // Probability 0.98
	log.Println("Count", res.Count)             // Count 68971

	info := client.Info()

	log.Println("Limit", info.Limit)         // Limit 1000
	log.Println("Remaining", info.Remaining) // Remaining 984
	log.Println("Reset", info.Reset)         // Reset 6m10s
}
