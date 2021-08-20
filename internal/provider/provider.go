package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/threpio/terraform-provider-bamboohr/internal/clients"
	"gopkg.in/errgo.v2/errors"
)

type ServiceRegistration interface {
	// Name is the name of this Service
	Name() string

	// WebsiteCategories returns a list of categories which can be used for the sidebar
	WebsiteCategories() []string

	// SupportedDataSources returns the supported Data Sources supported by this Service
	SupportedDataSources() map[string]*schema.Resource

	// SupportedResources returns the supported Resources supported by this Service
	SupportedResources() map[string]*schema.Resource
}

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

	if apiKey == "" {
		err := errors.New("API_KEY is not set or empty.")
		return nil, diag.FromErr(err)
	}
	if subdomain == "" {
		err := errors.New("BAMBOOHR_SUBDOMAIN is not set or empty.")
		return nil, diag.FromErr(err)
	}

	return buildClient(subdomain, apiKey)
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
