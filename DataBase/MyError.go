package DataBase

type Error struct {
	Message string
}

func (E *Error) Error() string {
	return E.Message
}

type Msg struct {
	Message string
}

func (M *Msg) Error() string {
	return M.Message
}
