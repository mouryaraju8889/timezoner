package timezone

import (
	"context"
	"fmt"
	"time"
)


// StreamCurrentTime starts a goroutine that sends the current system time 
// to the provided channel every second. The function listens to a context 
// for cancellation and stops gracefully when the context is canceled.
//
// Parameters:
//   - ctx: context.Context used to cancel the streaming goroutine.
//   - ch: a send-only channel where the time.Time values will be sent.
//
// Output:
//   - Sends the current time every second into the channel until the context is done
func StreamCurrentTime(ctx context.Context, ch chan<- time.Time) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine: received stop signal.")
			close(ch) // Cleanup
			return
		case t := <-ticker.C:
			// fmt.Println("Tick at:", time.Now())
			ch <- t
			i++
		}
	}
}



// StreamCurrentTimeInZone starts a goroutine that sends the current time
// in a specified timezone to the provided channel every second. It uses 
// a ticker and stops when the provided context is canceled.
//
// Parameters:
//   - ctx: context.Context used to stop the goroutine.
//   - ch: a send-only channel that receives time.Time values each second.
//   - timezoneName: a string representing the IANA timezone (e.g., "Asia/Kolkata", "America/New_York").
//
// Output:
//   - Sends the current time (converted to the given timezone) to the channel every second.
//   - If the timezone is invalid, prints an error and closes the channel immediately
func StreamCurrentTimeInZone(ctx context.Context, ch chan<- time.Time, timezoneName string) {
	loc, err := time.LoadLocation(timezoneName)
	if err != nil {
		fmt.Printf("Invalid timezone: %v\n", err)
		close(ch)
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		case t := <-ticker.C:
			ch <- t.In(loc)
		}
	}
}
