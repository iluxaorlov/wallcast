package parse

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/iluxaorlov/wallcast/internal/app/model"
	"github.com/iluxaorlov/wallcast/internal/app/store"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func (p *Parse) GetImage(store *store.Store) (*model.Image, error) {
	for {
		rand.Seed(time.Now().UnixNano())

		res, err := http.Get(fmt.Sprintf(p.url, rand.Intn(9999)))
		if err != nil {
			return nil, err
		}

		if res.StatusCode != http.StatusOK {
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}

		sel := doc.Find(".wallpapers__item__img")

		for i := range sel.Nodes {
			src, _ := sel.Eq(i).Attr("src")
			src = strings.Replace(src, "wallpaper", "original", 1)
			src = strings.Replace(src, "big", "720x1280", 1)

			image, err := store.Image().FindByUrl(src)
			if err != nil {
				return nil, err
			}

			if image != nil {
				continue
			}

			res, err := http.Get(src)
			if err != nil {
				return nil, err
			}

			if res.StatusCode != http.StatusOK {
				continue
			}

			image = &model.Image{
				Url: src,
			}

			return image, nil
		}
	}
}
