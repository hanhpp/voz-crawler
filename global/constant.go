package global

import "voz/utils"

const (
	F33           = "https://voz.vn/f/diem-bao.33"
	VozBaseURL    = "https://voz.vn"
	ThreadLink    = ".structItem-minor .structItem-parts"
	ThreadTitle   = ".structItem-title"
	ThreadIDRegex = "(.[0-9]\\w+/)"

	CommentIDRegex   = "([0-9]\\w+)"
	CommentStruct    = ".message-userContent"
	CommentNamespace = "data-lb-caption-desc"
	CommentId        = "data-lb-id"
	//TestThread       = "https://voz.vn/t/nsut-thanh-loc-toi-phai-ban-do-trong-nha-de-trang-trai-cuoc-song.196903/"
	//TestThread = "https://voz.vn/t/hai-lan-tuyet-vong-cua-chang-beo-tap-the-hinh.203141/"
	TestThread = "https://voz.vn/t/chi-tien-ti-co-ngay-ngoi-hoa-hau.203081/"
	//With positive look ahead and behind
	//ThreadIDRegex = "(?=.)([0-9]\\w+)(?=/)"
	//Not working in go
	//https://stackoverflow.com/questions/26771592/negative-look-ahead-in-go-regular-expressions
)

const (
	F17 = "https://voz.vn/f/chuyen-tro-linh-tinh.17"
)

var (
	F17_P2 = utils.AddPageSuffix(F17, 2)
	F17_P3 = utils.AddPageSuffix(F17, 3)
	F17_P4 = utils.AddPageSuffix(F17, 4)
	F17_P5 = utils.AddPageSuffix(F17, 5)
	F17_P6 = utils.AddPageSuffix(F17, 6)
	F17_P7 = utils.AddPageSuffix(F17, 7)
)
