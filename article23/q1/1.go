package main

func test() {
	lock.Lock()
	for mailbox == 1 {
		sendCond.Wait()
	}
	mailbox = 1
	lock.Unlock()
	recvCond.Signal()

}

func ceshi() {
	lock.RLock()
	for mailbox == 0 {
		recvCond.Wait()
	}
	mailbox = 0
	lock.RUnlock()
	sendCond.Signal()

}
