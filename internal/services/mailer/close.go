package mailer

func (m *MailerService) Close() error {
	return m.Client.Close()
}
