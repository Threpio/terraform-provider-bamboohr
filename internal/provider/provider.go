package provider

import (
	"github.com/threpio/terraform-provider-bamboohr/internal/clients"
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/diag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/errgo.v2/errors"
)

func BambooHRProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"subdomain": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BAMBOOHR_SUBDOMAIN", nil),
			},
			"apiKey": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_KEY", nil),
			},
		},
		ResourcesMap:         map[string]*schema.Resource{},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

//providerConfigure configures the provider in regards to the terraform provider
//TODO: Check ctx here?
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("API_KEY").(string)
	subdomain := d.Get("subdomain").(string)
	var diags diag.Diagnostics

	if apiKey == "" {
		err := errors.New("API_KEY is not set or empty.")
		return nil, diag.FromErr(err)
	}
	if subdomain == "" {
		err := errors.New("BAMBOOHR_SUBDOMAIN is not set or empty.")
		return nil, diag.FromErr(err)
	}

	return buildClient(subdomain, apiKey), diags
}

//buildClient returns a client object for use with auth.
//TODO: This might be redundant
func buildClient(subdomain, apiKey string) (*clients.Client, diag.Diagnostics) {
	client, err := clients.NewClient(subdomain, apiKey)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, nil
}
