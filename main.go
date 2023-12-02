package main

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/jstern/aoc2023/aoc"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "list":
		fmt.Println("\nAvailable solutions (aka keys):")
		for _, k := range aoc.ListSolutions() {
			fmt.Printf("  * %s\n", k)
		}
	case "run":
		key := args[1]
		res := run(key)
		if res != nil {
			logAttempt(key, *res)
		}
	case "submit":
		key := args[1]
		res := run(key)
		if res != nil {
			fmt.Println("===\nSubmitting...\n===")
			submitRes := submit(key, *res)
			fmt.Println(submitRes)
			logAttempt(key, *res, submitRes)
		}
	case "all":
		for _, k := range aoc.ListSolutions() {
			fmt.Printf("\n%s\n", k)
			run(k)
		}
	case "stubs":
		stubs(args[1])
	default:
		panic("first arg must be 'run' or 'start'")
	}
}

const defaultTimeout = "300" // 5 minutes

type result struct {
	answer   string
	duration time.Duration
}

func run(key string) *result {
	attempt := aoc.SolutionFor(key)
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

	rc := make(chan result)
	go func() {
		start := time.Now()
		answer := attempt(input)
		rc <- result{answer, time.Since(start)}
	}()

	select {
	case res := <-rc:
		fmt.Printf("\nAnswer in %v\n---\n%s\n", res.duration, res.answer)
		return &res
	case <-time.After(time.Duration(wait) * time.Second):
		fmt.Println("Too slow!")
		return nil
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
	cached := cachedInput(year, day)
	if cached != "" {
		fmt.Println("Using cached input")
		return cached
	}
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
	err = cacheInput(year, day, bodyBytes)
	if err != nil {
		fmt.Printf("error saving input: %v\n", err)
	}
	return string(bodyBytes)
}

func cachedInput(year, day string) string {
	inpPath := filepath.Join(".aoc", fmt.Sprintf("input-%s-%s.txt", year, day))
	b, err := os.ReadFile(inpPath)
	if err != nil {
		return ""
	}
	return string(b)
}

func cacheInput(year, day string, inp []byte) error {
	inpPath := filepath.Join(".aoc", fmt.Sprintf("input-%s-%s.txt", year, day))
	err := os.WriteFile(inpPath, inp, 0644)
	if err != nil {
		return err
	}
	return nil
}

func submit(key string, result result) string {
	year, day := parseKey(key)
	level := strings.Split(key, ":")[2]
	token := strings.TrimSpace(os.Getenv("AOC_SESSION"))

	form := url.Values{}
	form.Add("answer", result.answer)
	form.Add("level", level)

	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", year, day)

	var client http.Client
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", "session="+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("submit post returned unexpected status: %d", resp.StatusCode))
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	return doc.Find("main article").First().Text()
}

//go:embed templates/*
var templateFS embed.FS

func stubs(key string) {
	year, day := parseKey(key)

	// bail if source/test files already exist

	srcName := fmt.Sprintf("y%sd%s.go", year, day)
	srcPath := filepath.Join("aoc", srcName)
	_, err := os.Stat(srcPath)
	if !errors.Is(err, os.ErrNotExist) {
		panic(fmt.Sprintf("%s exists", srcPath))
	}

	tstName := fmt.Sprintf("y%sd%s_test.go", year, day)
	tstPath := filepath.Join("aoc", tstName)
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
	err = tmp.ExecuteTemplate(srcOut, "solution.go.tmpl", data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("created %s\n", srcPath)

	tstOut, err := os.OpenFile(tstPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	err = tmp.ExecuteTemplate(tstOut, "solution_test.go.tmpl", data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("created %s\n", tstPath)
}

func logAttempt(key string, res result, messages ...string) {
	f, err := os.OpenFile(
		filepath.Join(".aoc", "log.txt"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Printf("[%s] Answer in %v: %s\n", key, res.duration, res.answer)
	for _, msg := range messages {
		logger.Printf("[%s] %s", key, msg)
	}
}
