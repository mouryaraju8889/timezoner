// package main

// import (
// 	"fmt"
// 	"time"
// 	// "log"
// 	"github.com/mouryaraju8889/timezoner/pkg/timezone"
// )

// func main() {
// 	now := time.Now()
// 	from, converted, err := timezone.Convert(now, "Asia/Kolkata", "America/New_York")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("from Time:", from)

// 	fmt.Println("Converted Time:", converted)

// 	utc1, err := timezone.GetUTCTime(now, "Asia/Kolkata")

// 	fmt.Println("UTC Time:", utc1)

// 	diff, err := timezone.TimezoneDifference("America/New_York", "Asia/Kolkata")

// 	fmt.Println("Timezone Difference:", diff)

// 	// _ := int64(1700000000000000000)
// 	toLoc, _ := time.LoadLocation("Asia/Kolkata")

// 	nanoToTimezone, err := timezone.ConvertNanosecondsToTimezone(now.In(toLoc).UnixNano(), "Asia/Kolkata")

// 	fmt.Println("Nanoseconds to Timezone:", nanoToTimezone)

// 	nanoToTimezone2, err := timezone.ConvertToTimezone(now.In(toLoc).UnixMicro(), "microseconds", "Asia/Kolkata")

// 	fmt.Println("microseconds to Timezone:", nanoToTimezone2)

// 	fmt.Println("-------------------------------------------------------------")

// 	formattedTime, err := timezone.ConvertWithFormat(now, "UTC", "Asia/Kolkata", "RFC3339")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Formatted Time (RFC3339):", formattedTime)

// 	// Convert and format time using a custom format
// 	formattedCustom, err := timezone.ConvertWithFormat(now, "UTC", "Asia/Kolkata", "2006 mON")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Formatted Time (Custom):", formattedCustom)

// 	inputTime := "2025-03-30 18:30:00 IST"

// 	// Try parsing with nearest format
// 	detectedFormat, err := timezone.ParseWithNearestFormat(inputTime)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Detected Format:", detectedFormat)
// 		// fmt.Println("Parsed Time:", parsedTime)
// 	}

// 	timestamps := []uint64{
// 		0,                      // Epoch start (1970-01-01 00:00:00 UTC) → Seconds
// 		946684800,              // Year 2000 in seconds
// 		1711785600,             // Recent timestamp in seconds
// 		1711785600000,          // Milliseconds
// 		1711785600000000,       // Microseconds
// 		1711785600000000000,    // Nanoseconds
// 	}

// 	// Detect time units
// 	for _, ts := range timestamps {
// 		unit, err := timezone.ParseTimeUnit(ts)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 		} else {
// 			fmt.Printf("Timestamp: %d → Unit: %s\n", ts, unit)
// 		}
// 	}


// 	input := "yesterday at 5pm"

// 	parsedTime, err := timezone.ParseNaturalDate(input)
// 	if err != nil {
// 		fmt.Println("Error parsing natural date:", err)
// 		return
// 	}

// 	fmt.Println("Parsed natural language date:", parsedTime.Format(time.RFC1123))

// }
