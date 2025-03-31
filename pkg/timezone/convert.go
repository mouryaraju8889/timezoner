package timezone

import (
	"time"
	"fmt"
)


// Convert converts a given time from one timezone to another.
// Inputs:
//   - t: The time to be converted.
//   - fromZone: The source timezone as a string.
//   - toZone: The target timezone as a string.
// Outputs:
//   - fromTime: The original time in the source timezone.
//   - toTime: The converted time in the target timezone.
//   - error: An error if timezone loading fails.
func Convert(t time.Time, fromZone string, toZone string) (time.Time, time.Time, error) {
	// Load location
	fromLoc, err := time.LoadLocation(fromZone)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	toLoc, err := time.LoadLocation(toZone)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	// Convert time to UTC first, then to the target timezone
	utcTime := t.In(fromLoc).UTC()
	return t.In(fromLoc), utcTime.In(toLoc), nil
}

// GetUTCTime converts a given time from a specified timezone to UTC.
// Inputs:
//   - t: The time to be converted.
//   - fromZone: The source timezone as a string.
// Outputs:
//   - UTC time corresponding to the given input time.
//   - error: An error if timezone loading fails.
func GetUTCTime(t time.Time, fromZone string) (time.Time, error) {
	// Load location
	fromLoc, err := time.LoadLocation(fromZone)
	if err != nil {
		return time.Time{}, err
	}

	// Convert time to UTC first, then to the target timezone
	return t.In(fromLoc).UTC(), nil
}


// TimezoneDifference calculates the time difference between two timezones.
// Inputs:
//   - fromZone: The source timezone as a string.
//   - toZone: The target timezone as a string.
// Outputs:
//   - Time duration representing the difference between the two timezones.
//   - error: An error if timezone loading fails.
func TimezoneDifference(fromZone, toZone string) (time.Duration, error) {
	// Load timezones
	fromLoc, err := time.LoadLocation(fromZone)
	if err != nil {
		return 0, err
	}
	toLoc, err := time.LoadLocation(toZone)
	if err != nil {
		return 0, err
	}

	// Get the current time in UTC
	now := time.Now()

	// Get offsets of both timezones
	_, fromOffset := now.In(fromLoc).Zone()
	_, toOffset := now.In(toLoc).Zone()

	// Calculate the time difference in seconds
	offsetDiff := toOffset - fromOffset

	// Convert seconds to time.Duration
	return time.Duration(offsetDiff) * time.Second, nil
}

// ConvertNanosecondsToTimezone converts a given timestamp in nanoseconds to a specified timezone.
func ConvertNanosecondsToTimezone(nanoseconds int64, timezone string) (time.Time, error) {
	// Convert nanoseconds to a Time object
	t := time.Unix(0, nanoseconds) // 0 sec, nanoseconds provided

	// Load the target timezone
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	// Convert to the target timezone
	return t.In(loc), nil
}



func ConvertToTimezone(value int64, unit string, timezone string) (time.Time, error) {
	var t time.Time

	// Convert the input value to a time.Time object based on the unit
	switch unit {
	case "nanoseconds":
		t = time.Unix(0, value) // 0 seconds, value in nanoseconds
	case "microseconds":
		t = time.Unix(0, value*int64(time.Microsecond)) // Convert to nanoseconds
	case "milliseconds":
		t = time.Unix(0, value*int64(time.Millisecond)) // Convert to nanoseconds
	case "seconds":
		t = time.Unix(value, 0) // value in seconds
	default:
		return time.Time{}, fmt.Errorf("invalid unit: use 'nanoseconds', 'microseconds', 'milliseconds', or 'seconds'")
	}

	// Load the target timezone
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	// Convert to the target timezone
	return t.In(loc), nil
}

func ConvertWithFormat(t time.Time, fromZone, toZone, format string) (string, error) {
	// Load timezones
	fromLoc, err := time.LoadLocation(fromZone)
	if err != nil {
		return "", fmt.Errorf("invalid from timezone: %v", err)
	}
	toLoc, err := time.LoadLocation(toZone)
	if err != nil {
		return "", fmt.Errorf("invalid to timezone: %v", err)
	}

	// Convert time
	convertedTime := t.In(fromLoc).UTC().In(toLoc)

	// Check if the format exists in the predefined formats
	formats := ListFormats()
	if predefinedFormat, exists := formats[format]; exists {
		return convertedTime.Format(predefinedFormat), nil
	}
	format =  formats["Custom"]

	// If not found, assume the user provided a custom Go format
	return convertedTime.Format(format), nil
}

func EpochToTime(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}


func TimeToEpoch(t time.Time) int64 {
	return t.Unix()
}
