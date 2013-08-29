package main

import (
	"fmt"
	"github.com/AeroNotix/digitalocean"
	"strings"
)

func main() {
	digitalocean.Settings("ClientID", "APIKey")
	droplets, err := digitalocean.Droplets()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Repeat("-", 50))
	for _, droplet := range droplets {
		fmt.Println(digitalocean.DropletByID(droplet.ID))
	}
}
