package testexample

type Mailer interface {
	Send(recipient, subject, content string) error
}
