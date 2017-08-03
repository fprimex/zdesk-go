package zdesk

import (
	"net/url"
	"strings"
	"strconv"
)

func GetIDFromURL(rawurl string) string {
	url, err := url.Parse(rawurl)
	if err != nil {
		return ""
	}

	path := url.EscapedPath()
	pathparts := strings.Split(path, "/")
	lastpart := pathparts[len(pathparts)-1]
	nameparts := strings.Split(lastpart, ".")
	name := nameparts[0]

	if _, err := strconv.Atoi(name); err == nil {
		return name
	} else {
		return ""
	}
}

