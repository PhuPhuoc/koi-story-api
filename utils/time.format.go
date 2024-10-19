package utils

import "time"

func CreateDateTimeCurrentFormated() string {
	// Use time.FixedZone to create a fixed offset for Vietnam time (UTC+7)
	vietnamZone := time.FixedZone("Asia/Ho_Chi_Minh", 7*60*60)
	currentTime := time.Now().In(vietnamZone)
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	return formattedTime
}
