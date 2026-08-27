package main

import (
	"fmt"
	"time"
)

func main() {
	// Remember: The official release date and time of the Go programming language is used 
	// to get a formatted version of the date and time, which is "2006-01-02 15:04:05".

	currentDateTime := time.Now()
	fmt.Println("Current Date and Time:", currentDateTime)
	fmt.Printf("Data Type: %T\n", currentDateTime)

	// Format the time as "2006-01-02 15:04:05" or "yyyy-mm-dd hh:mm:ss"
	formattedTime := currentDateTime.Format("2006-01-02, Monday 15:04:05")
	fmt.Println("Formatted time:", formattedTime)

	// Format the time as "02 January 2006, Monday 03:04 PM" or "dd month yyyy, weekday hh:mm AM/PM"
	formattedTime = currentDateTime.Format("02 January 2006, Monday 03:04 PM")
	fmt.Println("Formatted time:", formattedTime)

	// Get the current date
	currentDate := currentDateTime.Format("2006-01-02")
	fmt.Println("Current date:", currentDate)

	// Get the current time in 24-hour format
	currentTime := currentDateTime.Format("15:04:05")
	fmt.Println("Current time:", currentTime)

	// Get the current time in 12-hour format with seconds
	currentTime = currentDateTime.Format("03:04:05 PM")
	fmt.Println("Current time:", currentTime)

	// Get the current time in 12-hour format without seconds
	currentTime = currentDateTime.Format("03:04 PM")
	fmt.Println("Current time:", currentTime)

	// Get the current day of the week
	currentDay := currentDateTime.Format("Monday")
	fmt.Println("Current day:", currentDay)

	// Get the year, month, day, hour, minute, second and weekday from the current time
	year := currentDateTime.Year()
	month := currentDateTime.Month()
	day := currentDateTime.Day()
	hour := currentDateTime.Hour()
	minute := currentDateTime.Minute()
	second := currentDateTime.Second()
	weekday := currentDateTime.Weekday()
	fmt.Printf("Year: %d, Month: %s, Day: %d, Hour: %d, Minute: %d, Second: %d, Weekday: %s\n", year, month, day, hour, minute, second, weekday	)

	// String to time conversion
	str := "2023-04-05 12:34:56"
	// Remember: Both str and format should be in the same format
	format := "2006-01-02 15:04:05"
	dateTime, _ := time.Parse(format, str)
	fmt.Println("Date and time:", dateTime)
	formattedTime = dateTime.Format("02 January 2006, Monday 03:04 PM")
	fmt.Println("Formatted time:", formattedTime)

	// Update the date
	updatedDate := dateTime.AddDate(0, 0, 1)
	fmt.Println("Updated date (+1 day):", updatedDate)

	// Update the time
	updatedTime := dateTime.Add(time.Hour * 2)
	fmt.Println("Updated time (+2 hours):", updatedTime)
}
