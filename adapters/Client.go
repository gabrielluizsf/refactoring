package adapters

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(comp Computer) {
    fmt.Println("Client inserts Lightning connector into computer.")
    comp.InsertIntoLightningPort()
}
