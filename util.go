package zdesk

import (
	"net/url"
	"strings"
	"strconv"
)

func GetIDFromURL(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	path := url.EscapedPath()
	pathparts := strings.Split(path, "/")
	lastpart := pathparts[len(pathparts)-1]
	nameparts := strings.Split(lastpart, ".")
	name := nameparts[len(nameparts)-1]

	if _, err := strconv.Atoi(name); err == nil {
		return name, nil
	} else {
		return "", err
	}
}

