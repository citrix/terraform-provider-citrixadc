package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcNsservicepath() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsservicepathFunc,
		ReadContext:   readNsservicepathFunc,
		DeleteContext: deleteNsservicepathFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"servicepathname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNsservicepathFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsservicepathFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicepathName := d.Get("servicepathname").(string)
	nsservicepath := ns.Nsservicepath{
		Servicepathname: d.Get("servicepathname").(string),
	}

	_, err := client.AddResource(service.Nsservicepath.Type(), nsservicepathName, &nsservicepath)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsservicepathName)

	return readNsservicepathFunc(ctx, d, meta)
}

func readNsservicepathFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsservicepathFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicepathName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsservicepath state %s", nsservicepathName)
	data, err := client.FindResource(service.Nsservicepath.Type(), nsservicepathName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsservicepath state %s", nsservicepathName)
		d.SetId("")
		return nil
	}
	d.Set("servicepathname", data["servicepathname"])

	return nil

}

func deleteNsservicepathFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsservicepathFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicepathName := d.Id()
	err := client.DeleteResource(service.Nsservicepath.Type(), nsservicepathName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
