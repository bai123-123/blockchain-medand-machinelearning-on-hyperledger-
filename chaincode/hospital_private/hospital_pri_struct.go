package main

type EMR_pri struct {
	EMRNo string

	Date string

	DoctorNo string
	Doctor string
	Medical_department string
	PCD string
	PMH string
	DD string
	Medicine string
	Quantity string
	Amount string
	PPatient Patient

}

type Patient struct {
	Name string
	Gender string
	Age string
	BirthDate string
	EmergencyContact string
	Address string
	IdCardNo string
	MartialState string
	ContactNumber string
	Email string
}





