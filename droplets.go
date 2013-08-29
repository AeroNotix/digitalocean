package digitalocean

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	Endpoint string = "https://api.digitalocean.com/droplets/%s?client_id=%s&api_key=%s"
)

func Settings(ClientID, APIKey string) {
	Endpoint = fmt.Sprintf(Endpoint, "%s", ClientID, APIKey)
}

type Droplet struct {
	ID             int64
	Name           string
	Image_ID       int
	Size_ID        int
	Region_ID      int
	Backups_Active bool
	IP_Address     string
	Locked         bool
	Status         string
	Created_At     time.Time
}

func baserequest(URL string, Type interface{}) (interface{}, error) {
	resp, err := http.Get(fmt.Sprintf(Endpoint, ""))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, Type)
	return Type, nil
}

func Droplets() ([]Droplet, error) {
	type Request struct {
		Status   string
		Droplets []Droplet
	}
	resp, err := baserequest(fmt.Sprintf(Endpoint, ""), &Request{})
	if err != nil {
		return nil, err
	}
	if val, ok := resp.(*Request); ok {
		return val.Droplets, nil
	}
	return nil, err
}
