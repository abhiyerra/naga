package main

import (
	"bytes"
	"flag"
	"fmt"
	. "github.com/abhiyerra/naga/gerrit"
	"log"
	"os/exec"
	"strings"
)

var (
	changeId   string
	revisionId string

	gerrit Gerrit
)

func cloneRepo(c *Change) (repoDir string) {
	fmt.Println("\nChecking out Repo:", c.Project)

	reposStr := strings.Split(c.Project, "/")
	repo := reposStr[len(reposStr)-1]
	repoDir = fmt.Sprintf("/tmp/%s", repo)

	cmd := exec.Command("git", "clone", "--depth", "1", gerrit.ProjectRepo(c.Project))
	cmd.Dir = "/tmp"
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Check out the change into the clone
	fmt.Printf("Checking out change to repo %s\n", repo)

	cmd2 := exec.Command("git", "fetch", gerrit.ProjectRepo(c.Project), c.Ref())
	cmd2.Dir = repoDir
	err = cmd2.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fetching head")
	cmd3 := exec.Command("git", "cherry-pick", "FETCH_HEAD")
	cmd3.Dir = repoDir
	err = cmd3.Run()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func bootProject(repoDir string) {
	fmt.Println("Booting up")
	cmd := exec.Command("vagrant", "up")
	cmd.Dir = repoDir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func setupProject(repoDir string) {
	fmt.Println("Booting up")
	cmd := exec.Command("vagrant", "ssh", "./naga-setup.sh")
	cmd.Dir = repoDir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func init() {
	gerrit = NewGerrit()

	flag.StringVar(&changeId, "change_id", "", "The change id to run")
	flag.StringVar(&revisionId, "revision", "", "The revision number to run")
	flag.Parse()
}

func main() {
	change := gerrit.Change(changeId, revisionId)
	repo := cloneRepo(change)
	bootProject(repo)
	setupProject(repo)
}
