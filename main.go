// Really dumb simple vendor checker for MS repos across container projects
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type githubRepo struct {
	Tag string `json:"tag_name,omitempty"`
}

type importRepo struct {
	repoName string
	version  map[string]string
}

var microsoftRepos = []string{
	"Microsoft/opengcs",
	"Microsoft/hcsshim",
	"Microsoft/go-winio",
}

var externalRepos = []string{

	//	"moby/moby",
	"containerd/cri",
	"containerd/containerd",
	"containerd/continuity", // easier for testing...
}

type externalRepoInfo struct {
	commits map[string]string
}

var allExternalRepos map[string]*externalRepoInfo

func main() {

	allExternalRepos = make(map[string]*externalRepoInfo)

	// Find tags of latest release from each of the Microsoft repos
	fmt.Println("Finding latest releases...")
	msRepos := make(map[string]string)
	for _, repo := range microsoftRepos {
		fmt.Printf("-%10s : ", repo)
		ghr := githubRepo{}
		if err := getJson(fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo), &ghr); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", ghr.Tag)
		msRepos[repo] = ghr.Tag
	}

	// Get vendor.conf from each of the external repos
	fmt.Println("Analysing vendor.conf from external repos recursively...")
	for _, repo := range externalRepos {
		seedAllExternalReposFrom(repo)
		fmt.Println()
	}

	fmt.Println("We found:")
	for _, foo := range allExternalRepos {
		fmt.Println("...", foo)
	}
	//fmt.Printf("%+v", allExternalRepos)

}

func getJson(url string, target interface{}) error {
	c := &http.Client{Timeout: 10 * time.Second}
	r, err := c.Get(url)
	if err != nil {
		fmt.Println("Oops:", url)
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func getvndrconf(repo string) (string, error) {
	c := &http.Client{Timeout: 10 * time.Second}
	r, err := c.Get(fmt.Sprintf("https://raw.githubusercontent.com/%s/master/vendor.conf", repo))
	if err != nil {
		fmt.Println("Failed to find", fmt.Sprintf("https://raw.githubusercontent.com/%s/master/vendor.conf", repo))
		return "", err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return "", nil
	}
	vc, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(vc), nil
}
func seedAllExternalReposFrom(repo string) {

	fmt.Printf("- %s : \n", repo)

	if _, ok := allExternalRepos[repo]; !ok {
		eri := &externalRepoInfo{
			commits: make(map[string]string),
		}
		eri.commits[repo] = ""
		allExternalRepos[repo] = eri
	}

	vc, err := getvndrconf(repo)
	if err != nil {
		log.Fatal(err)
	}
	if vc == "" {
		return // Might be empty
	}

	vc = strings.Replace(vc, "\r", "", -1)
	vcLines := strings.Split(vc, "\n")

	for _, line := range vcLines {
		// Ignore blanks, comments
		if len(line) == 0 || string(line[0]) == "#" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		if len(lineSplit) != 2 { // Can't cope with things with an alias
			//log.Printf("WARN: Ignoring %s", lineSplit)
			continue
		}
		if !strings.HasPrefix(lineSplit[0], "github.com") { // For now anyway
			//log.Printf("WARN: Ignoring %s", lineSplit)
			continue
		}
		lineSplit[0] = strings.TrimPrefix(lineSplit[0], "github.com/")
		if _, ok := allExternalRepos[lineSplit[0]]; !ok {
			//fmt.Println(lineSplit[0], "is not in allExternalRepos")
			eri := &externalRepoInfo{
				commits: make(map[string]string),
			}
			eri.commits[lineSplit[0]] = lineSplit[1]
			allExternalRepos[lineSplit[0]] = eri

			seedAllExternalReposFrom(lineSplit[0]) // Recurse...
		}
	}
}
