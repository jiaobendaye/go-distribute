package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}

	res, err := http.Post(ServerURL, "application/json", buf)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("falied to register service. Registry sercie responded with code %v", res.StatusCode)
	}

	return nil
}

func ShtdownService(url string) error {
	req, err := http.NewRequest(http.MethodDelete, ServerURL, bytes.NewBuffer([]byte(url)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to deregister servcie, Registry service reponded with code %v",
			res.StatusCode)
	}

	return nil
}
