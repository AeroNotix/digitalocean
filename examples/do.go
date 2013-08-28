package main

import (
	"fmt"
	"github.com/AeroNotix/digitalocean"
)

func main() {
	digitalocean.Settings("ClientCode", "APIKey")
	fmt.Println(digitalocean.Droplets())
}
