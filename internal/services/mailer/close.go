package mailer

func (m *MailerService) Close() error {
	return m.client.Close()
}
