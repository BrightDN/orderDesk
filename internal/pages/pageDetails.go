package pages

type PageData struct {
	Title string
	Type  pageType
}

type pageType string

const (
	OwnerType    pageType = "owner"
	AdminType    pageType = "admin"
	EmployeeType pageType = "employee"
	BusinessType pageType = "business"
)
