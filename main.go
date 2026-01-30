package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type Operation struct {
	OpID       string `json:"op_id"`
	Type       string `json:"type"`
	Value      int    `json:"value"`
	OccurredAt string `json:"occurred_at"` // format "2006-01-02T15:04:05Z07:00" = time.RFC3339
}

func main() {
	// read json input
	data, err := os.ReadFile("input.json")
	if err != nil {
		log.Fatal(err)
	}

	var ops []Operation

	if err := json.Unmarshal(data, &ops); err != nil {
		log.Fatal(err)
	}

	dedupOps := make([]Operation, 0, len(ops))
	seen := make(map[string]struct{})

	// process loop to deduplicate input
	for _, v := range ops {
		if _, ok := seen[v.OpID]; ok {
			continue
		}

		seen[v.OpID] = struct{}{}
		dedupOps = append(dedupOps, v)
	}

	sort.Slice(dedupOps, func(i, j int) bool {
		// time index of element i
		ti, err := time.Parse(time.RFC3339, dedupOps[i].OccurredAt)
		if err != nil {
			log.Fatal(err)
		}

		// time index of element j
		tj, err := time.Parse(time.RFC3339, dedupOps[j].OccurredAt)
		if err != nil {
			log.Fatal(err)
		}

		return ti.Before(tj)
	})

	finalValue := 0

	// if operation is increment, we add the value
	// if operation is decrement, we subtract the value
	for _, op := range dedupOps {
		switch op.Type {
		case "increment":
			finalValue += op.Value
		case "decrement":
			finalValue -= op.Value
		}
	}

	result := map[string]int{
		"final_value": finalValue,
	}

	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println(string(out))
}
