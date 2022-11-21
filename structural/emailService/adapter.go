package emailService

// old email service that can't be changed
type OldEmailSender interface {
	Send(from, to, subject, body string) error
}

type OldEmailService struct{}

func (s OldEmailService) Send(from, to, subject, body string) error {
	//Send email...
	return nil
}

// new email service which has states
// NewEmailSender interface supported by the client
type NewEmailSender interface {
	Send() error
}

// NewEmailService has states such as From, To, Subject, Body
type NewEmailService struct {
	From, To, Subject, Body string
}

// Send implements sendings of emails
func (s NewEmailService) Send() error {
	//state already initialized just send
	return nil
}

// OldEmailServiceAdapter is an adapter for the old email service, and meets the NewEmailSender interface
type OldEmailServiceAdapter struct {
	From, To, Subject, Body string
	OldEmailService         OldEmailSender
}

// Send of the OldEmailServiceAdapter calls the Send() of the OldEmailService
func (a OldEmailServiceAdapter) Send() error {
	return a.OldEmailService.Send(a.From, a.To, a.Subject, a.Body)
}

/*
    client which only supports the new email service interface
	EmailClient{ Mail: OldEmailServiceAdapter{...}} old service
	EmailClient{ Mail: NewEmailService{...}} new service
*/
// EmailClient  only supports the new email service interface
type EmailClient struct {
	Mail NewEmailSender
}

func (e EmailClient) SendEmail(From, To, Subject, Body string) error {
	return e.Mail.Send()
}

func NewEmailClient(emailSender NewEmailSender) EmailClient {
	return EmailClient{Mail: emailSender}
}
