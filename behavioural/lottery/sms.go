package lottery

import "fmt"

type Sms struct {
}

func (s *Sms) genRandomCode(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random code %s\n", randomOTP)
	return randomOTP
}

func (s *Sms) saveCodeCache(otp string) {
	fmt.Printf("SMS: saving code: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS code for lottery is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
