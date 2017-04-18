package zdesk

import (
	"fmt"
)

type TicketShowInput struct {
	ID string
}

func (c *Client) TicketShow(i *TicketShowInput) (interface{}, error) {
	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/api/v2/tickets/%s.json", i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var e interface{}

	if err := decodeJSON(&e, resp.Body); err != nil {
		return nil, err
	}

	return e, nil
}

