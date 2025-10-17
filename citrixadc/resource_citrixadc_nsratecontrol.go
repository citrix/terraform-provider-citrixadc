package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsratecontrol() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsratecontrolFunc,
		ReadContext:   readNsratecontrolFunc,
		UpdateContext: updateNsratecontrolFunc,
		DeleteContext: deleteNsratecontrolFunc,
		Schema: map[string]*schema.Schema{
			"icmpthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcprstthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"udpthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsratecontrolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsratecontrolFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsratecontrolName string
	nsratecontrolName = resource.PrefixedUniqueId("tf-nsratecontrol-")
	nsratecontrol := ns.Nsratecontrol{}

	if raw := d.GetRawConfig().GetAttr("icmpthreshold"); !raw.IsNull() {
		nsratecontrol.Icmpthreshold = intPtr(d.Get("icmpthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tcprstthreshold"); !raw.IsNull() {
		nsratecontrol.Tcprstthreshold = intPtr(d.Get("tcprstthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tcpthreshold"); !raw.IsNull() {
		nsratecontrol.Tcpthreshold = intPtr(d.Get("tcpthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("udpthreshold"); !raw.IsNull() {
		nsratecontrol.Udpthreshold = intPtr(d.Get("udpthreshold").(int))
	}

	err := client.UpdateUnnamedResource(service.Nsratecontrol.Type(), &nsratecontrol)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsratecontrolName)

	return readNsratecontrolFunc(ctx, d, meta)
}

func readNsratecontrolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsratecontrolFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nsratecontrol state")
	data, err := client.FindResource(service.Nsratecontrol.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsratecontrol state")
		d.SetId("")
		return nil
	}
	value, _ := strconv.Atoi(data["icmpthreshold"].(string))
	d.Set("icmpthreshold", value)
	value, _ = strconv.Atoi(data["tcprstthreshold"].(string))
	d.Set("tcprstthreshold", value)
	value, _ = strconv.Atoi(data["tcpthreshold"].(string))
	d.Set("tcpthreshold", value)
	value, _ = strconv.Atoi(data["udpthreshold"].(string))
	d.Set("udpthreshold", value)

	return nil

}

func updateNsratecontrolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsratecontrolFunc")
	client := meta.(*NetScalerNitroClient).client

	nsratecontrol := ns.Nsratecontrol{}

	if raw := d.GetRawConfig().GetAttr("icmpthreshold"); !raw.IsNull() {
		nsratecontrol.Icmpthreshold = intPtr(d.Get("icmpthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tcprstthreshold"); !raw.IsNull() {
		nsratecontrol.Tcprstthreshold = intPtr(d.Get("tcprstthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tcpthreshold"); !raw.IsNull() {
		nsratecontrol.Tcpthreshold = intPtr(d.Get("tcpthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("udpthreshold"); !raw.IsNull() {
		nsratecontrol.Udpthreshold = intPtr(d.Get("udpthreshold").(int))
	}

	err := client.UpdateUnnamedResource(service.Nsratecontrol.Type(), &nsratecontrol)
	if err != nil {
		return diag.Errorf("Error updating nsratecontrol")
	}

	return readNsratecontrolFunc(ctx, d, meta)
}

func deleteNsratecontrolFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsratecontrolFunc")
	// nsratecontrol do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
