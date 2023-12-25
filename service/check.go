package service

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func CheckStatus(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	status := doc.Find("#status-txt").Text()

	return status, nil
}

func CheckLastOnline(logs *os.File) (time.Time, error) {

	logs.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(logs)
	var onlineLines []string

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		if len(tokens) > 3 {
			if strings.Contains(tokens[3], "online") {
				onlineLines = append(onlineLines, line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	lastLine := onlineLines[len(onlineLines)-1]

	lastOnline, err := time.Parse("2006/01/02 15:04:05",
		fmt.Sprintf("%s %s", strings.Split(lastLine, " ")[0], strings.Split(lastLine, " ")[1]))
	if err != nil {
		return time.Time{}, err
	}
	return lastOnline, nil

}
