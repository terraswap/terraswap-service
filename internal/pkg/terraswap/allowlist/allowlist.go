package allowlist

import (
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	jsPrefix = "module.exports = "
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

	yamlBody := strings.Replace(string(body), jsPrefix, "", 1)

	var listMap T
	err = yaml.Unmarshal([]byte(yamlBody), &listMap)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse response(%v)", yamlBody)
	}
	return &listMap, nil
}
