package user

type GetUserReqest struct {
	ID int `json:"id"`
}
type GetListUserReqest struct {
	Word    string `json:"word"`
	LoginID string `json:"login_id"`
}
