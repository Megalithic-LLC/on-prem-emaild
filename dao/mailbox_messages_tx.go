package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
)

func (self MailboxMessagesDAO) CreateTx(tx *genji.Tx, mailboxMessage *model.MailboxMessage) error {
	if table, err := tx.GetTable(model.MailboxMessageTable); err != nil {
		return err
	} else {
		_, err := table.Insert(mailboxMessage)
		return err
	}
}

func (self MailboxMessagesDAO) FindTx(tx *genji.Tx, where query.Expr, limit int, iter func(mailboxMessage *model.MailboxMessage) error) error {
	mailboxMessageTable, err := tx.GetTable(model.MailboxMessageTable)
	if err != nil {
		return err
	}
	selectStmt := query.Select().From(mailboxMessageTable)
	if where != nil {
		selectStmt = selectStmt.Where(where)
	}
	if limit > 0 {
		selectStmt = selectStmt.Limit(limit)
	}
	return selectStmt.Run(tx).Iterate(func(recordID []byte, r record.Record) error {
		var mailboxMessage model.MailboxMessage
		if err := mailboxMessage.ScanRecord(r); err != nil {
			return err
		}
		return iter(&mailboxMessage)
	})
}

func (self MailboxMessagesDAO) FindByIDsTx(tx *genji.Tx, mailboxID, messageID string) (*model.MailboxMessage, error) {
	messageTable, err := tx.GetTable(model.MailboxMessageTable)
	if err != nil {
		return nil, err
	}
	searchFor := &model.MailboxMessage{MailboxID: mailboxID, MessageID: messageID}
	mailboxMessagePK, err := searchFor.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := messageTable.GetRecord(mailboxMessagePK)
	if err != nil {
		return nil, err
	}
	var mailboxMessage model.MailboxMessage
	if err := mailboxMessage.ScanRecord(r); err != nil {
		return nil, err
	}
	return &mailboxMessage, nil
}

func (self MailboxMessagesDAO) ReplaceTx(tx *genji.Tx, mailboxMessage *model.MailboxMessage) error {
	mailboxMessageTable, err := tx.GetTable(model.MailboxMessageTable)
	if err != nil {
		return err
	}
	pk, err := mailboxMessage.PrimaryKey()
	if err != nil {
		return err
	}
	return mailboxMessageTable.Replace(pk, mailboxMessage)
}