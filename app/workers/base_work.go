package workers

//BaseWork BaseWork
type BaseWork interface {
	Init() error
	Status()
	Run()
	Stop()
	Close()
}