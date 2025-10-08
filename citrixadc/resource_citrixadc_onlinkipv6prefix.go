package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcOnlinkipv6prefix() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createOnlinkipv6prefixFunc,
		ReadContext:   readOnlinkipv6prefixFunc,
		UpdateContext: updateOnlinkipv6prefixFunc,
		DeleteContext: deleteOnlinkipv6prefixFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ipv6prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"autonomusprefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"decrementprefixlifetimes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"depricateprefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"onlinkprefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefixpreferredlifetime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"prefixvalidelifetime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createOnlinkipv6prefixFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createOnlinkipv6prefixFunc")
	client := meta.(*NetScalerNitroClient).client
	onlinkipv6prefixName := d.Get("ipv6prefix").(string)

	onlinkipv6prefix := network.Onlinkipv6prefix{
		Autonomusprefix:          d.Get("autonomusprefix").(string),
		Decrementprefixlifetimes: d.Get("decrementprefixlifetimes").(string),
		Depricateprefix:          d.Get("depricateprefix").(string),
		Ipv6prefix:               d.Get("ipv6prefix").(string),
		Onlinkprefix:             d.Get("onlinkprefix").(string),
		Prefixpreferredlifetime:  d.Get("prefixpreferredlifetime").(int),
		Prefixvalidelifetime:     d.Get("prefixvalidelifetime").(int),
	}

	_, err := client.AddResource(service.Onlinkipv6prefix.Type(), onlinkipv6prefixName, &onlinkipv6prefix)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(onlinkipv6prefixName)

	return readOnlinkipv6prefixFunc(ctx, d, meta)
}

func readOnlinkipv6prefixFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readOnlinkipv6prefixFunc")
	client := meta.(*NetScalerNitroClient).client
	onlinkipv6prefixName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading onlinkipv6prefix state %s", onlinkipv6prefixName)
	data, err := client.FindResource(service.Onlinkipv6prefix.Type(), onlinkipv6prefixName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing onlinkipv6prefix state %s", onlinkipv6prefixName)
		d.SetId("")
		return nil
	}
	d.Set("ipv6prefix", data["ipv6prefix"])
	d.Set("autonomusprefix", data["autonomusprefix"])
	d.Set("decrementprefixlifetimes", data["decrementprefixlifetimes"])
	d.Set("depricateprefix", data["depricateprefix"])
	d.Set("ipv6prefix", data["ipv6prefix"])
	d.Set("onlinkprefix", data["onlinkprefix"])
	setToInt("prefixpreferredlifetime", d, data["prefixpreferredlifetime"])
	setToInt("prefixvalidelifetime", d, data["prefixvalidelifetime"])

	return nil

}

func updateOnlinkipv6prefixFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateOnlinkipv6prefixFunc")
	client := meta.(*NetScalerNitroClient).client
	onlinkipv6prefixName := d.Get("ipv6prefix").(string)

	onlinkipv6prefix := network.Onlinkipv6prefix{
		Ipv6prefix: d.Get("ipv6prefix").(string),
	}
	hasChange := false
	if d.HasChange("autonomusprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Autonomusprefix has changed for onlinkipv6prefix %s, starting update", onlinkipv6prefixName)
		onlinkipv6prefix.Autonomusprefix = d.Get("autonomusprefix").(string)
		hasChange = true
	}
	if d.HasChange("decrementprefixlifetimes") {
		log.Printf("[DEBUG]  citrixadc-provider: Decrementprefixlifetimes has changed for onlinkipv6prefix %s, starting update", onlinkipv6prefixName)
		onlinkipv6prefix.Decrementprefixlifetimes = d.Get("decrementprefixlifetimes").(string)
		hasChange = true
	}
	if d.HasChange("depricateprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Depricateprefix has changed for onlinkipv6prefix %s, starting update", onlinkipv6prefixName)
		onlinkipv6prefix.Depricateprefix = d.Get("depricateprefix").(string)
		hasChange = true
	}
	if d.HasChange("onlinkprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Onlinkprefix has changed for onlinkipv6prefix %s, starting update", onlinkipv6prefixName)
		onlinkipv6prefix.Onlinkprefix = d.Get("onlinkprefix").(string)
		hasChange = true
	}
	if d.HasChange("prefixpreferredlifetime") {
		log.Printf("[DEBUG]  citrixadc-provider: Prefixpreferredlifetime has changed for onlinkipv6prefix %s, starting update", onlinkipv6prefixName)
		onlinkipv6prefix.Prefixpreferredlifetime = d.Get("prefixpreferredlifetime").(int)
		hasChange = true
	}
	if d.HasChange("prefixvalidelifetime") {
		log.Printf("[DEBUG]  citrixadc-provider: Prefixvalidelifetime has changed for onlinkipv6prefix %s, starting update", onlinkipv6prefixName)
		onlinkipv6prefix.Prefixvalidelifetime = d.Get("prefixvalidelifetime").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Onlinkipv6prefix.Type(), &onlinkipv6prefix)
		if err != nil {
			return diag.Errorf("Error updating onlinkipv6prefix %s", onlinkipv6prefixName)
		}
	}
	return readOnlinkipv6prefixFunc(ctx, d, meta)
}

func deleteOnlinkipv6prefixFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteOnlinkipv6prefixFunc")
	client := meta.(*NetScalerNitroClient).client
	onlinkipv6prefixName := d.Id()
	err := client.DeleteResource(service.Onlinkipv6prefix.Type(), onlinkipv6prefixName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
