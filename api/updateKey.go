package api

import "time"

func UpdateKey(token, contentHash string) error {
	time, err := time.ParseDuration("24h")
	if err != nil {
		return err
	}

	_, err = s.PublishWithDetails(contentHash, token, time, 10000, true)

	return err
}
