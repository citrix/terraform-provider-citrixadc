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

func resourceCitrixAdcNstcpbufparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNstcpbufparamFunc,
		ReadContext:   readNstcpbufparamFunc,
		UpdateContext: updateNstcpbufparamFunc,
		DeleteContext: deleteNstcpbufparamFunc,
		Schema: map[string]*schema.Schema{
			"memlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNstcpbufparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstcpbufparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var nstcpbufparamName string
	// there is no primary key in nstcpbufparam resource. Hence generate one for terraform state maintenance
	nstcpbufparamName = resource.PrefixedUniqueId("tf-nstcpbufparam-")
	nstcpbufparam := ns.Nstcpbufparam{}

	if raw := d.GetRawConfig().GetAttr("memlimit"); !raw.IsNull() {
		nstcpbufparam.Memlimit = intPtr(d.Get("memlimit").(int))
	}
	if raw := d.GetRawConfig().GetAttr("size"); !raw.IsNull() {
		nstcpbufparam.Size = intPtr(d.Get("size").(int))
	}

	err := client.UpdateUnnamedResource(service.Nstcpbufparam.Type(), &nstcpbufparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nstcpbufparamName)

	return readNstcpbufparamFunc(ctx, d, meta)
}

func readNstcpbufparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstcpbufparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nstcpbufparam state")
	data, err := client.FindResource(service.Nstcpbufparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nstcpbufparam state")
		d.SetId("")
		return nil
	}
	setToInt("memlimit", d, data["memlimit"])
	setToInt("size", d, data["size"])

	return nil

}

func updateNstcpbufparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNstcpbufparamFunc")
	client := meta.(*NetScalerNitroClient).client

	nstcpbufparam := ns.Nstcpbufparam{}

	if raw := d.GetRawConfig().GetAttr("memlimit"); !raw.IsNull() {
		nstcpbufparam.Memlimit = intPtr(d.Get("memlimit").(int))
	}
	hasChange := false
	if d.HasChange("memlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Memlimit has changed for nstcpbufparam, starting update")
		hasChange = true
	}
	if d.HasChange("size") {
		log.Printf("[DEBUG]  citrixadc-provider: Size has changed for nstcpbufparam, starting update")
		nstcpbufparam.Size = intPtr(d.Get("size").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nstcpbufparam.Type(), &nstcpbufparam)
		if err != nil {
			return diag.Errorf("Error updating nstcpbufparam")
		}
	}
	return readNstcpbufparamFunc(ctx, d, meta)
}

func deleteNstcpbufparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstcpbufparamFunc")

	d.SetId("")

	return nil
}
