package main

type UserInfo struct {
	Name string
	EMRNo string
	Category string
	MaxNum string
	CurrentNum string
	Historys	[]HistoryItem
}
type HistoryItem struct {
	TxId	string
	User     UserInfo
}