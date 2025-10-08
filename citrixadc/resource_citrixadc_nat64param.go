package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNat64param() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNat64paramFunc,
		ReadContext:   readNat64paramFunc,
		UpdateContext: updateNat64paramFunc,
		DeleteContext: deleteNat64paramFunc,
		Schema: map[string]*schema.Schema{
			"nat64fragheader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64ignoretos": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64v6mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nat64zerochecksum": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNat64paramFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNat64paramFunc")
	client := meta.(*NetScalerNitroClient).client
	var nat64paramName string
	// there is no primary key in nat64param resource. Hence generate one for terraform state maintenance
	nat64paramName = resource.PrefixedUniqueId("tf-nat64param-")

	nat64param := network.Nat64param{
		Nat64fragheader:   d.Get("nat64fragheader").(string),
		Nat64ignoretos:    d.Get("nat64ignoretos").(string),
		Nat64v6mtu:        d.Get("nat64v6mtu").(int),
		Nat64zerochecksum: d.Get("nat64zerochecksum").(string),
		Td:                d.Get("td").(int),
	}

	err := client.UpdateUnnamedResource("nat64param", &nat64param)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nat64paramName)

	return readNat64paramFunc(ctx, d, meta)
}

func readNat64paramFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNat64paramFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nat64param state")
	data, err := client.FindResource("nat64param", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nat64param state")
		d.SetId("")
		return nil
	}
	d.Set("nat64fragheader", data["nat64fragheader"])
	d.Set("nat64ignoretos", data["nat64ignoretos"])
	val, _ := strconv.Atoi(data["nat64v6mtu"].(string))
	d.Set("nat64v6mtu", val)
	d.Set("nat64zerochecksum", data["nat64zerochecksum"])
	val, _ = strconv.Atoi(data["td"].(string))
	d.Set("td", val)

	return nil

}

func updateNat64paramFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNat64paramFunc")
	client := meta.(*NetScalerNitroClient).client

	nat64param := network.Nat64param{}
	hasChange := false
	if d.HasChange("nat64fragheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Nat64fragheader has changed for nat64param, starting update")
		nat64param.Nat64fragheader = d.Get("nat64fragheader").(string)
		hasChange = true
	}
	if d.HasChange("nat64ignoretos") {
		log.Printf("[DEBUG]  citrixadc-provider: Nat64ignoretos has changed for nat64param, starting update")
		nat64param.Nat64ignoretos = d.Get("nat64ignoretos").(string)
		hasChange = true
	}
	if d.HasChange("nat64v6mtu") {
		log.Printf("[DEBUG]  citrixadc-provider: Nat64v6mtu has changed for nat64param, starting update")
		nat64param.Nat64v6mtu = d.Get("nat64v6mtu").(int)
		hasChange = true
	}
	if d.HasChange("nat64zerochecksum") {
		log.Printf("[DEBUG]  citrixadc-provider: Nat64zerochecksum has changed for nat64param, starting update")
		nat64param.Nat64zerochecksum = d.Get("nat64zerochecksum").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for nat64param, starting update")
		nat64param.Td = d.Get("td").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("nat64param", &nat64param)
		if err != nil {
			return diag.Errorf("Error updating nat64param")
		}
	}
	return readNat64paramFunc(ctx, d, meta)
}

func deleteNat64paramFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNat64paramFunc")

	d.SetId("")

	return nil
}
