package gerrit

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type Gerrit struct {
	Host     string
	Username string
	Password string
}

func (g Gerrit) request(path string) []byte {
	out, err := exec.Command(
		"curl",
		"--silent",
		"-X", "GET",
		"--digest",
		"-u", g.Auth(),
		"--insecure",
		fmt.Sprintf("%s/%s", g.Host, path)).Output()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Ugh.... Gerrit returns a [)]} as the first line. Need
	// to kill that why this mess exists.
	return []byte(strings.Join(strings.Split(string(out), "\n")[1:], "\n"))
}

func (g Gerrit) Change(changeId string, revisionId string) (c *Change) {
	c = &Change{}

	changePath := fmt.Sprintf("a/changes/%s/revisions/%s/review", changeId, revisionId)
	changesBytes := g.request(changePath)

	err := json.Unmarshal(changesBytes, c)
	if err != nil {
		fmt.Println("error:", err)
	}

	return
}

func (g Gerrit) Auth() string {
	return fmt.Sprintf("%s:%s", g.Username, g.Password)
}

func (g Gerrit) ProjectRepo(project string) string {
	u, err := url.Parse("http://gerrit")
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("https://%s@%s/%s", u.Host, g.Auth(), project)
}

func NewGerrit() Gerrit {
	gerrit := Gerrit{}

	gerrit.Host = os.Getenv("BOOTUP_GERRIT_HOST")
	if gerrit.Host == "" {
		log.Println("Didn't set BOOTUP_GERRIT_HOST")
		gerrit.Host = "https://gerrit"
	}

	gerrit.Username = os.Getenv("BOOTUP_GERRIT_USERNAME")
	if gerrit.Username == "" {
		log.Fatal("Didn't set BOOTUP_GERRIT_USERNAME")
	}

	gerrit.Password = os.Getenv("BOOTUP_GERRIT_PASSWORD")
	if gerrit.Password == "" {
		log.Println("Didn't set BOOTUP_GERRIT_PASSWORD")
	}

	return gerrit
}
