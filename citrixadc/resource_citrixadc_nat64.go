package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNat64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNat64Func,
		ReadContext:   readNat64Func,
		UpdateContext: updateNat64Func,
		DeleteContext: deleteNat64Func,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"acl6name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"netprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNat64Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNat64Func")
	client := meta.(*NetScalerNitroClient).client
	nat64Name := d.Get("name").(string)
	nat64 := network.Nat64{
		Acl6name:   d.Get("acl6name").(string),
		Name:       d.Get("name").(string),
		Netprofile: d.Get("netprofile").(string),
	}

	_, err := client.AddResource(service.Nat64.Type(), nat64Name, &nat64)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nat64Name)

	return readNat64Func(ctx, d, meta)
}

func readNat64Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNat64Func")
	client := meta.(*NetScalerNitroClient).client
	nat64Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nat64 state %s", nat64Name)
	data, err := client.FindResource(service.Nat64.Type(), nat64Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nat64 state %s", nat64Name)
		d.SetId("")
		return nil
	}
	d.Set("acl6name", data["acl6name"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])

	return nil

}

func updateNat64Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNat64Func")
	client := meta.(*NetScalerNitroClient).client
	nat64Name := d.Get("name").(string)

	nat64 := network.Nat64{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("acl6name") {
		log.Printf("[DEBUG]  citrixadc-provider: Acl6name has changed for nat64 %s, starting update", nat64Name)
		nat64.Acl6name = d.Get("acl6name").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Netprofile has changed for nat64 %s, starting update", nat64Name)
		nat64.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nat64.Type(), nat64Name, &nat64)
		if err != nil {
			return diag.Errorf("Error updating nat64 %s", nat64Name)
		}
	}
	return readNat64Func(ctx, d, meta)
}

func deleteNat64Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNat64Func")
	client := meta.(*NetScalerNitroClient).client
	nat64Name := d.Id()
	err := client.DeleteResource(service.Nat64.Type(), nat64Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
