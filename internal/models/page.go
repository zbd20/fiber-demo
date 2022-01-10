package models

type Page struct {
	PageSize   int    `json:"limit"`
	Offset     int    `json:"offset"`
	Page       int    `json:"page"`
	TotalCount int    `json:"-"`
	Query      string `json:"-"`
	OrderBy    string `json:"order_by"`
	Sort       string `json:"sort"`
}

type Result struct {
	TotalCount  *int        `json:"total_count,omitempty"`
	PageCount   *int        `json:"page_count,omitempty"`
	CurrentPage *int        `json:"current_page,omitempty"`
	PageSize    *int        `json:"page_size,omitempty"`
	Data        interface{} `json:"data"`
	Code        int64       `json:"code"`
	Message     string      `json:"message"`
}

func NewResult(count int, page *Page, code int64, msg string, data interface{}) Result {
	var result Result

	if page != nil {
		result = Result{
			TotalCount:  &count,
			CurrentPage: &page.Page,
			PageSize:    &page.PageSize,
			Data:        data,
			Message:     msg,
			Code:        code,
		}

		pc := count / page.PageSize
		result.PageCount = &pc
		if count%page.PageSize > 0 {
			*(result.PageCount) += 1
		}
		// 处理跳转至随机页，该情况下，currentPage 不为 1，开始模糊查询的问题
		if *result.CurrentPage > *result.PageCount {
			*result.CurrentPage = 1
		}
	} else {
		result = Result{
			Data:    data,
			Message: msg,
			Code:    code,
		}

	}

	return result
}
