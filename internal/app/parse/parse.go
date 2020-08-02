package parse

type Parse struct {
	url string
}

func New() *Parse {
	return &Parse{
		url: "https://www.goodfon.ru/catalog/landscapes/720x1280/index-%d.html",
	}
}
