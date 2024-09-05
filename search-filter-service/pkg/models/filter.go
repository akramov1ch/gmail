package models

import "time"

type Email struct {
    EmailId  string
    Sender   string
    Subject  string
    Preview  string
    DateSent time.Time
}

type FilteredEmail struct {
    EmailId string
    Folder  string
}
