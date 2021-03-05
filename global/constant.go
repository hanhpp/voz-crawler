package global

const (
	VozBaseURL    = "https://voz.vn"
	ThreadLink    = ".structItem-minor .structItem-parts"
	ThreadTitle   = ".structItem-title"
	ThreadStruct = ".structItem.structItem--thread"
	ThreadIDRegex = "(.[0-9]\\w+/)"

	CommentIDRegex   = "([0-9]\\w+)"
	CommentStruct    = ".message-userContent"
	//CommentStruct    = ".message-userContent .bbWrapper"
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
	F33           = "https://voz.vn/f/diem-bao.33"
)

const (
	MinPage uint64 = 1
	MaxPage uint64 = 10
)

const (
	//10 second
	CrawlInterval uint64 = 10
)
var (
	//F17_P2 = utils.AddPageSuffix(F17, 2)
	//F17_P3 = utils.AddPageSuffix(F17, 3)
	//F17_P4 = utils.AddPageSuffix(F17, 4)
	//F17_P5 = utils.AddPageSuffix(F17, 5)
	F17_Pages = []string{F17}//,F17_P2,F17_P3}//,F17_P4,F17_P5}
)

var (
	//F33_P2 = utils.AddPageSuffix(F33, 2)
	//F33_P3 = utils.AddPageSuffix(F33, 3)
	//F33_P4 = utils.AddPageSuffix(F33, 4)
	//F33_P5 = utils.AddPageSuffix(F33, 5)
	F33_Pages = []string{F33}//,F33_P2,F33_P3}//,F33_P4,F33_P5}
)

