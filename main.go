// Simple vendor checker for MS repos across container projects.
// It also looks for inconsistencies in vendoring across projects.
// Currently it is limited to only look through github.com hosted projects.
//
// By John Howard, September 2018. @jhowardmsft

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

// This is a list of top-level repos which we use as starting points in our
// search for all vendoring. containerd/cri and containerd/containerd are
// good ones.
var externalRepos = []string{
	//	"moby/moby",
	"containerd/cri",
	//"containerd/containerd",
	//"containerd/continuity", // easier for testing...
}

// This is a list of known bad repos - ones which shouldn't be present
var badRepos = []string{
	"boltdb/bolt",
}

type externalRepoInfo struct {
	commits map[string]string
}

var allExternalRepos map[string]*externalRepoInfo

func main() {

	allExternalRepos = make(map[string]*externalRepoInfo)

	// Find tags of latest release from each of the Microsoft repos
	fmt.Printf("Finding latest releases:\n\n")
	msRepos := make(map[string]string)
	for _, repo := range microsoftRepos {
		fmt.Printf("%20s : ", repo)
		ghr := githubRepo{}
		if err := getJson(fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo), &ghr); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", ghr.Tag)
		msRepos[repo] = ghr.Tag
	}

	// Get vendor.conf from each of the external repos
	fmt.Printf("\nAnalysing vendor.conf dependency chain:\n\n")
	for _, repo := range externalRepos {
		seedAllExternalReposFrom(repo, "", "")
		fmt.Println()
	}

	fmt.Println("We found:")
	for _, foo := range allExternalRepos {
		fmt.Println("...", foo)
	}
	//fmt.Printf("%+v", allExternalRepos)

}

// getJson gets json from a URL and decodes it
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

// getvndrconf gets vendor.conf from the root of a repo. If not found (or some
// other error occurs, it returns blank silently. Not the best, but it works.
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

// seedAllExternalReposFrom seeds our global `allExternalRepos` structure.
// It works recursively until all unique instances of mentioned repos have
// been found.
func seedAllExternalReposFrom(repo, atCommit, parentRepo string) {
	fmt.Printf(".")

	// Add this repo if it's not already present.
	if _, ok := allExternalRepos[repo]; !ok {
		eri := &externalRepoInfo{
			commits: make(map[string]string),
		}
		eri.commits[repo] = atCommit
		allExternalRepos[repo] = eri

		for _, bad := range badRepos {
			if bad == repo {
				fmt.Printf("\n\nWARN: %q vendors known bad repo %q at %q\n\n", parentRepo, bad, atCommit)
			}
		}
	}

	// Get the repo's vendor.conf
	vc, err := getvndrconf(repo)
	if err != nil {
		log.Fatal(err)
	}

	// Nothing to do if this repo doesn't vendor
	if vc == "" {
		return
	}

	// Get all lines of the vendor.conf file into an array
	vc = strings.Replace(vc, "\r", "", -1)
	vcLines := strings.Split(vc, "\n")

	// Loop through each line in the vendor.conf
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
		vendoredRepo := strings.TrimPrefix(lineSplit[0], "github.com/")
		vendoredAt := lineSplit[1]

		// Look to see if this line in the repos vendor.conf is already present
		// in our structure holding all the external repos we know about.
		if _, ok := allExternalRepos[vendoredRepo]; !ok {
			// As we just added a repo to the list of known, recurse into that
			// so we get the full tree built.
			seedAllExternalReposFrom(vendoredRepo, vendoredAt, repo)
		}
	}
}
