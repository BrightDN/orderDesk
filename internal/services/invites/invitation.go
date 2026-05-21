package invites

type Invitation struct {
	Type    iType
	Email   string
	Company string
	Token   string
}

type iType string

const (
	CompanyType  iType = "company"
	EmployeeType iType = "employee"
)
