package contest

type MyMutex struct {
	lockChan chan struct{}
}

func New() Mutex {
	l := make(chan struct{}, 1)
	l <- struct{}{}
	return &MyMutex{lockChan: l}
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
