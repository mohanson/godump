package jany

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mohanson/godump/doa"
)

// See: https://datatracker.ietf.org/doc/html/rfc7159#section-13
var co string = strings.TrimSpace(`
{
    "Image": {
        "Width": 800,
        "Height": 600,
        "Title": "View from 15th Floor",
        "Thumbnail": {
            "Url": "http://www.example.com/image/481989943",
            "Height": 125,
            "Width": 100
        },
        "Animated": false,
        "IDs": [
            116,
            943,
            234,
            38793
        ]
    }
}
`)

var cl string = strings.TrimSpace(`
[
    {
        "precision": "zip",
        "Latitude": 37.7668,
        "Longitude": -122.3959,
        "Address": "",
        "City": "SAN FRANCISCO",
        "State": "CA",
        "Zip": "94107",
        "Country": "US"
    },
    {
        "precision": "zip",
        "Latitude": 37.371991,
        "Longitude": -122.026020,
        "Address": "",
        "City": "SUNNYVALE",
        "State": "CA",
        "Zip": "94085",
        "Country": "US"
    }
]
`)

func TestJany(t *testing.T) {
	o := doa.Try(Reader(bytes.NewReader([]byte(co))))
	doa.Doa(o.Get("Image").Get("Width").Uint32() == 800)
	doa.Doa(o.Get("Image").Get("Title").String() == "View from 15th Floor")
	doa.Doa(o.Get("Image").Get("IDs").Idx(0).Uint32() == 116)
	doa.Doa(o.Get("Image").Get("IDs").Idx(3).Uint32() == 38793)
	l := doa.Try(Reader(bytes.NewReader([]byte(cl))))
	doa.Doa(l.Idx(0).Get("Longitude").Float64() == -122.3959)
	doa.Doa(l.Idx(1).Get("City").String() == "SUNNYVALE")
}
