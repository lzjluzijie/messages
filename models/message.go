package models

import "time"

type Message struct {
	ID        int64 `xorm:"pk autoincr"`
	Text      string
	Name      string
	IP        string
	CreatedAt time.Time `xorm:"created"`
}

func InsertMessage(text, name, ip string) (err error) {
	m := &Message{
		Text: text,
		Name: name,
		IP:   ip,
	}

	_, err = x.Insert(m)

	return
}

func GetMessage() (msgs []Message, err error) {
	msgs = make([]Message, 0)
	err = x.Find(&msgs)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
