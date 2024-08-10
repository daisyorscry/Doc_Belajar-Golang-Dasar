package helper

import "log"

func DatabaseErr(err error) error {
	if err != nil {
		log.Panicf("Database error: %v", err)
		return err
	}

	return nil
}
