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

func resourceCitrixAdcSystemCreatebackup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSystembackupCreateFunc,
		ReadContext:   readSystembackupCreateFunc,
		DeleteContext: deleteSystembackupCreateFunc,
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"filename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"includekernel": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"level": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"uselocaltimezone": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createSystembackupCreateFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystembackupFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName := resource.PrefixedUniqueId(d.Get("filename").(string) + "-")

	systembackup := system.Systembackup{
		Filename:         d.Get("filename").(string),
		Uselocaltimezone: d.Get("uselocaltimezone").(bool),
		Level:            d.Get("level").(string),
		Includekernel:    d.Get("includekernel").(string),
		Comment:          d.Get("comment").(string),
	}

	err := client.ActOnResource(service.Systembackup.Type(), &systembackup, "create")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(systembackupName)

	return readSystembackupCreateFunc(ctx, d, meta)
}

func readSystembackupCreateFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystembackupFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading systembackup state %s", systembackupName)
	data, err := client.FindResource(service.Systembackup.Type(), d.Get("filename").(string)+".tgz")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing systembackup state %s", systembackupName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	//d.Set("filename", data["filename"])
	//d.Set("includekernel", data["includekernel"])
	//d.Set("level", data["level"])
	//d.Set("skipbackup", data["skipbackup"])
	d.Set("uselocaltimezone", data["uselocaltimezone"])

	return nil

}

func deleteSystembackupCreateFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystembackupFunc")
	client := meta.(*NetScalerNitroClient).client
	systembackupName := d.Get("filename").(string) + ".tgz"
	err := client.DeleteResource(service.Systembackup.Type(), systembackupName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
