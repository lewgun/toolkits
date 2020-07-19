

type BasicInfo struct {
	Title string
	SubTitle string
	ISBN string
	Translator string  //译者
	Image string
	Author string
	OriginAuthor string  //原作者
	OriginTitle string //原作名
	Publisher string   //出版社
	PublishAt string
	PageNum string ``
	Price string
	Binding string //装帧
	Producer string //出品方
	Score string
	Series string //从书

}

type ReadInfo struct {
	Reading uint32  //在读
	Readed uint32  //已读
	WantRead uint32 //想读
}

type SaleInfo struct {
	Vendor string
	Link string
}
type Book struct {
	BasicInfo
	ReadInfo

	AboutAuthor  string     //作者简介
	Introduction string     //内容简介
	Catalog      string     //目录
	Tags         []string   //标签
	OnSales      []SaleInfo //当前在售
	Recommends   []string   //相关推荐
}




const doubanAPI = "http://www.douban.com/isbn/"
func Query (isbn string) (*Book, error) {
	url :=  doubanAPI + isbn
	//fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("Error reading request. " + err.Error())
	}

	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		panic("Error reading response. " + err.Error())

	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	book := &Book{}

	book.Title =  doc.Find("#wrapper > h1 > span").Text()
	book.Image, _ =  doc.Find("#mainpic > a > img").Attr("src")
	book.Score =  doc.Find("#interest_sectl > div > div.rating_self > strong").Text()

	//内容简介
	//book.Introduction = strings.TrimSpace(doc.Find("#link-report > div > div > p").Text())
	extractIntroduction(doc, book)

	//推荐
	doc.Find("#db-rec-section").Find("dd > a") .Each(func(i int, selection *goquery.Selection) {
		book.Recommends = append(book.Recommends,strings.TrimSpace(selection.Text()) )
	})

	//作者简介
	extractAboutAuthor(doc,book)

	//标签
	extractTags(doc, book)

	//想读
	extractReadInfo(doc, book)

	//销售信息
	extractSaleInfo(doc, book)

	//目录
	extractCatalog(doc, book)

	//基本信息
	extractBasicInfo(doc, book)

	return book, nil
}

func extractIntroduction(doc *goquery.Document, book *Book) {

	relatedInfo := doc.Find("div.related_info div.indent")

	intro := ""
	var content *goquery.Selection
	if relatedInfo.Eq(0).Find("span").Nodes == nil {
		content = relatedInfo.Eq(0).Find("div.intro")

	} else {
		// 有些简介需点击展开全部查看完整内容
		content = relatedInfo.Eq(1).Find("div.intro").Eq(1)
	}

	content.Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Find("p").Text())
		intro = fmt.Sprintf("%s%s\n", intro, text)
	})
	book.Introduction = intro

}

//https://blog.csdn.net/luckytanggu/article/details/79684470
func extractBasicInfo(doc *goquery.Document, book *Book) {
	// 书籍基本信息(一整块的信息)
	info := doc.Find("div#info")

	// 把全角/半角冒号统一替换为半角冒号, 因为发现有些书籍信息会用全角冒号，例如--> 译者：xxx
	// 影响到下面一些字符的判断
	lines := strings.Split(strings.Replace(info.Text(), ":|：", ":", -1), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && strings.ContainsAny(line, ":") {
			line = fmt.Sprintf("%s ", line)
			for _, nextLine := range lines[i+1:] {
				nextLine = strings.TrimSpace(nextLine)
				if nextLine == "" {
					continue
				}
				if strings.ContainsAny(nextLine, ":") {
					break
				} else {
					line = fmt.Sprintf("%s%s", line, nextLine)
				}
			}
			//fmt.Println(line)
			val := strings.SplitN(line, ":", 2)[1]
			switch strings.SplitN(line, ":", 2)[0] {
			case "作者":
				book.Author = val

			case "出版社":
				book.Producer = val

			case "原作名":
				book.OriginTitle = val

			case "原作者":
				book.OriginAuthor = val

			case "译者":
				book.Translator = val

			case "出版年":
				book.PublishAt = val

			case "页数":
				book.PageNum = val

			case "定价":
				re := regexp.MustCompile("[0-9.]+")

				fmt.Println()
				fields := re.FindAllString(val, -1)
				if len(fields ) > 0  {
					book.Price = fields[0]
				}


			case "装帧":
				book.Binding = val

			case "丛书":
				book.Series = val

			case "ISBN":
				book.ISBN = val

			case "副标题":
				book.SubTitle = val
			}
		}
	}

}

const (
	reading = "人在读"
	readed = "人读过"
	wantRead = "人想读"
)
func extractReadInfo( doc *goquery.Document, book *Book) {
	doc.Find("#collector > p > a").Each(func(i int, selection *goquery.Selection) {
		txt :=  strings.TrimSpace(selection.Text())

		var (
			p *uint32
			trimed string
		)
		if strings.HasSuffix(txt, reading) {
			p = &book.Reading
			trimed = strings.TrimSuffix(txt, reading)

		} else 	if strings.HasSuffix(txt, readed) {
			p = &book.Readed
			trimed = strings.TrimSuffix(txt, readed)

		} else if strings.HasSuffix(txt, wantRead) {
			p = &book.WantRead
			trimed = strings.TrimSuffix(txt, wantRead)
		}

		num, _ := strconv.ParseUint(trimed, 10,64)
		*p = uint32(num)
	})

}

func extractSaleInfo(doc *goquery.Document, book *Book) {
	doc.Find("#buyinfo-printed").Find(" div.vendor-name > a").Each(func(i int, selection *goquery.Selection) {
		link, _ := selection.Attr("href")
		book.OnSales = append(book.OnSales, SaleInfo{
			Vendor: strings.TrimSpace(selection.Text()),
			Link:   link,
		})
	})
}

//https://blog.csdn.net/luckytanggu/java/article/details/79684470
func extractAboutAuthor(doc *goquery.Document, book *Book) {
	////作者简介
	//doc.Find("span").EachWithBreak(func(i int, s *goquery.Selection) bool {
	//	if s.Text() == "作者简介"{
	//		book.AboutAuthor = strings.TrimSpace(s.ParentFiltered("h2").NextFiltered("div").Find("div > div.intro").Text())
	//		return false
	//	}
	//	return true
	//})

	relatedInfo := doc.Find("div.related_info div.indent")
	aboutAuthor := ""
	var author *goquery.Selection
	if relatedInfo.Eq(1).Find("span").Nodes == nil {
		author = relatedInfo.Eq(1).Find("div.intro")

	} else {
		author = relatedInfo.Eq(1).Find("div.intro").Eq(1)
	}
	author.Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Find("p").Text())
		aboutAuthor = fmt.Sprintf("%s%s\n", aboutAuthor, text)
	})
	book.AboutAuthor = aboutAuthor

}

func extractTags(doc *goquery.Document, book *Book) {
	doc.Find("#db-tags-section > div.indent > span").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		//最多20个标签
		if i < 20 {
			book.Tags = append(book.Tags, strings.TrimSpace(selection.Text()))
			return true
		}
		return false
	})
}

func extractCatalog(doc *goquery.Document, book *Book) {
	doc.Find("span").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if s.Text() == "目录"{
			book.Catalog = strings.TrimSpace(s.ParentFiltered("h2").NextFiltered("div").Text())
			return false
		}
		return true
	})

}
