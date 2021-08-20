package main

import (
	"github.com/threpio/terraform-provider-bamboohr/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return provider.BambooHRProvider()
}
