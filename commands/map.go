package commands

import "fmt"

func Mapf(c *Config) error {
	if c.LocationsNext == "" && c.LocationsPrev != "" {
		fmt.Println("you're on the last page")
		return nil
	}
	areasResponse, err := c.Client.GetLocations(c.LocationsNext)
	if err != nil {
		return err
	}
	for _, area := range areasResponse.Results {
		fmt.Println(area.Name)
	}
	c.LocationsNext = areasResponse.Next
	c.LocationsPrev = areasResponse.Previous
	return nil
}

func Mapb(c *Config) error {
	if c.LocationsPrev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	areasResponse, err := c.Client.GetLocations(c.LocationsPrev)
	if err != nil {
		return err
	}
	for _, area := range areasResponse.Results {
		fmt.Println(area.Name)
	}
	c.LocationsNext = areasResponse.Next
	c.LocationsPrev = areasResponse.Previous
	return nil
}
