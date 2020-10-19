package interfaces

import "net/mail"

type Sender interface {
	// Send a mail to a given address with a subject and text.
	Send(message mail.Message) error
  }