package advent

import (
	"fmt"
	"net/http"
	"path/filepath"
	"log"
	"io"
	"time"
	"flag"
	"os"

	"github.com/joho/godotenv"
)

func ParseFlags() (day, year int, cookie string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	today := time.Now()

	flag.IntVar(&day, "day", today.Day(), "day to fetch")
	flag.IntVar(&year, "year", today.Year(), "year to fetch")
	flag.StringVar(&cookie, "cookie", os.Getenv("SESSION_COOKIE"), "session cookie")
	flag.Parse()

	if day < 1 || day > 25  {
		log.Fatalf("day out of range %d", day)
	}

	if year < 2015 || year > today.Year() {
		log.Fatalf("year out of range %d", year)
	}

	if cookie == "" {
		log.Fatalf("need to set a cookie in .env SESSION_COOKIE or pass as cli argument -cookie")
	}
	
	return day, year, cookie
}

func GetInput(day, year int, cookie string) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	sessionCookie := http.Cookie{
		Name: "session",
		Value: cookie,
	}
	req.AddCookie(&sessionCookie)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error reading resp: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading resp body: %s", err)
	}

	fmt.Println("response length is", len(body))

	wd, err := os.Getwd()


	filename := filepath.Join(wd, fmt.Sprintf("src/Y%d/Day%d/input.txt", year, day))

	os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	os.WriteFile(filename, body, os.FileMode(0644))

	fmt.Println("Wrote file: ", filename)
}