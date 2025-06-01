package pokeapi

import (
	"fmt"
	"os"
)

func (c *Client) GetImage(url string) (*os.File, error) {
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch image: %s", resp.Status)
	}

	file, err := os.CreateTemp("", "sprite-*.png")
	if err != nil {
		return file, err
	}

	_, err = file.ReadFrom(resp.Body)
	file.Seek(0, 0)
	return file, err
}
