package mongo

import "fmt"

// tags
const (
	DEFAULT      = iota + 100 //100
	CONNECT                   //101
	INDEX_CREATE              //102
	SAVE                      //103
	COUNT                     //104
	UPDATE                    //105
	DELETE                    //106
	GET_CURSOR                //107
	PAGINATE                  //108
)

type Error struct {
	Code int
	Msg  string
}

func errorsCode(tag int) string {
	switch tag {
	case DEFAULT:
		return "err@"
	case CONNECT:
		return "connect@"
	case INDEX_CREATE:
		return "index_create@"
	case SAVE:
		return "save@"
	case UPDATE:
		return "update@"
	case DELETE:
		return "delete@"
	case PAGINATE:
		return "paginate@"
	default:
		return "err"

	}
}

func (m *Mongo) handleErrors(collection string, tag int, err error) Error {
	e := Error{
		Code: 0,
		Msg:  "Ok",
	}
	if err != nil {
		e.Code = tag
		e.Msg = err.Error()
		fmt.Println("err=", e.Msg)
		key := errorsCode(tag) + collection
		logErrors(m, key, err.Error())
	}

	return e
}

func logErrors(m *Mongo, key, value string) {
	m.logger.Set(key, value)
}
