package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRemotefile() *schema.Resource {
	return &schema.Resource{
		Description: "File on remote host.",

		ReadContext: dataSourceRemotefileRead,

		Schema: map[string]*schema.Schema{
			"connection": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Connection to host where files are located.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The username on the target host.",
						},
						"private_key": {
							Type:        schema.TypeString,
							Required:    true,
							Sensitive:   true,
							Description: "The private key used to login to the target host.",
						},
						"host": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The target host.",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The ssh port to the target host.",
						},
					},
				},
			},
			"path": {
				Description: "Path to file on remote host.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceRemotefileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	err := resourceRemotefileRead(ctx, d, meta)

	d.SetId(d.Get("path").(string))

	return err
}
