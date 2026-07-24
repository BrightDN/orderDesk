package mailer

import "bytes"

type Mail struct {
	Subject    string
	Receiver   string
	Body       string
	Attachment *Attachment
}

type Attachment struct {
	Filename string
	Reader   *bytes.Reader
}
