// pov is a command-line tool that read a "geotag" GeoJSON Feature on SDTIN and outputs its
// camera's "point of view area as a GeoJSON geometry element.
package main

import (
	"encoding/json"
	"github.com/sfomuseum/go-geojson-geotag/v2"
	"io"
	"log"
	"os"
)

func main() {

	body, err := io.ReadAll(os.Stdin)

	if err != nil {
		log.Fatalf("Failed to read data from STDIN, %v", err)
	}

	f, err := geotag.NewGeotagFeature(body)

	if err != nil {
		log.Fatalf("Failed to create geotag feature, %v", err)
	}

	fov, err := f.PointOfView()

	if err != nil {
		log.Fatalf("Failed to derive field of view from geotag feature, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(fov)

	if err != nil {
		log.Fatalf("Failed to encode field of view, %v", err)
	}
}
