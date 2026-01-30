# Backend Engineer Take-Home Assignment

## Code Test 0 - Fundamentals

This program processes a list of counter operations provided as JSON input.

### Features
- Reads operations from `input.json`
- Deduplicates operations by `op_id`
- Sorts operations by `occurred_at` (ascending)
- Applies increment/decrement sequentially
- Outputs deterministic final value

### How to Run

```bash
go run main.go