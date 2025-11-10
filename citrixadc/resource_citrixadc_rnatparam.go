package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcRnatparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createRnatparamFunc,
		ReadContext:   readRnatparamFunc,
		UpdateContext: updateRnatparamFunc,
		DeleteContext: deleteRnatparamFunc,
		Schema: map[string]*schema.Schema{
			"srcippersistency": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpproxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRnatparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createRnatparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var rnatparamName string

	// there is no primary key in rnatparam resource. Hence generate one for terraform state maintenance
	rnatparamName = resource.PrefixedUniqueId("tf-rnatparam-")
	rnatparam := network.Rnatparam{
		Srcippersistency: d.Get("srcippersistency").(string),
		Tcpproxy:         d.Get("tcpproxy").(string),
	}

	err := client.UpdateUnnamedResource(service.Rnatparam.Type(), &rnatparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(rnatparamName)

	return readRnatparamFunc(ctx, d, meta)
}

func readRnatparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readRnatparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading rnatparam state")
	data, err := client.FindResource(service.Rnatparam.Type(), "")
	log.Println(data)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rnatparam state")
		d.SetId("")
		return nil
	}

	d.Set("srcippersistency", data["srcippersistency"])
	d.Set("tcpproxy", data["tcpproxy"])

	return nil

}

func updateRnatparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRnatparamFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatparam := network.Rnatparam{}
	hasChange := false
	if d.HasChange("srcippersistency") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcippersistency has changed for rnatparam, starting update")
		rnatparam.Srcippersistency = d.Get("srcippersistency").(string)
		hasChange = true
	}
	if d.HasChange("tcpproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpproxy has changed for rnatparam, starting update")
		rnatparam.Tcpproxy = d.Get("tcpproxy").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Rnatparam.Type(), &rnatparam)
		if err != nil {
			return diag.Errorf("Error updating rnatparam")
		}
	}
	return readRnatparamFunc(ctx, d, meta)
}

func deleteRnatparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRnatparamFunc")

	d.SetId("")

	return nil
}
