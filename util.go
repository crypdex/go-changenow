package changenow

import (
	"fmt"
	"strconv"
	"strings"

	resty "gopkg.in/resty.v1"
)

func checkResponse(response *resty.Response, err error) error {
	if err != nil {
		return err
	}

	if response.Error().(*RequestError).IsError() {
		return response.Error().(*RequestError)
	}
	return nil
}

func fromTo(source, target string) string {
	return fmt.Sprintf("%s_%s", strings.ToLower(source), strings.ToLower(target))
}

func sourceTarget(fromTo string) (string, string) {
	parts := strings.Split(fromTo, "_")
	return parts[0], parts[1]
}

func floatToString(num float64, options ...int) string {
	precision := 8
	if len(options) > 0 {
		precision = options[0]
	}
	// to convert a float number to a string
	return strconv.FormatFloat(num, 'f', precision, 64)
}
