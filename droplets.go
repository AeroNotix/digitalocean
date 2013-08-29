package digitalocean

import (
	"encoding/json"
	"errors"
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
	resp, err := http.Get(URL)
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
	type Response struct {
		Status   string
		Droplets []Droplet
	}
	resp, err := baserequest(fmt.Sprintf(Endpoint, ""), &Response{})
	if err != nil {
		return nil, err
	}
	if val, ok := resp.(*Response); ok {
		return val.Droplets, nil
	}
	return nil, err
}

func DropletByID(id int64) (*Droplet, error) {
	type Response struct {
		Status  string
		Droplet Droplet
	}
	resp, err := baserequest(
		fmt.Sprintf(Endpoint, fmt.Sprintf("%d", id)), &Response{},
	)
	if err != nil {
		return nil, err
	}
	if val, ok := resp.(*Response); ok {
		return &val.Droplet, nil
	}
	return nil, err
}

func RebootDroplet(id int64) error {
	type Response struct {
		Status   string
		Event_ID int
	}
	resp, err := baserequest(
		fmt.Sprintf(Endpoint, fmt.Sprintf("%d/reboot", id)), &Response{},
	)
	if err != nil {
		return err
	}
	if _, ok := resp.(*Response); ok {
		return nil
	}
	return errors.New("Invalid response from rebooting droplet.")
}
