package zdesk

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"strconv"
)

func GetBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

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

