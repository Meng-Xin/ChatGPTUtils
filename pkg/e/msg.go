package e

var ErrMsg = map[int]string{
	SUCCESS: "Success",
	ERROR:   "Fail",
}

func GetMsg(code int) string {
	return ErrMsg[code]
}
