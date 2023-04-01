package model

type Group struct {
	Id         uint
	MaxCurrent float32
}

type ChargePoint struct {
	Id      uint
	GroupId uint
}

type ChargePointConnector struct {
	Id            uint
	ChargePointId uint
	Status        string
}
