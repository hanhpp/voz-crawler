package global

const (
	F17           = "https://voz.vn/f/chuyen-tro-linh-tinh.17/"
	F33           = "https://voz.vn/f/diem-bao.33/"
	ThreadLink    = ".structItem-minor .structItem-parts"
	ThreadTitle   = ".structItem-title"
	ThreadIDRegex = "(.[0-9]\\w+/)"
	//With positive look ahead and behind
	//ThreadIDRegex = "(?=.)([0-9]\\w+)(?=/)"
	//Not working in go
	//https://stackoverflow.com/questions/26771592/negative-look-ahead-in-go-regular-expressions
)
