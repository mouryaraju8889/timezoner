package timezone


import (
	"time"
	// "fmt"
	"errors"
)


func ParseWithNearestFormat(timeStr string) (string, error) {
	//use ListFormats() to get the formats
	formats := ListFormats()
	for _, format := range formats {
		_, err := time.Parse(format, timeStr)
		// fmt.Println("Parsed Time:", parsedTime, "Format:", format, "Error:", err)
		if err == nil {
			return format, nil
		}
	}
	return "", errors.New("no matching format found")
}



func ParseTimeUnit(timestamp uint64) (string, error) {
	// Define reasonable epoch ranges
	lowerBound := uint64(0)           // Epoch start (1970)
	now := uint64(time.Now().Unix())  // Current Unix time in seconds
	upperBound := now + 10            // Allow slight future timestamps

	// Convert timestamp to different units
	seconds := timestamp
	milliseconds := timestamp / 1_000
	microseconds := timestamp / 1_000_000
	nanoseconds := timestamp / 1_000_000_000

	// Determine the most likely unit
	switch {
	case lowerBound <= seconds && seconds <= upperBound:
		return "seconds", nil
	case lowerBound <= milliseconds && milliseconds <= upperBound:
		return "milliseconds", nil
	case lowerBound <= microseconds && microseconds <= upperBound:
		return "microseconds", nil
	case lowerBound <= nanoseconds && nanoseconds <= upperBound:
		return "nanoseconds", nil
	default:
		return "", errors.New("invalid timestamp: out of expected range")
	}
}