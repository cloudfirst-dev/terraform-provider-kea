package kea

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KEA_URL", os.Getenv("KEA_URL")),
				Description: "KEA API Server",
			},
		},
		ConfigureFunc: ConfigureProvider,
		ResourcesMap: map[string]*schema.Resource{
			"kea_reservation": resourceReservation(),
		},
	}
}

func ConfigureProvider(d *schema.ResourceData) (interface{}, error) {
	return NewClient(d.Get("url").(string)), nil
}
