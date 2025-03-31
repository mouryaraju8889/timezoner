package timezone

import (
	"testing"
	"time"
	"github.com/mouryaraju8889/timezoner/pkg/timezone"
)

// Test Convert function
func TestConvert(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		fromZone string
		toZone   string
	}{
		{"IST to UTC", "Asia/Kolkata", "UTC"},
		{"UTC to EST", "UTC", "America/New_York"},
		{"PST to IST", "America/Los_Angeles", "Asia/Kolkata"},
	}

	// Get the current time
	now := time.Now()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fromTime, toTime, err := timezone.Convert(now, tc.fromZone, tc.toZone)
			if err != nil {
				t.Errorf("Convert(%s, %s) failed: %v", tc.fromZone, tc.toZone, err)
			}

			// Ensure conversion happened correctly
			if fromTime.IsZero() || toTime.IsZero() {
				t.Errorf("Conversion returned zero time for %s to %s", tc.fromZone, tc.toZone)
			}
		})
	}
}

// Test TimezoneDifference function
func TestTimezoneDifference(t *testing.T) {
	tests := []struct {
		fromZone string
		toZone   string
	}{
		{"UTC", "America/New_York"},
		{"Asia/Kolkata", "Europe/London"},
		{"America/Los_Angeles", "Asia/Tokyo"},
	}

	for _, tc := range tests {
		t.Run(tc.fromZone+" to "+tc.toZone, func(t *testing.T) {
			diff, err := timezone.TimezoneDifference(tc.fromZone, tc.toZone)
			if err != nil {
				t.Errorf("TimezoneDifference(%s, %s) failed: %v", tc.fromZone, tc.toZone, err)
			}

			// Ensure the difference is not zero (most timezones have differences)
			if diff == 0 {
				t.Errorf("Expected non-zero difference for %s to %s", tc.fromZone, tc.toZone)
			}
		})
	}
}

// Test ConvertToTimezone function
func TestConvertToTimezone(t *testing.T) {
	tests := []struct {
		value    int64
		unit     string
		timezone string
	}{
		{1672531199000000000, "nanoseconds", "America/New_York"},
		{1672531199000, "milliseconds", "Asia/Kolkata"},
		{1672531199, "seconds", "UTC"},
	}

	for _, tc := range tests {
		t.Run(tc.unit+" to "+tc.timezone, func(t *testing.T) {
			result, err := timezone.ConvertToTimezone(tc.value, tc.unit, tc.timezone)
			if err != nil {
				t.Errorf("ConvertToTimezone(%d, %s, %s) failed: %v", tc.value, tc.unit, tc.timezone, err)
			}

			// Ensure the returned time is valid
			if result.IsZero() {
				t.Errorf("Expected valid time conversion for %d %s to %s", tc.value, tc.unit, tc.timezone)
			}
		})
	}
}

// Test ParseWithNearestFormat function
func TestParseWithNearestFormat(t *testing.T) {
	tests := []struct {
		timeStr string
	}{
		{"2024-03-30T14:30:00Z"},
		{"Mon, 30 Mar 2024 14:30:00 MST"},
		{"30 Mar 2024"},
		{"14:30:00"},
	}

	for _, tc := range tests {
		t.Run(tc.timeStr, func(t *testing.T) {
		    format, err := timezone.ParseWithNearestFormat(tc.timeStr)
			if err != nil {
				t.Errorf("ParseWithNearestFormat(%s) failed: %v", tc.timeStr, err)
			}

			if format == "" {
				t.Errorf("Expected a valid format for %s, but got empty", tc.timeStr)
			}
		})
	}
}
