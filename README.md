## Timezoner - A Golang Package for Timezone Conversions

Timezoner is a Golang package that provides various utilities for timezone conversions, epoch time handling, and format parsing.

### 📦 Installation
```sh
go get github.com/mouryaraju8889/timezoner
```

### 🚀 Features
- Convert time between different timezones
- Get UTC time from any timezone
- Calculate the difference between two timezones
- Convert epoch timestamps to time
- Parse user input with nearest matching format
- List available timezones and formats

### 📚 Usage
```go
package main

import (
	"fmt"
	"time"
	"github.com/mouryaraju8889/timezoner/pkg/timezone"
)

func main() {
	// Get current time
	t := time.Now()
	fromZone := "Asia/Kolkata"
	toZone := "America/New_York"

	// Convert time between timezones
	fromTime, toTime, err := timezone.Convert(t, fromZone, toZone)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Original Time in", fromZone, ":", fromTime)
	fmt.Println("Converted Time in", toZone, ":", toTime)

	// Get UTC time
	utcTime, err := timezone.GetUTCTime(t, fromZone)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("UTC Time:", utcTime)
}
```

### 🛠 Available Functions
- `Convert(t time.Time, fromZone string, toZone string) (time.Time, time.Time, error)`
- `GetUTCTime(t time.Time, fromZone string) (time.Time, error)`
- `TimezoneDifference(fromZone, toZone string) (time.Duration, error)`
- `EpochToTime(epoch int64) time.Time`
- `TimeToEpoch(t time.Time) int64`
- `ParseWithNearestFormat(timeStr string) (string, error)`
- `ListTimezones() []string`
- `ListFormats() map[string]string`

### 🤝 Contributing
Contributions are welcome! Feel free to submit pull requests and report issues.

### 📜 License
This project is licensed under the MIT License.