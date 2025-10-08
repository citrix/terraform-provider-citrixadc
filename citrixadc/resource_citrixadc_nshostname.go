package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNshostname() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNshostnameFunc,
		ReadContext:   readNshostnameFunc,
		UpdateContext: updateNshostnameFunc,
		DeleteContext: deleteNshostnameFunc,
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNshostnameFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNshostnameFunc")
	client := meta.(*NetScalerNitroClient).client
	nshostnameName := resource.PrefixedUniqueId("tf-nshostname-")
	nshostname := ns.Nshostname{
		Hostname:  d.Get("hostname").(string),
		Ownernode: d.Get("ownernode").(int),
	}

	err := client.UpdateUnnamedResource(service.Nshostname.Type(), &nshostname)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nshostnameName)

	return readNshostnameFunc(ctx, d, meta)
}

func readNshostnameFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNshostnameFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nshostname state")
	data, err := client.FindResource(service.Nshostname.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nshostname state")
		d.SetId("")
		return nil
	}
	d.Set("hostname", data["hostname"])
	setToInt("ownernode", d, data["ownernode"])

	return nil

}

func updateNshostnameFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNshostnameFunc")
	client := meta.(*NetScalerNitroClient).client

	nshostname := ns.Nshostname{
		Hostname: d.Get("hostname").(string),
	}

	hasChange := false
	if d.HasChange("hostname") {
		log.Printf("[DEBUG]  citrixadc-provider: Hostname has changed for nshostname, starting update")
		nshostname.Hostname = d.Get("hostname").(string)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for nshostname, starting update")
		nshostname.Ownernode = d.Get("ownernode").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nshostname.Type(), &nshostname)
		if err != nil {
			return diag.Errorf("Error updating nshostname")
		}
	}
	return readNshostnameFunc(ctx, d, meta)
}

func deleteNshostnameFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNshostnameFunc")
	//nshostname does not support delete operation
	d.SetId("")

	return nil
}
