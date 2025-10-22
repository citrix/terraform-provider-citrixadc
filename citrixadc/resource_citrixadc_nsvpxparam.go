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

func resourceCitrixAdcNsvpxparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsvpxparamFunc,
		ReadContext:   readNsvpxparamFunc,
		DeleteContext: deleteNsvpxparamFunc,
		Schema: map[string]*schema.Schema{
			"cpuyield": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"masterclockcpu1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ownernode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNsvpxparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsvpxparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nsvpxparamName := resource.PrefixedUniqueId("tf-nsvpxparam-")

	nsvpxparam := ns.Nsvpxparam{
		Cpuyield:        d.Get("cpuyield").(string),
		Masterclockcpu1: d.Get("masterclockcpu1").(string),
	}
	if raw := d.GetRawConfig().GetAttr("ownernode"); !raw.IsNull() {
		nsvpxparam.Ownernode = intPtr(d.Get("ownernode").(int))
	}

	err := client.UpdateUnnamedResource("nsvpxparam", &nsvpxparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsvpxparamName)

	return readNsvpxparamFunc(ctx, d, meta)
}

func readNsvpxparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsvpxparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nsvpxparamName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsvpxparam state %s", nsvpxparamName)
	findParams := service.FindParams{
		ResourceType: "nsvpxparam",
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return diag.FromErr(err)
	}

	ownernode, ownernodeOk := d.GetOk("ownernode")

	foundIndex := -1
	if ownernodeOk {
		for index, value := range dataArr {
			if ownernode == value["ownernode"] {
				foundIndex = index
			}
		}
	} else {
		// In standalone VPX there is only one entry for nsvpxparam
		foundIndex = 0
	}

	if foundIndex == -1 {
		// Clear state for resource
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]

	d.Set("cpuyield", data["cpuyield"])
	// d.Set("masterclockcpu1", data["masterclockcpu1"])
	setToInt("ownernode", d, data["ownernode"])

	return nil

}

func deleteNsvpxparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsvpxparamFunc")
	// Just delete the reference
	// Actual configuration cannot be deleted
	d.SetId("")

	return nil
}
