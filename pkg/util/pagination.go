package util

// Page :
type Page struct {
	Start int `json:"start"`
	No    int `json:"no"`
	Size  int `json:"size"`
	Total int `json:"total"`
	Pages int `json:"pages"`
}

// Cal : calculate
func (p *Page) Cal() {
	if p.No < 0 || p.Size <= 0 {
		return
	}

	p.Pages = p.Total / p.Size
	if p.Total%p.Size > 0 {
		p.Pages++
	}
	p.Start = p.No*p.Size - p.Size
}
