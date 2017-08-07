package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/fprimex/zdesk-go"
	"github.com/tidwall/gjson"
)

func main() {
	// must add "os" to imports for this to work
	//client, err := zdesk.NewClient(fmt.Sprintf("%s/token",
	//	os.Getenv("ZDESK_USERNAME")),
	//	os.Getenv("ZDESK_TOKEN"),
	//	os.Getenv("ZDESK_DOMAIN"))

	// The DefaultClient uses all of the above env vars for you.
	client, err := zdesk.DefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	/*
	 * Tickets
	 */
	newticket := []byte(`{"ticket":{"subject":"zdesk-go test ticket","description":"test ticket description"}}`)

	// create
	resp, err := client.TicketCreate(&zdesk.RequestOptions{
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       bytes.NewReader(newticket),
		BodyLength: int64(len(newticket))})
	if err != nil {
		log.Fatal(fmt.Sprintf("TicketCreate: %s", err))
	}

	fmt.Printf("Created ticket location:\n%s\n\n", resp.Header.Get("Location"))

	// Get id from url using the convenience function from zdesk util.go
	ticketid := zdesk.GetIDFromURL(resp.Header.Get("Location"))
	fmt.Printf("Created ticket ID: %s\n\n", ticketid)

	// retrieve
	resp, err = client.TicketShow(&zdesk.TicketShowInput{ID: ticketid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("TicketShow: %s", err))
	}

	body, err := zdesk.GetBody(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Retrieved ticket %s subject:\n%s\n\n", ticketid, gjson.GetBytes(body, "ticket.subject"))

	// update
	ticketupdate := []byte(`{"ticket": {"status": "open", "comment": { "body": "Update to the test ticket"}}}`)
	resp, err = client.TicketUpdate(&zdesk.TicketUpdateInput{ID: ticketid},
		&zdesk.RequestOptions{
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       bytes.NewReader(ticketupdate),
			BodyLength: int64(len(ticketupdate))})
	if err != nil {
		log.Fatal(fmt.Sprintf("TicketUpdate: %s", err))
	}
	fmt.Printf("Update ticket %s return code: %d\n\n", ticketid, resp.StatusCode)

	// list
	resp, err = client.TicketsList(nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("TicketsList: %s", err))
	}

	body, err = zdesk.GetBody(resp)
	if err != nil {
		log.Fatal(err)
	}

	println("Listing tickets:")
	ticketlist := gjson.GetBytes(body, "tickets")
	ticketlist.ForEach(func(key, value gjson.Result) bool {
		fmt.Printf("%s: %s\n",
			gjson.Get(value.String(), "id"),
			gjson.Get(value.String(), "subject"))
		return true // keep iterating
	})
	println("")

	// delete
	resp, err = client.TicketDelete(&zdesk.TicketDeleteInput{ID: ticketid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("TicketDelete: %s", err))
	}

	fmt.Printf("Delete ticket %s return code: %d\n\n", ticketid, resp.StatusCode)
}
