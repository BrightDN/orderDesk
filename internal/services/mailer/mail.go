package mailer

type Mail struct {
	Subject     string
	Receiver    string
	Body        string
	Attachments []string
}
