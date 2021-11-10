package userTABLE

var USERDATA map[string]string

func GETDATA(s string) (string, bool) {
	USERDATA = map[string]string{
		"test":"test",
	}
	if _, exist := USERDATA[s]; exist {
		return USERDATA[s], true
	} else {
		return "", false
	}
}
