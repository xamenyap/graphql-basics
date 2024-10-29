package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ProductQuery struct {
	Query     string `json:"query"`
	Variables struct {
		ID string `json:"id"`
	} `json:"variables"`
}

func main() {
	client := &http.Client{}
	query := ProductQuery{
		Query: `query getProduct($id: ID!) {
			getProduct(id: $id) {
				id
				name
				reviews {
					content
					rating
				}
			}
		}`,
	}
	query.Variables.ID = "1"

	jsonData, err := json.Marshal(query)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/query", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Printf("Response: %+v\n", result)
}
