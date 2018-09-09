package types

type Reminder interface {
	Remind() (out string)
}

type Server interface {
	Run() (err error)
}
