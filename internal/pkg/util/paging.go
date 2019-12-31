package util

import (
	"strconv"
)

type PaginationRequest struct {
	CurrentPage int `json:"currentPage"`
	Limit       int `json:"limit"`
	Offset      int `json:"offset"`
}

type Pagination struct {
	PaginationRequest
	TotalPage  int `json:"totalPage"`
	TotalItems int `json:"totalItems"`
}
type PaginationList struct {
	Pagination
	Items interface{} `json:"items"`
}

func NewPaginationRequest(limitRequest string, pageRequest string) PaginationRequest {
	currentPage := 1
	limit := 10

	if pageRequest != "" {
		pageInt, err := strconv.Atoi(pageRequest)
		if err == nil {
			currentPage = pageInt
		}
	}

	if limitRequest != "" {
		limitInt, err := strconv.Atoi(limitRequest)
		if err == nil {
			limit = limitInt
		}
	}

	paginationRequest := PaginationRequest{}
	paginationRequest.CurrentPage = currentPage
	paginationRequest.Offset = limit * (currentPage - 1)
	paginationRequest.Limit = limit

	return paginationRequest
}

//func (pagination *PaginationRequest) PaginationFilter(q *orm.Query) (*orm.Query, error) {
//	if pagination.Limit != - 1 {
//		q.Limit(pagination.Limit)
//	}
//	q.Offset(pagination.Offset)
//	return q, nil
//}

func ListPaginate(list []interface{}, skip int, size int) interface{} {
	limit := func() int {
		if skip+size > len(list) {
			return len(list)
		} else {
			return skip + size
		}

	}

	start := func() int {
		if skip > len(list) {
			return len(list)
		} else {
			return skip
		}

	}
	return list[start():limit()]
}

func PaginationResponse(paginationRequest PaginationRequest, totalItems int, list interface{}) PaginationList {

	paginationList := PaginationList{}
	totalPage := totalItems / paginationRequest.Limit
	if totalItems > (totalPage * paginationRequest.Limit) {
		totalPage += 1
	}

	paginationList.TotalItems = totalItems
	paginationList.Items = list
	paginationList.CurrentPage = paginationRequest.CurrentPage
	paginationList.Limit = paginationRequest.Limit
	paginationList.Offset = paginationRequest.Offset
	paginationList.TotalPage = totalPage
	return paginationList
}
