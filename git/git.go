package git

import (
	"fmt"
	"runtime/debug"
)

var (
	commitTag    string //nolint:gochecknoglobals // Injected variable.
	commitAuthor string //nolint:gochecknoglobals // Injected variable.
)

type CommitInfo struct {
	Project  string
	Revision string
	Time     string
	Author   string
}

func (c CommitInfo) String() string {
	return fmt.Sprintf("%s at %s by %s", c.Revision, c.Time, c.Author)
}

func GetCommitInfo() CommitInfo {
	var project, revision, revisionTime string

	if info, ok := debug.ReadBuildInfo(); ok {
		project = info.Main.Path // <-- module/project name

		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				revision = setting.Value
			}

			if setting.Key == "vcs.time" {
				revisionTime = setting.Value
			}
		}
	}

	if commitTag != "" {
		revision = commitTag
	}

	return CommitInfo{
		Project:  project,
		Revision: revision,
		Time:     revisionTime,
		Author:   commitAuthor,
	}
}
