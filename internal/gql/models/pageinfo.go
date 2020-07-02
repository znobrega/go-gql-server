package models

type PageInfo struct {
	StartCursor   *string `json:"startCursor"`
	EndCursor     *string `json:"endCursor"`
	HasNextPage   bool    `json:"hasNextPage"`
	HasBeforePage bool    `json:"hasBeforePage"`
	BeforeCursor  *string `json:"beforeCursor"`
	NextCursor    *string `json:"nextCursor"`
}

func HasPage(cursor *string) bool {
	if cursor == nil {
		return false
	}

	return true
}
