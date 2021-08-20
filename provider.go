package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/threpio/terraform-provider-bamboohr/internal/provider"
)

func Provider() *schema.Provider {
	return provider.BambooHRProvider()
}
