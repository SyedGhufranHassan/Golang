package main

import (
	"fmt"
	"log"

	"github.com/datumbrain/otters"
)

func main() {
	var totaltransactionamount float64
	var avgDeal float64

	// Load CSV
	df, err := otters.ReadCSV("contributions.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Basic overview
	fmt.Println(df.Describe())
	fmt.Println(df.Shape())
	fmt.Println(df.Columns())
	fmt.Println(df)

	// Aggregations
	totaltransactionamount, _ = df.Sum("transactionAmount")
	fmt.Printf("Total Transaction Amount: %.2f\n", totaltransactionamount)

	avgDeal, _ = df.Mean("transactionAmount")
	fmt.Printf("Average Transaction Amount: %.2f\n", avgDeal)

	// Multiple filters
	highEarners := df.
		Filter("transactionAmount", "<", 11).
		Filter("transactionSubTypeDesc", "==", "Itemized Monetary").
		Filter("transactionSource", "==", "Individual")

	fmt.Println("Filtered High Earners:")
	fmt.Println(highEarners)

	// Select specific columns
	summary := df.Select(
		"transactionAmount",
		"sourceName",
		"employerName",
		"occupation",
		"sourceAddress",
		"transactionSource",
		"reportName",
		"transactionSubTypeDesc",
	)

	fmt.Println("Summary Table:")
	fmt.Println(summary)

	// Head
	fmt.Println("Head of the DataFrame:")
	fmt.Println(summary.Head(5))

	// Sort by transaction amount
	fmt.Println("Sort on Transaction Amount:")
	top := summary.Sort("transactionAmount", false)
	fmt.Println(top)

	// Multi-column sort
	fmt.Println("Multi-col sort on transaction amount and transaction source:")
	ranked := df.SortBy(
		[]string{"transactionSource", "transactionAmount"},
		[]bool{true, false},
	)
	fmt.Println(ranked)

	// Drop unnecessary columns
	essential := df.Drop("guid", "filerRegistrationGuid", "totalRows", "transactionCategory")
	fmt.Println("Essential columns only:")
	fmt.Println(essential)

	err = df.WriteCSV("clean.csv")
	if err != nil {
		log.Fatal(err)
	}
}
