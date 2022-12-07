package models

type CaseStatusType string

const (
	CasePending  CaseStatusType = "pending"
	CaseBuilding CaseStatusType = "building"
	CaseBuilt    CaseStatusType = "finished"
)

type Case struct {
	Id          string
	MobileModel string
	Status      CaseStatusType
}
