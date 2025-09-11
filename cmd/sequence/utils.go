package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func convertNumbers(input interface{}) interface{} {
	switch v := input.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			result[key] = convertNumbers(value)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(v))
		for i, value := range v {
			result[i] = convertNumbers(value)
		}
		return result
	case json.Number:
		if strings.Contains(string(v), ".") {
			f, _ := v.Float64()
			return f
		}
		n, _ := v.Int64()
		if n == 0 {
			return int64(0) // Ensure 0 is converted to int64
		}
		return n
	case float64:
		if v == float64(int64(v)) {
			return int64(v) // Convert whole numbers to int64
		}
		return v
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		if v <= math.MaxInt64 {
			return int64(v)
		}
		return v
	case string:
		// Try to parse string as number
		if n, err := json.Number(v).Int64(); err == nil {
			return n
		}
		if f, err := json.Number(v).Float64(); err == nil {
			if f == float64(int64(f)) {
				return int64(f)
			}
			return f
		}
		return v
	default:
		return v
	}
}

// readStdin reads all input from stdin until EOF
func readStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	var input []byte

	for {
		chunk, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read from stdin: %w", err)
		}
		input = append(input, chunk)
	}

	return string(input), nil
}

// fromPosOrStdin gets the value from either the positional argument or stdin
func fromPosOrStdin(args []string, pos int) (string, error) {
	if len(args) > pos {
		return args[pos], nil
	}

	input, err := readStdin()
	if err != nil {
		return "", err
	}

	if input == "" {
		return "", fmt.Errorf("no input provided")
	}

	return input, nil
}
