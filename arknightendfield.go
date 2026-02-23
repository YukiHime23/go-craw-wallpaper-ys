package crawal

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// ExtractHrefsFromScript extracts href values from script tags containing window.__resource
func ExtractHrefsFromScript(url string) ([]string, error) {
	client := &http.Client{Timeout: 30 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	htmlContent := string(body)
	return extractHrefsFromHTML(htmlContent), nil
}

// extractHrefsFromHTML finds script tags containing window.__resource and extracts href values
func extractHrefsFromHTML(htmlContent string) []string {
	scriptRegex := regexp.MustCompile(`<script[^>]*>(.*?)</script>`)
	scriptMatches := scriptRegex.FindAllStringSubmatch(htmlContent, -1)

	var hrefs []string
	for _, match := range scriptMatches {
		if len(match) > 1 {
			scriptContent := match[1]
			if strings.Contains(scriptContent, "window.__resource") {
				extractedHrefs := extractHrefsFromJS(scriptContent)
				hrefs = append(hrefs, extractedHrefs...)
			}
		}
	}

	return hrefs
}

// extractHrefsFromJS extracts href values from JavaScript content using regex
func extractHrefsFromJS(jsContent string) []string {
	hrefRegex := regexp.MustCompile(`href\s*:\s*"([^"]+)"`)
	matches := hrefRegex.FindAllStringSubmatch(jsContent, -1)

	var hrefs []string
	for _, match := range matches {
		if len(match) > 1 {
			hrefs = append(hrefs, match[1])
		}
	}

	return hrefs
}

// FilterImageURLs filters URLs to only include images with format {num}.{hash}.jpg
func FilterImageURLs(hrefs []string) []string {
	imageRegex := regexp.MustCompile(`/(\d{4})\.([a-f0-9]{6})\.jpg$`)
	var imageURLs []string

	for _, href := range hrefs {
		if imageRegex.MatchString(href) {
			imageURLs = append(imageURLs, href)
		}
	}

	return imageURLs
}

// DownloadImages downloads all filtered image URLs to a specified directory
func DownloadImages(imageURLs []string, downloadDir string) error {
	if len(imageURLs) == 0 {
		return fmt.Errorf("no image URLs to download")
	}

	fmt.Printf("Starting download of %d images to: %s\n", len(imageURLs), downloadDir)

	for i, url := range imageURLs {
		filename := filepath.Base(url)
		fmt.Printf("Downloading %d/%d: %s\n", i+1, len(imageURLs), filename)

		err := DownloadFile(url, "", downloadDir)
		if err != nil {
			fmt.Printf("Failed to download %s: %v\n", filename, err)
			continue
		}
	}

	fmt.Printf("Download completed!\n")
	return nil
}

// TestExtractHrefs demonstrates how to use the ExtractHrefsFromScript function
func TestExtractHrefs() {
	url := "https://endfield.gryphline.com/special/over-the-frontier"

	fmt.Printf("Fetching hrefs from: %s\n", url)
	hrefs, err := ExtractHrefsFromScript(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found %d href values:\n", len(hrefs))
	for i, href := range hrefs {
		fmt.Printf("%d: %s\n", i+1, href)
	}
}

// DownloadEndfieldImages extracts hrefs and downloads filtered images
func DownloadEndfieldImages() {
	url := "https://endfield.gryphline.com/special/over-the-frontier"

	fmt.Printf("Fetching hrefs from: %s\n", url)
	hrefs, err := ExtractHrefsFromScript(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found %d total href values\n", len(hrefs))

	imageURLs := FilterImageURLs(hrefs)
	fmt.Printf("Filtered to %d image URLs with format {num}.{hash}.jpg\n", len(imageURLs))

	downloadDir, err := CreateFolder("Downloads/endfield-images")
	if err != nil {
		fmt.Printf("Error creating download directory: %v\n", err)
		return
	}

	err = DownloadImages(imageURLs, downloadDir)
	if err != nil {
		fmt.Printf("Error downloading images: %v\n", err)
		return
	}
}
