package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func download(url string, name string) {
	// Get the data
	fmt.Printf("downloading %s\n", name)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return
	}

	// Create the file
	out, err := os.Create(name)
	if err != nil {
		return
	}
	defer out.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}
}

func main() {

	file, err := os.Open("chunklist_dvr.m3u8")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "8t6") {
			name := "media_" + strings.Split(text, "_")[3]
			download(fmt.Sprintf("https://bcovlive-a.akamaihd.net/a1a395c883004c05b3184dc2ea9570f1/eu-west-1/6252938537001/profile_0/%s", text), name)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
