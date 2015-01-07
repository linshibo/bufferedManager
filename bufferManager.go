package bufferedManager

//const (
//MaxBucket  = 1024
//)

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
	buffer   chan *Token
	resource []byte
	MaxBucket int
}

func NewBufferManager(size int, MaxBucket int) *BufferManager {
	ret := &BufferManager{buffer: make(chan *Token, size)}
	ret.resource = make([]byte, size*MaxBucket)
	ret.MaxBucket=MaxBucket
	for i := 0; i < size; i++ {
		ret.buffer <- &Token{owner: ret.buffer, base: ret.resource[i*MaxBucket : (i+1)*MaxBucket]}
	}
	return ret
}

func (b *BufferManager) GetToken(size int) *Token {
	var t *Token
	if size > b.MaxBucket {
		t = new(Token)
		t.base = make([]byte, size)
	} else {
		select {
		case t = <-b.buffer:
		default:
			t = new(Token)
			t.base = make([]byte, size)
		}
	}
	t.Data = t.base[:size]
	return t
}
