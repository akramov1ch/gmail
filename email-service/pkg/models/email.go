package models

type Email struct {
    ID         int64
    Sender     string
    Recipients []string
    CC         []string
    BCC        []string
    Subject    string
    Body       string
    Attachments []string
    CreatedAt  string
    IsTrashed  bool
}
