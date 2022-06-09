package terraswap

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func ToTerraAmount(number string, decimals int) (string, error) {
	if decimals < 0 {
		return "", errors.New("decimals must be bigger than 0")
	}

	if err := checkFormat(number, decimals); err != nil {
		return "", errors.Wrapf(err, "wrong format amount(%s)", number)
	}

	zeros := strings.Repeat("0", int(decimals))
	baseStr := fmt.Sprintf("%s%s", "1", zeros)

	base, err := decimal.NewFromString(baseStr)
	if err != nil {
		return "", errors.Wrap(err, "cannot make base number")
	}

	d, err := decimal.NewFromString(number)
	if err != nil {
		err = errors.Wrapf(err, "cannot convert amount(%s)", number)
		return "", err
	}

	return d.Mul(base).String(), nil
}

func checkFormat(data string, decimals int) error {
	regexStr := fmt.Sprintf(`^[0-9]{1,30}(|\.[0-9]{1,%d})$`, decimals)
	re := regexp.MustCompile(regexStr)
	if !re.Match([]byte(data)) {
		return errors.New("data format is wrong, requires positive number")
	}
	return nil
}
