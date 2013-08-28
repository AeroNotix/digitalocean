package digitalocean

import (
	"net/http"
	"testing"
)

var droplets_get_response_good = `
{
    "status": "OK",
    "droplets": [
        {
            "id": 12345,
            "name": "the_name_of_the_server",
            "image_id": 350424,
            "size_id": 66,
            "region_id": 2,
            "backups_active": false,
            "ip_address": "37.139.18.183",
            "locked": false,
            "status": "active",
            "created_at": "2013-08-14T18:19:58Z"
        },
        {
            "id": 12346,
            "name": "the_other_name_of_the_server",
            "image_id": 350424,
            "size_id": 66,
            "region_id": 2,
            "backups_active": false,
            "ip_address": "37.139.17.114",
            "locked": false,
            "status": "active",
            "created_at": "2013-08-25T17:06:28Z"
        }
    ]
}`

func TestGetDroplets(t *testing.T) {
	httpTestsSetUp(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(droplets_get_response_good))
	})
	_, err := Droplets()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestFailGetDroplets(t *testing.T) {
	httpTestsSetUp(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`not_json`))
	})
	_, err := Droplets()
	if err == nil {
		t.Error(`Droplets() failed to handle receiving invalid data`)
	}
}
