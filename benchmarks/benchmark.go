package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Prompt struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Prompt   string `json:"prompt"`
	Expected string `json:"expected"`
}

type Result struct {
	ID         string  `json:"id"`
	Category   string  `json:"category"`
	Prompt     string  `json:"prompt"`
	Expected   string  `json:"expected"`
	Response   string  `json:"response"`
	RawOutput  string  `json:"raw_output"`
	DurationMS int64   `json:"duration_ms"`
	Success    bool    `json:"success"`
	Error      *string `json:"error,omitempty"`
}

func extractResponse(raw, prompt string) string {
	marker := "> " + prompt

	start := strings.Index(raw, marker)
	if start == -1 {
		return strings.TrimSpace(raw)
	}

	// Remove tudo que veio antes do prompt.
	response := raw[start+len(marker):]

	// Remove as métricas do llama.cpp e tudo que vem depois.
	if end := strings.Index(response, "[ Prompt:"); end != -1 {
		response = response[:end]
	}

	return strings.TrimSpace(response)
}

func main() {
	model := flag.String(
		"model",
		"",
		"Path to the GGUF model",
	)

	dataset := flag.String(
		"dataset",
		"datasets/educational-v0.1/prompts.json",
		"Path to evaluation dataset",
	)

	output := flag.String(
		"output",
		"experiments/001-smollm2-135m-baseline/educational-results-f16.json",
		"Path to output JSON",
	)

	flag.Parse()

	if *model == "" {
		fmt.Fprintln(os.Stderr, "error: --model is required")
		os.Exit(1)
	}

	// Carrega o dataset.
	data, err := os.ReadFile(*dataset)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"error reading dataset: %v\n",
			err,
		)
		os.Exit(1)
	}

	var prompts []Prompt

	if err := json.Unmarshal(data, &prompts); err != nil {
		fmt.Fprintf(
			os.Stderr,
			"error parsing dataset: %v\n",
			err,
		)
		os.Exit(1)
	}

	if len(prompts) == 0 {
		fmt.Fprintln(os.Stderr, "error: dataset contains no prompts")
		os.Exit(1)
	}

	fmt.Printf("Model: %s\n", *model)
	fmt.Printf("Dataset: %s\n", *dataset)
	fmt.Printf("Prompts: %d\n\n", len(prompts))

	results := make([]Result, 0, len(prompts))

	for i, p := range prompts {
		fmt.Printf(
			"Prompt %d/%d: %s (%s)\n",
			i+1,
			len(prompts),
			p.ID,
			p.Category,
		)

		start := time.Now()

		cmd := exec.Command(
			"llama-cli",
			"-m", *model,
			"-p", p.Prompt,
			"-n", "128",
			"--temp", "0",
			"-ngl", "0",
			"--single-turn",
			"--simple-io",
			"--log-disable",
			"--no-display-prompt",
		)

		outputBytes, err := cmd.CombinedOutput()

		duration := time.Since(start)

		rawOutput := string(outputBytes)

		result := Result{
			ID:         p.ID,
			Category:   p.Category,
			Prompt:     p.Prompt,
			Expected:   p.Expected,
			Response:   extractResponse(rawOutput, p.Prompt),
			RawOutput:  strings.TrimSpace(rawOutput),
			DurationMS: duration.Milliseconds(),
			Success:    err == nil,
		}

		if err != nil {
			errMessage := err.Error()
			result.Error = &errMessage

			fmt.Printf(
				"  ERROR: %s\n",
				errMessage,
			)
		} else {
			fmt.Printf(
				"  completed in %.2fs\n",
				duration.Seconds(),
			)
		}

		results = append(results, result)
	}

	resultData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"error encoding results: %v\n",
			err,
		)
		os.Exit(1)
	}

	if err := os.WriteFile(*output, resultData, 0644); err != nil {
		fmt.Fprintf(
			os.Stderr,
			"error writing results: %v\n",
			err,
		)
		os.Exit(1)
	}

	fmt.Printf("\nFinished.\n")
	fmt.Printf("Results: %s\n", *output)
}
