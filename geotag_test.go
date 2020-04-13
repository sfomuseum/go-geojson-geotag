package geotag

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestParseGeotagFeature(t *testing.T) {

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	fixtures := filepath.Join(cwd, "fixtures")
	feature_path := filepath.Join(fixtures, "geotag.geojson")

	feature_fh, err := os.Open(feature_path)

	if err != nil {
		t.Fatal(err)
	}

	defer feature_fh.Close()

	feature, err := NewGeotagFeatureWithReader(feature_fh)

	if err != nil {
		t.Fatal(err)
	}

	pov, err := feature.PointOfView()

	if err != nil {
		t.Fatal(err)
	}

	if pov.Coordinates[0] != -122.37499034916583 {
		t.Fatal("Invalid point of view longitude")
	}

	if pov.Coordinates[1] != 37.62868949010699 {
		t.Fatal("Invalid point of view longitude")
	}

	hl, err := feature.HorizonLine()

	if err != nil {
		t.Fatal(err)
	}

	if hl.Coordinates[0][0] != -122.3942243343624 {
		t.Fatal("Invalid horizon line longitude")
	}

	if hl.Coordinates[0][1] != 37.58120271661385 {
		t.Fatal("Invalid horizon line longitude")
	}

	if hl.Coordinates[1][0] != -122.43198473259156 {
		t.Fatal("Invalid horizon line longitude")
	}

	if hl.Coordinates[1][1] != 37.60749730586552 {
		t.Fatal("Invalid horizon line longitude")
	}

	fov, err := feature.FieldOfView()

	if err != nil {
		t.Fatal(err)
	}

	if len(fov.Coordinates[0]) != 4 {
		t.Fatal("Invalid field of view length")
	}

	/*
		trg, err := feature.Target()

		if err != nil {
			t.Fatal(err)
		}
	*/

	fov_feature := FieldOfViewFeature{
		Type:       "Feature",
		Properties: feature.Properties,
		Geometry:   fov,
	}

	enc, err := json.Marshal(fov_feature)

	if err != nil {
		t.Fatal(err)
	}

	hash := sha256.Sum256(enc)
	hash_str := fmt.Sprintf("%x", hash[:])

	if hash_str != "67f9a804bb0c2946905219bc9cfca0f3fa7ffeb78d10f168ea53783db4f393f8" {
		t.Fatal("Invalid field of view feature hash")
	}
}
