package commands

import "fmt"

func Mapf(c *Config, _ []string) error {
	if c.Data.LocationsNext == "" && c.Data.LocationsPrev != "" {
		fmt.Println("you're on the last page")
		return nil
	}
	areasResponse, err := c.Client.GetLocations(c.Data.LocationsNext)
	if err != nil {
		return err
	}
	for _, area := range areasResponse.Results {
		fmt.Println(area.Name)
	}
	c.Data.LocationsNext = areasResponse.Next
	c.Data.LocationsPrev = areasResponse.Previous
	return nil
}

func Mapb(c *Config, _ []string) error {
	if c.Data.LocationsPrev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	areasResponse, err := c.Client.GetLocations(c.Data.LocationsPrev)
	if err != nil {
		return err
	}
	for _, area := range areasResponse.Results {
		fmt.Println(area.Name)
	}
	c.Data.LocationsNext = areasResponse.Next
	c.Data.LocationsPrev = areasResponse.Previous
	return nil
}
