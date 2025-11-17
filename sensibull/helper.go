package sensibull

import (
	"encoding/json"
	"fmt"
	"time"
)

func UnmarshalChartData(jsonBytes []byte) ([]SensibullData, error) {
	var root Root
	if err := json.Unmarshal(jsonBytes, &root); err != nil {
		return nil, err
	}
	var result []SensibullData
	for k, v := range root.Payload.ChartData {
		parsedTime, err := time.Parse(time.RFC3339, k)
		if err != nil {
			parsedTime, err = time.Parse("2006-01-02T15:04:05-07:00", k)
			if err != nil {
				return nil, fmt.Errorf("failed to parse timestamp key %q: %w", k, err)
			}
		}
		v.Timestamp = parsedTime
		result = append(result, v)
	}
	return result, nil
}