package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var tagPattern = regexp.MustCompile(`^v(\d+)\.(\d+)\.(\d+)$`)

type version struct {
	major int
	minor int
	patch int
}

func (v version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.major, v.minor, v.patch)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: nexttag <x|y|z>")
		fmt.Fprintln(os.Stderr, "  x = major bump, y = minor bump, z = patch bump")
		os.Exit(1)
	}

	bump := os.Args[1]
	if bump != "x" && bump != "y" && bump != "z" {
		fmt.Fprintf(os.Stderr, "invalid bump type %q: must be x, y, or z\n", bump)
		os.Exit(1)
	}

	current, err := highestTag()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching tags: %v\n", err)
		os.Exit(1)
	}

	next := current
	switch bump {
	case "x":
		next.major++
		next.minor = 0
		next.patch = 0
	case "y":
		next.minor++
		next.patch = 0
	case "z":
		next.patch++
	}

	fmt.Printf("git tag release/%s\n", next)
}

// highestTag returns the highest semver tag of the form vX.Y.Z.
// If no matching tag exists, it returns v0.0.0.
func highestTag() (version, error) {
	out, err := exec.Command("git", "tag").Output()
	if err != nil {
		return version{}, err
	}

	highest := version{}
	found := false

	for _, line := range strings.Split(string(out), "\n") {
		tag := strings.TrimSpace(line)
		m := tagPattern.FindStringSubmatch(tag)
		if m == nil {
			continue
		}

		major, _ := strconv.Atoi(m[1])
		minor, _ := strconv.Atoi(m[2])
		patch, _ := strconv.Atoi(m[3])

		v := version{major: major, minor: minor, patch: patch}
		if !found || compare(v, highest) > 0 {
			highest = v
			found = true
		}
	}

	return highest, nil
}

// compare returns -1, 0, or 1 depending on whether a is less than,
// equal to, or greater than b.
func compare(a, b version) int {
	if a.major != b.major {
		if a.major < b.major {
			return -1
		}
		return 1
	}
	if a.minor != b.minor {
		if a.minor < b.minor {
			return -1
		}
		return 1
	}
	if a.patch != b.patch {
		if a.patch < b.patch {
			return -1
		}
		return 1
	}
	return 0
}
