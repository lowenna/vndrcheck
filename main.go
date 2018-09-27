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

// For getting repo latest tag from github API
type githubRepo struct {
	Tag string `json:"tag_name,omitempty"`
}

// The versions of a repo being used and who is using it.
type version struct {
	usingRepos    []string
	vendorVersion string
}

// Information about a repo being used.
type externalRepoInfo struct {
	commits map[string]*version
}

var (

	// The Microsoft repos of particular interest
	microsoftRepos = []string{
		"Microsoft/opengcs",
		"Microsoft/hcsshim",
		"Microsoft/go-winio",
	}

	// This is a list of top-level repos which we use as starting points in our
	// search for all vendoring. containerd/cri and containerd/containerd are
	// good ones, as is moby/moby.
	externalRepos = []string{
		"moby/moby",
		"containerd/cri",
		"containerd/containerd",
		//"containerd/continuity", // easier for testing...
	}

	// This is a list of known bad repos - ones which shouldn't be present
	badRepos = []string{
		"boltdb/bolt",
	}

	allExternalRepos   map[string]*externalRepoInfo
	warnings           int
	count              int
	mismatchingImports int
)

func main() {

	allExternalRepos = make(map[string]*externalRepoInfo)

	// Find tags of latest release from each of the Microsoft repos
	fmt.Printf("Finding latest releases:\n\n")
	msRepos := make(map[string]string)
	for _, repo := range microsoftRepos {
		fmt.Printf("- %-23s : ", repo)
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

	fmt.Printf("Scanned %d repo(s) under github.com.\n", len(allExternalRepos))

	fmt.Printf("\nAnalysing the results:\n")
	for importedRepo, eri := range allExternalRepos {
		if len(eri.commits) > 1 {
			warnings++
			mismatchingImports++
			fmt.Printf("\n\nWARN: %s has %d versions imported\n", importedRepo, len(eri.commits))
			for _, importedBy := range eri.commits {
				fmt.Printf("\t%s by:\n", importedBy.vendorVersion)
				for _, usingRepo := range importedBy.usingRepos {
					fmt.Printf("\t\t%s\n", usingRepo)
				}
			}
		}
	}

	if warnings > 0 {
		fmt.Printf("\n%d warning(s) were found.\n", warnings)
	}

	if mismatchingImports > 0 {
		fmt.Printf("\n%d repo(s) are imported at different revisions.\n", mismatchingImports)
	}

}

// getJson gets json from a URL and decodes it
func getJson(url string, target interface{}) error {
	c := &http.Client{Timeout: 10 * time.Second}
	r, err := c.Get(url)
	if err != nil {
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
	//fmt.Printf(".")

	count++
	fmt.Printf("%4d: %-30.30s  %-16.16s  %s\n", count, repo, atCommit, parentRepo)

	if eriItemForRepo, ok := allExternalRepos[repo]; ok {
		// An entry is present in allExternalRepos. Does it match an existing commit?
		//fmt.Printf("eriItemForRepo: %+v\n", eriItemForRepo)

		for _, knownVersion := range eriItemForRepo.commits {
			//fmt.Printf("knownversion: %+v\n", knownVersion)
			if knownVersion.vendorVersion == atCommit {

				// Is this repo already in usingRepos?
				for _, usingRepo := range knownVersion.usingRepos {
					if usingRepo == parentRepo {
						// Nothing to do as already present. Stop recursing
						return
					}
				}
				// So we need to append that to the list of repos using this version
				knownVersion.usingRepos = append(knownVersion.usingRepos, parentRepo)
				return // Done - stop recursing further
			}
		}

		// So we have another version of this same repo in use. Add another version.
		eriItemForRepo.commits[atCommit] = &version{
			vendorVersion: atCommit,
			usingRepos:    []string{parentRepo},
		}

		// Stop recursing further
		return
	}

	// Add this repo at the commit.
	eri := &externalRepoInfo{
		commits: make(map[string]*version),
	}
	eri.commits[repo] = &version{
		usingRepos:    []string{parentRepo},
		vendorVersion: atCommit,
	}
	allExternalRepos[repo] = eri

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

		for _, bad := range badRepos {
			if bad == vendoredRepo {
				warnings++
				fmt.Printf("\n\nWARN: %q vendors known bad repo %q at %q\n\n", repo, bad, vendoredAt)
			}
		}

		// Go recusive  // TODO - could be multithreaded here...
		seedAllExternalReposFrom(vendoredRepo, vendoredAt, repo)
	}
}
