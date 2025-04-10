package utils

import (
    "math"
    "time"
)

func DaysSince(date time.Time) int {
    now := time.Now()
    duration := now.Sub(date)
    return int(duration.Hours() / 24)
}

func CalculateBalance(initialAmount float64, daysPassed int, dailyInterestRate float64) float64 {
    balance := initialAmount * math.Pow(1+dailyInterestRate, float64(daysPassed))
    return balance
}