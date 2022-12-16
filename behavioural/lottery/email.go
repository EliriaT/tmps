package lottery

import "fmt"

type Email struct {
}

func (s *Email) genRandomCode(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random code %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveCodeCache(otp string) {
	fmt.Printf("EMAIL: saving code: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "EMAIL Code for lottery is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}
