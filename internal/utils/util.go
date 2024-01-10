package utils

import (
	"log"
	"regexp"
	"strings"
)

var VERSION_FORMAT = regexp.MustCompile(`v[0-9](\.[0-9]+)+`)

func LogVersionWarnings(version string) {
	if strings.Contains(version, "test") {
		log.Println(
			"WARNING: The following version is a test version, do not run this in production",
		)
	} else if strings.Contains(version, "staging") {
		log.Println("WARNING: The following version is a staging version, do not run this in production")
	} else if !VERSION_FORMAT.Match([]byte(version)) {
		log.Fatalln("FATAL: Version must match format v[0-9](\\.[0-9]+)+")
	}
}
