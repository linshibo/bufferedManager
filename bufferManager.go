package bufferedManager

//import (
//"sync"
//)

const (
	MaxBucket = 2048
)

type Token struct {
	base  []byte
	Data  []byte
	owner chan<- *Token
}

func (t *Token) Return() {
	if t.owner != nil {
		t.Data = nil
		t.owner <- t
	}
}

type BufferManager struct {
	buffer chan *Token
	resource []byte
}

func NewBufferManager(size int) *BufferManager {
	ret := &BufferManager{buffer:make(chan *Token, size)}
	ret.resource=make([]byte,size*MaxBucket)
	for i := 0; i < size; i++ {
		ret.buffer <- &Token{owner: ret.buffer, base:ret.resource[i*MaxBucket:(i+1)*MaxBucket]}
	}
	return ret
}

func (b *BufferManager) GetToken(size int) *Token {
	var t *Token 
	select{
	case t = <-b.buffer:
	default :
		t = new(Token)
		t.base= make([]byte, size)
	}
	t.Data = t.base[:size]
	return t
}
