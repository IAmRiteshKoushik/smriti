package common

import (
	"fmt"
	"strings"
)

func ExtractSSEData(body []byte) ([]byte, error) {
	lines := strings.Split(string(body), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "data:") {
			return []byte(strings.TrimSpace(
				strings.TrimPrefix(line, "data:"),
			)), nil
		}
	}
	return nil, fmt.Errorf("no data field found")
}
