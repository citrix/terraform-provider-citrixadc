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

func resourceCitrixAdcSystembackupRestore() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSystembackupRestoreFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"filename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"skipbackup": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createSystembackupRestoreFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystembackupRestoreFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName := resource.PrefixedUniqueId(d.Get("filename").(string) + "-")

	systembackup := system.Systembackup{
		Filename:   d.Get("filename").(string),
		Skipbackup: d.Get("skipbackup").(bool),
	}

	err := client.ActOnResource(service.Systembackup.Type(), &systembackup, "restore")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(systembackupName)
	return nil
}
