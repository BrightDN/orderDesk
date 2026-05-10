package mailer

type Mail struct {
	Subject     string
	Receiver    string
	Sender      string
	Body        string
	Attachments []string
}
