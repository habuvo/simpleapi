package models

//Purchase is data type for purchase
type Purchase struct {
	ID                int64
	AffectingContract int64
	CreditsSpent      uint
}
