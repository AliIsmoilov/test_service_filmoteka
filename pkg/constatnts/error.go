package constatnts

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	ErrRowsAffectedZero = Sentinel("no rows affected by sql command")
	ErrRecordNotFound   = Sentinel("record not found")
)
