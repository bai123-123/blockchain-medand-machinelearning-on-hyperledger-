package main

type EMR_common struct {
	EMRNo string
	MedNo string
	DoctorNo string
	POM string
	Date string
	CommonInfo Common
	VerifyDate string
}

type Common struct {
	Name string
	Gender string
	Age string
	BirthDate string
	Contact string
	Medicine string
	QuantityFixed string
	AmountFixed string
	AmountCurrent string
}
