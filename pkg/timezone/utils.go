package timezone

// ListFormats returns a list of common date/time formats in Go.
func ListFormats() map[string]string {
	return map[string]string{
		"ISO8601":     "2006-01-02T15:04:05Z07:00",
		"RFC3339":     "2006-01-02T15:04:05Z07:00",
		"RFC1123":     "Mon, 02 Jan 2006 15:04:05 MST",
		"RFC822":      "02 Jan 06 15:04 MST",
		"YYYY-MM-DD":  "2006-01-02",
		"HH:MM:SS":    "15:04:05",
		"FullDate":    "Monday, 02 January 2006",
		"ShortDate":   "02 Jan 2006",
		"Time12Hour":  "03:04:05 PM",
		"Time24Hour":  "15:04:05",
		"Custom":      "2006-01-02 15:04:05 MST",
	}
}


func ListTimezones() []string {
	return []string{
		"UTC", "America/New_York", "America/Los_Angeles", "America/Chicago",
		"Europe/London", "Europe/Berlin", "Asia/Tokyo", "Asia/Kolkata",
		"Australia/Sydney", "Africa/Johannesburg",
		// Add more if needed
	}
}