package gerrit

type Change struct {
	Project   string `json:"project"`
	Branch    string `json:"branch"`
	Revisions map[string]struct {
		_Number int
		Fetch   map[string]map[string]string
	}
}

func (c *Change) Ref() string {
	for _, v := range c.Revisions {
		return v.Fetch["http"]["ref"]
	}

	return ""
}
