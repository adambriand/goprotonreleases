package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ReleaseInfo struct {
	Name        string `json:"name"`
	TagName     string `json:"tag_name"`
	Url         string `json:"html_url"`
	PublishDate string `json:"published_at"`
}

func getReleaseInfo(url string) (ReleaseInfo, error) {
	var release ReleaseInfo

	resp, err := http.Get(url)
	if err != nil {
		return release, errors.New("Error - cannot retrieve release url")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return release, errors.New("Error - cannot read response body")
	}

	errJson := json.Unmarshal([]byte(body), &release)

	if errJson != nil {
		return release, errors.New("Error - cannot unmarshal json")
	}

	return release, nil
}

func main() {
	colorBlue := "\033[34m"
	colorCyan := "\033[36m"
	colorReset := "\033[0m"

	Proton := "https://api.github.com/repos/ValveSoftware/Proton/releases/latest"
	ProtonGe := "https://api.github.com/repos/GloriousEggroll/proton-ge-custom/releases/latest"

	urlArray := []string{Proton, ProtonGe}

	for _, url := range urlArray {
		release, err := getReleaseInfo(url)
		if err != nil {
			fmt.Println(err)
			break;
		}

		fmt.Println(string(colorBlue), release.Name, "(", release.TagName, ")")
		publishDateTime, _ := time.Parse(time.RFC3339, release.PublishDate)
		fmt.Println(string(colorCyan), "Released on", publishDateTime.Format("Jan-02-2006"))
		fmt.Println(string(colorCyan), release.Url)
		fmt.Println(string(colorReset))
	}
}
