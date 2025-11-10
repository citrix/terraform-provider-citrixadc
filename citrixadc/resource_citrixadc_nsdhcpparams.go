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

func resourceCitrixAdcNsdhcpparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsdhcpparamsFunc,
		ReadContext:   readNsdhcpparamsFunc,
		UpdateContext: updateNsdhcpparamsFunc,
		DeleteContext: deleteNsdhcpparamsFunc,
		Schema: map[string]*schema.Schema{
			"dhcpclient": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saveroute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsdhcpparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsdhcpparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsdhcpparamsName string
	// there is no primary key in nsdhcpparams resource. Hence generate one for terraform state maintenance
	nsdhcpparamsName = resource.PrefixedUniqueId("tf-nsdhcpparams-")
	nsdhcpparams := ns.Nsdhcpparams{
		Dhcpclient: d.Get("dhcpclient").(string),
		Saveroute:  d.Get("saveroute").(string),
	}

	err := client.UpdateUnnamedResource(service.Nsdhcpparams.Type(), &nsdhcpparams)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsdhcpparamsName)

	return readNsdhcpparamsFunc(ctx, d, meta)
}

func readNsdhcpparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsdhcpparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nsdhcpparams state")
	data, err := client.FindResource(service.Nsdhcpparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsdhcpparams state")
		d.SetId("")
		return nil
	}
	d.Set("dhcpclient", data["dhcpclient"])
	d.Set("saveroute", data["saveroute"])

	return nil

}

func updateNsdhcpparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsdhcpparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	nsdhcpparams := ns.Nsdhcpparams{}
	hasChange := false
	if d.HasChange("dhcpclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhcpclient has changed for nsdhcpparams , starting update")
		nsdhcpparams.Dhcpclient = d.Get("dhcpclient").(string)
		hasChange = true
	}
	if d.HasChange("saveroute") {
		log.Printf("[DEBUG]  citrixadc-provider: Saveroute has changed for nsdhcpparams , starting update")
		nsdhcpparams.Saveroute = d.Get("saveroute").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nsdhcpparams.Type(), &nsdhcpparams)
		if err != nil {
			return diag.Errorf("Error updating nsdhcpparams")
		}
	}
	return readNsdhcpparamsFunc(ctx, d, meta)
}

func deleteNsdhcpparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsdhcpparamsFunc")

	d.SetId("")

	return nil
}
