package version

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestCommittedVersionMatchesMakefile(t *testing.T) {
	makefile, err := os.ReadFile("../../Makefile")
	if err != nil {
		t.Fatal(err)
	}
	match := regexp.MustCompile(`(?m)^DEFAULT_VERSION := (\S+)$`).FindSubmatch(makefile)
	if len(match) != 2 {
		t.Fatal("Makefile must define DEFAULT_VERSION")
	}
	if Version != string(match[1]) {
		t.Fatalf("binary version %q differs from Makefile version %q", Version, match[1])
	}
}

func TestBuildInfoIncludesReleaseVersion(t *testing.T) {
	if info := Get(); info.Version != Version {
		t.Fatalf("build info version %q differs from release version %q", info.Version, Version)
	}
	if info := GetFormattedBuildInfo(); !strings.Contains(info, "Version="+Version+",") {
		t.Fatalf("build info does not report the release version: %s", info)
	}
}
