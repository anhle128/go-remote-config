package remoteconfig

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//LoadConfigFromURL Downloads a configuration JSON
func LoadConfigFromURL(configURL string, configStruct interface{}) error {
	resp, err := http.Get(configURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request to '%s' returned non-200 OK status '%d: %s'", configURL, resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Do a streaming JSON decode
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(configStruct); err != nil {
		return fmt.Errorf("Failed to decode JSON, with error, %s", err.Error())
	}

	return nil
}
