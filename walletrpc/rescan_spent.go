package walletrpc

// RescanSpent Rescan the blockchain for spent outputs.
func (c *Client) RescanSpent() error {
	err := c.do("rescan_spent", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
