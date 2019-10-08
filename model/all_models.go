package model

import (
	"github.com/asdine/genji/record"
)

const (
	AccountTable        = "a"
	MailboxTable        = "mbx"
	MailboxMessageTable = "mbx_msg"
	MessageRawBodyTable = "msg_braw"
	MessageTable        = "msg"
	PropertyTable       = "prop"
	SnapshotTable       = "s"
)

var (
	Tables = map[string]record.Record{
		AccountTable:        new(Account),
		MailboxTable:        new(Mailbox),
		MailboxMessageTable: new(MailboxMessage),
		MessageRawBodyTable: new(MessageRawBody),
		MessageTable:        new(Message),
		PropertyTable:       new(Property),
		SnapshotTable:       new(Snapshot),
	}
)
