package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	TIMEOUT = 10 * time.Second
	URL     = "https://adventofcode.com/2020"
)

func main() {
	args := os.Args[1:]
	if err := run(args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("new expects a single arg 'day', got: %v", args)
	}
	// Arg should be a day integer
	day, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("argument to new should be a valid integer, got: %s", args[0])
	}

	_, here, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("could not get root directory")
	}
	root := filepath.Join(here, "../..")

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return err
	}

	session, ok := os.LookupEnv("AOC_SESSION")
	if !ok {
		return errors.New("missing AOC_SESSION in .env")
	}

	data, err := getInput(day, session)
	if err != nil {
		return err
	}

	err = makeDay(root, day, data)
	if err != nil {
		return err
	}

	return nil
}

// getInput gets a day's puzzle input and returns it as a byte slice
func getInput(day int, session string) ([]byte, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/day/%d/input", URL, day)

	client := &http.Client{
		Jar:     jar,
		Timeout: TIMEOUT,
	}

	sessionCookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(sessionCookie)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// makeDay creates the directory and files for a days advent of code and saves the input
func makeDay(root string, day int, data []byte) error {
	dayPath := filepath.Join(root, fmt.Sprintf("day%02d", day))

	err := os.Mkdir(dayPath, os.ModeSticky|os.ModePerm)
	if err != nil {
		return err
	}

	dayGo := filepath.Join(dayPath, fmt.Sprintf("day%02d.go", day))
	dayTest := filepath.Join(dayPath, fmt.Sprintf("day%02d_test.go", day))
	dayInput := filepath.Join(dayPath, fmt.Sprintf("day%02d.txt", day))

	err = os.WriteFile(dayGo, []byte("package main\n"), os.ModeSticky|os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(dayTest, []byte("package main\n"), os.ModeSticky|os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(dayInput, data, os.ModeSticky|os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
