// Package cpubenchmarknet is an unofficial library
// for downloading the CPU Mega List dataset by PassMark Software.
package cpubenchmarknet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const userAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"

// extractCookie visits a URL, extracts
// a cookie, and returns its value if it exists.
func extractCookie(url string, cookieName string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", userAgent)
	if err != nil {
		return "", err
	}
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == cookieName {
			return cookie.Value, nil
		}
	}
	return "", fmt.Errorf("%s cookie not found", cookieName)
}

// getJSONString sends a HTTP request and returns
// its body as string if it is valid JSON.
func getJSONString(req *http.Request) (string, error) {
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if !json.Valid(b) {
		return "", errors.New("HTTP Response Body is not valid JSON")
	}
	return string(b), nil
}

// GetCPUMegaList downloads CPU Mega List from cpubenchmark.net and returns its contents.
func GetCPUMegaList() (string, error) {
	cookieName := "PHPSESSID"
	cookieValue, err := extractCookie("https://cpubenchmark.net/CPU_mega_page.html", cookieName)
	if err != nil {
		return "", err
	}

	timeNowMilli := time.Now().UnixMilli()
	req, err := http.NewRequest("GET", fmt.Sprintf("https://cpubenchmark.net/data/?_=%d", timeNowMilli), nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", userAgent)
	req.AddCookie(&http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
	})

	return getJSONString(req)
}
