package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSystembackup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSystembackupFunc,
		Read:          schema.Noop,
		DeleteContext: deleteSystembackupFunc,
		Schema: map[string]*schema.Schema{
			"filename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSystembackupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystembackupFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName := resource.PrefixedUniqueId(d.Get("filename").(string) + "-")

	systembackup := system.Systembackup{
		Filename: d.Get("filename").(string),
	}

	_, err := client.AddResource(service.Systembackup.Type(), systembackupName, &systembackup)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(systembackupName)
	return nil
}

func deleteSystembackupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystembackupFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName := d.Get("filename").(string)
	err := client.DeleteResource(service.Systembackup.Type(), systembackupName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
