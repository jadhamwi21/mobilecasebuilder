package models

type CaseStatus int

const (
	CasePending CaseStatus = iota
	CaseBuilding
	CasePacked
)

func (caseStatus CaseStatus) String() string {
	return []string{"pending", "building", "packed"}[caseStatus]
}
