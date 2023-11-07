package main

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/jstern/aoc2023/solution"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "list":
		fmt.Println("\nAvailable solutions (aka keys):")
		for _, k := range solution.List() {
			fmt.Printf("  * %s\n", k)
		}
	case "run":
		run(args[1])
	case "all":
		for _, k := range solution.List() {
			fmt.Printf("\n%s\n", k)
			run(k)
		}

	case "stubs":
		start(args[1])
	default:
		panic("first arg must be 'run' or 'start'")
	}
}

const defaultTimeout = "300" // 5 minutes

type result struct {
	answer   string
	duration time.Duration
}

func run(key string) {
	attempt := solution.For(key)
	if attempt == nil {
		fmt.Println("no solution available for key")
		os.Exit(1)
	}

	year, day := parseKey(key)
	input := fetchInput(year, day)

	timeout := os.Getenv("AOC_TIMEOUT")
	if timeout == "" {
		timeout = defaultTimeout
	}
	wait, err := strconv.Atoi(timeout)
	if err != nil {
		panic(err)
	}

	res := make(chan result)
	go func() {
		start := time.Now()
		answer := attempt(input)
		result := result{answer, time.Since(start)}
		res <- result
	}()

	select {
	case result := <-res:
		fmt.Printf("---\nAnswer in %v\n---\n%s\n", result.duration, result.answer)
	case <-time.After(time.Duration(wait) * time.Second):
		fmt.Println("Too slow!")
	}
}

func parseKey(key string) (string, string) {
	parts := strings.Split(key, ":")
	if len(parts) < 2 {
		panic("expected at least <year>:<day> in key")
	}
	return parts[0], parts[1]
}

func fetchInput(year, day string) string {
	token := strings.TrimSpace(os.Getenv("AOC_SESSION"))

	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)

	var client http.Client
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", "session="+token)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("input fetch returned unexpected status: %d", resp.StatusCode))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(bodyBytes)
}

//go:embed templates/*
var templateFS embed.FS

func start(key string) {
	year, day := parseKey(key)

	// bail if source/test files already exist

	srcName := fmt.Sprintf("y%sd%s.go", year, day)
	srcPath := filepath.Join("solution", srcName)
	_, err := os.Stat(srcPath)
	if !errors.Is(err, os.ErrNotExist) {
		panic(fmt.Sprintf("%s exists", srcPath))
	}

	tstName := fmt.Sprintf("y%sd%s_test.go", year, day)
	tstPath := filepath.Join("solution", tstName)
	_, err = os.Stat(srcPath)
	if !errors.Is(err, os.ErrNotExist) {
		panic(fmt.Sprintf("%s exists", tstPath))
	}

	tmp, err := template.ParseFS(templateFS, "*/*.tmpl")
	if err != nil {
		panic(err)
	}

	data := map[string]string{
		"Year":     year,
		"Day":      day,
		"FuncName": fmt.Sprintf("y%sd%spart", year, day),
	}

	srcOut, err := os.OpenFile(srcPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	tmp.ExecuteTemplate(srcOut, "solution.go.tmpl", data)
	fmt.Printf("created %s\n", srcPath)

	tstOut, err := os.OpenFile(tstPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	tmp.ExecuteTemplate(tstOut, "solution_test.go.tmpl", data)
	fmt.Printf("created %s\n", tstPath)
}
