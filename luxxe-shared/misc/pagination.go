package misc

type PaginationStruct struct {
	Page  int
	Limit int
}

func Pagination(pgs PaginationStruct) (int, int) {
	page, limit := pgs.Page, pgs.Limit

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}

	skip := 0

	// If the current page is greater than 1, the function calculates how many items to skip. 
	// For example, if you are on page 3 with a limit of 20 items per page, you need to skip the first 40 items
	if page > 1 {
		skip = (page - 1) * limit
	}

	return skip, limit
}
