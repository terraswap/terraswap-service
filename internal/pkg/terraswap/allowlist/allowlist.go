package allowlist

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func GetAllowlistMapResponse[T AllowlistResponse](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var listMap T
	err = json.Unmarshal(body, &listMap)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse response(%v)", body)
	}
	return &listMap, nil
}
