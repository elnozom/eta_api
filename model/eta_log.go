package model

type Log struct {
	InternalID   string
	SubmissionID string
	StoreCode    int
	Serials      string
	LogText      string
	ErrText      string
	Posted       bool
}
