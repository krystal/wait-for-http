package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	flag "github.com/spf13/pflag"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

func main() {
	checkQuantity := flag.Int("quantity", 30, "number of times to check")
	sleepTime := flag.Int("sleep", 5, "sleep time between checks")
	timeoutSeconds := flag.Int("timeout", 5, "HTTP timeout")
	statusCodesStr := flag.String("statuses", "200", "acceptable HTTP status codes")
	insecure := flag.Bool("insecure", false, "should SSL issues be ignored?")

	flag.Parse()

	url := flag.Arg(0)
	if url == "" {
		logger.Fatal("must specify hostname")
	}

	statusCodes := parseStatusCodes(*statusCodesStr)

	checksCompleted := 0
	for {
		statusCode := runCheck(url, *timeoutSeconds, *insecure)

		if arrayOfIntsIncludes(statusCodes, statusCode) {
			logger.Printf("service was available at %s with status %d\n", url, statusCode)
			os.Exit(0)
		}

		if checksCompleted >= *checkQuantity {
			logger.Fatalf("did not become available after %d checks", *checkQuantity)
		}

		checksCompleted += 1
		logger.Printf("%s not available (status: %d), sleeping %d seconds", url, statusCode, *sleepTime)
		time.Sleep(time.Duration(*sleepTime) * time.Second)
	}
}

func arrayOfIntsIncludes(array []int, value int) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}

	return false
}

func runCheck(url string, timeoutSeconds int, insecure bool) int {
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeoutSeconds),
	}

	if insecure {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = transport
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Print(err)
		return 0
	}

	req.Header.Set("User-Agent", "wait-for-http")

	resp, err := client.Do(req)
	if err != nil {
		logger.Print(err)
		return 0
	}

	return resp.StatusCode
}

func parseStatusCodes(str string) []int {
	ints := []int{}
	for _, code := range strings.Split(str, ",") {
		int, err := strconv.Atoi(code)
		if err == nil {
			ints = append(ints, int)
		}
	}
	return ints
}
