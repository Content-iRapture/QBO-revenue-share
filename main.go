package main

import (
	"net/http"
	"fmt"
	"os"
	"log"
	"io"
	"net/url"
)

func main() {
	// https://sandbox-quickbooks.api.intuit.com/v3/company/<realmId>/query?query=<select_statement>
	realmId := os.Getenv("REALM_ID")
	query := "SELECT * FROM Invoice MAXRESULTS 1"
	query = url.QueryEscape(query)
	
	URL := fmt.Sprintf(
		"https://sandbox-quickbooks.api.intuit.com/v3/company/%s/query?query=%s",
		realmId,
		query,
	)
	
	res, err := http.Get(URL)
	
	if (err != nil) {
		log.Fatalf("Error on GET: %v", err)	
	}

	fmt.Printf("Code: %v - %s\n", res.StatusCode, res.Status)
	
	body, _ := io.ReadAll(res.Body)
	fmt.Printf("Body: %s\n", body)
}
