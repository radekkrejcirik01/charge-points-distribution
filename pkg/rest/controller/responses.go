package controller

type Response struct {
	Status  string `json:""`
	Message string `json:",omitempty"`
}
type ResponseGet struct {
	Status  string `json:""`
	Message string `json:",omitempty"`
	Data    string `json:",omitempty"`
}
