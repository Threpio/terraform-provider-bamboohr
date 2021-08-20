package employees

import "strings"

// EmployeeResponse is the top level response from the API
type EmployeeResponse struct {
	Employees []Employee
}

// EmployeeFields holds a slice of EmployeeField which are fields that can be requested on GetEmployee
type EmployeeFields []EmployeeField

// EmployeeField are fields that can be requested on GetEmployee
type EmployeeField string

// Join concatenates the elements of EmployeeFields to create a single string.  The separator is placed between elements in the resulting string.
func (ef EmployeeFields) Join(sep string) string {
	switch len(ef) {
	case 0:
		return ""
	case 1:
		return string(ef[0])
	}
	n := len(sep) * (len(ef) - 1)
	for i := 0; i < len(ef); i++ {
		n += len(ef[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(string(ef[0]))
	for _, s := range ef[1:] {
		b.WriteString(sep)
		b.WriteString(string(s))
	}
	return b.String()
}

// Fields for GetEmployee
const (
	DisplayName        EmployeeField = "DisplayName"
	FirstName                        = "FirstName"
	LastName                         = "LastName"
	PreferredName                    = "PreferredName"
	Gender                           = "Gender"
	JobTitle                         = "JobTitle"
	WorkPhone                        = "WorkPhone"
	MobilePhone                      = "MobilePhone"
	WorkEmail                        = "WorkEmail"
	Department                       = "Department"
	Location                         = "Location"
	Division                         = "Division"
	LinkedIn                         = "LinkedIn"
	WorkPhoneExtension               = "WorkPhoneExtension"
	PhotoUploaded                    = "PhotoUploaded"
	PhotoURL                         = "PhotoURL"
	CanUploadPhoto                   = "CanUploadPhoto"
)

// Employee represents a single person
type Employee struct {
	ID                 string
	DisplayName        string
	FirstName          string
	LastName           string
	PreferredName      string
	Gender             string
	JobTitle           string
	WorkPhone          string
	MobilePhone        string
	WorkEmail          string
	Department         string
	Location           string
	Division           string
	LinkedIn           string
	WorkPhoneExtension string
	PhotoUploaded      *bool // to avoid false when it's empty
	PhotoURL           string
	CanUploadPhoto     *int // to avoid 0 when it's empty
}
