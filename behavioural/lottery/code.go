package lottery

type ICode interface {
	genRandomCode(int) string
	saveCodeCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Code struct {
	InterfCode ICode
}

func (o *Code) GenAndSendCode(otpLength int) error {
	otp := o.InterfCode.genRandomCode(otpLength)
	o.InterfCode.saveCodeCache(otp)
	message := o.InterfCode.getMessage(otp)
	err := o.InterfCode.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
