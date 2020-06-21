package dto

type Product struct {
	Title       string `json:"name"`
	Thumbnail   string `json:"imageURL"`
	Reviews     string `json':"totalReviews"`
	Description string `json':"Description"`
}

type CrawlInfo struct {
	URL         string  `json:"Url"`
	Crawltime   string  `json:"CrawlTime"`
	ProductInfo Product `json:"Product"`
}
