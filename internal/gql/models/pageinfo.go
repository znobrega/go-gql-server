package models

type PageInfo struct {
	StartCursor   *string `json:"startCursor"`
	EndCursor     *string `json:"EndCursor"`
	HasNextPage   *string `json:"HasNextPage"`
	HasBeforePage *string `json:"HasBeforePage"`
	BeforeCursor  *string `json:"beforeCursor"`
	NextCursor    *string `json:"nextCursor"`
}
