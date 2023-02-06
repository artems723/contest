package contest

type MyMutex struct {
	lockChan chan struct{}
}

func New() Mutex {
	l := make(chan struct{}, 1)
	l <- struct{}{}
	m := MyMutex{
		lockChan: l,
	}
	return &m
}

func (m *MyMutex) Lock() {
	<-m.lockChan
}

func (m *MyMutex) Unlock() {
	m.lockChan <- struct{}{}
}

func (m *MyMutex) LockChannel() <-chan struct{} {
	return m.lockChan
}
