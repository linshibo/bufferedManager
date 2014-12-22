package bufferedManager

//import (
//"sync"
//)

type Token struct {
	base [2048]byte
	Data  []byte
	owner chan<- *Token
}

func (t *Token) Return() {
	t.Data = nil
	t.owner <- t
}

type BufferManager struct {
	buffer chan *Token
}

func NewBufferManager(size int) *BufferManager {
	buffer := make(chan *Token, size)
	ret := &BufferManager{buffer}
	for i := 0; i < size; i++ {
		buffer <- &Token{owner: buffer}
	}
	return ret
}

func (b *BufferManager) GetToken(size int) *Token {
	t := <-b.buffer
	t.Data = t.base[:size]
	return t
}
