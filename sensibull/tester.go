package sensibull

import (
	"fmt"
	"os"
)

func main() {
    // Read JSON from local file
    jsonBytes, err := os.ReadFile("testdata.json")
    if err != nil {
        panic(err)
    }

    // Unmarshal and extract data
    dataSlice, err := UnmarshalChartData(jsonBytes)
    if err != nil {
        panic(err)
    }

    // Print results
    for _, d := range dataSlice {
        fmt.Printf("Timestamp: %v, Nifty: %v\n", d.Timestamp, d.Nifty)
    }
}