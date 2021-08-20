package provider

import (
	"github.com/threpio/terraform-provider-bamboohr/internal/services/employees"
)

func SupportedServices() []ServiceRegistration {
	return []ServiceRegistration{

		employees.Registration{},
	}
}
