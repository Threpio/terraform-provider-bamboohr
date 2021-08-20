package employees

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/threpio/terraform-provider-bamboohr/internal/clients"
	"net/http"
	"strconv"
)

func employeeDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: employeeDataSourceRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"employee": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"first_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"preferred_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gender": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"job_title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"work_phone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mobile_phone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"work_email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"department": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"division": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"linkedIn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"work_phone_extension": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func employeeDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	c := meta.(*clients.Client)

	var diags diag.Diagnostics
	var employee Employee

	employeeID := strconv.Itoa(d.Get("id").(int))

	requestURL := fmt.Sprintf("%s/employees/%s", c.HostURL, employeeID)

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
	if err := json.Unmarshal(body, &employee); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("employee", employee); err != nil {
		return diag.FromErr(err)
	}

	//Set tf returns
	d.SetId(employee.ID)

	return diags
}
