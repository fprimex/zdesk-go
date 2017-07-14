package zdesk

import (
	"fmt"
	"errors"
    "io/ioutil"
)

type TicketShowInput struct {
	ID string
}

func (c *Client) TicketShow(i *TicketShowInput) ([]byte, error) {
	if i.ID == "" {
		return nil, errors.New("Missing required field 'ID'")
	}

	path := fmt.Sprintf("/api/v2/tickets/%s.json", i.ID)
	resp, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}

