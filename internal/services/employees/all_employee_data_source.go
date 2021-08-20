package employees

import (
	"../../clients"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/diag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"net/http"
)

func allEmployeeDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: allEmployeeDataSourceRead,
		Schema: map[string]*schema.Schema{
			"employees": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"employee": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"display_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"first_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"preferred_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"gender": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"job_title": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"work_phone": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mobile_phone": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"work_email": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"department": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"location": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"division": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"linkedIn": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"work_phone_extension": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func allEmployeeDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	c := meta.(*clients.Client)

	var diags diag.Diagnostics
	var employeeResponse EmployeeResponse

	requestURL := fmt.Sprintf("%s/employees/diretory", c.HostURL)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req = req.WithContext(ctx)

	// Do request
	body, err := c.DoRequest(req)
	if err != nil {
		return diag.FromErr(err)
	}

	//Unmarshall body and set it
	if err := json.Unmarshal(body, &employeeResponse); err != nil {
		return diag.FromErr(err)
	}

	//TODO: Iterate and set
	if err := d.Set("employee", employeeResponse); err != nil {
		return diag.FromErr(err)
	}

	employeeList := make([]map[string]interface{}, 0)
	for _, e := range employeeResponse.Employees {
		//TODO: Add user verification for errors
		employee := make(map[string]interface{})

		employee["display_name"] = e.DisplayName
		employee["first_name"] = e.FirstName
		employee["last_name"] = e.LastName
		employee["preferred_name"] = e.PreferredName
		employee["gender"] = e.Gender
		employee["job_title"] = e.JobTitle
		employee["work_phone"] = e.WorkPhone
		employee["mobile_phone"] = e.MobilePhone
		employee["work_email"] = e.WorkEmail
		employee["department"] = e.Department
		employee["location"] = e.Location
		employee["division"] = e.Division
		employee["linkedIn"] = e.LinkedIn
		employee["work_phone_extension"] = e.WorkPhoneExtension

		employeeList = append(employeeList, employee)
	}

	//Set tf returns
	if err := d.Set("employees", employeeList); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
