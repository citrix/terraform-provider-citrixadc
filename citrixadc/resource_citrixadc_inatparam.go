package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcInatparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createInatparamFunc,
		ReadContext:   readInatparamFunc,
		UpdateContext: updateInatparamFunc,
		DeleteContext: deleteInatparamFunc,
		Schema: map[string]*schema.Schema{
			"nat46fragheader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46ignoretos": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46v6mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nat46v6prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46zerochecksum": {
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

func createInatparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createInatparamFunc")
	client := meta.(*NetScalerNitroClient).client
	inatparamName := resource.PrefixedUniqueId("tf-inatparam-")

	inatparam := make(map[string]interface{})

	if v, ok := d.GetOk("nat46fragheader"); ok {
		inatparam["nat46fragheader"] = v.(string)
	}
	if v, ok := d.GetOk("nat46ignoretos"); ok {
		inatparam["nat46ignoretos"] = v.(string)
	}
	if v, ok := d.GetOk("nat46v6mtu"); ok {
		inatparam["nat46v6mtu"] = v.(int)
	}
	if v, ok := d.GetOk("nat46v6prefix"); ok {
		inatparam["nat46v6prefix"] = v.(string)
	}
	if v, ok := d.GetOk("nat46zerochecksum"); ok {
		inatparam["nat46zerochecksum"] = v.(string)
	}
	if v, ok := d.GetOk("td"); ok {
		inatparam["td"] = v.(int)
	}

	err := client.UpdateUnnamedResource(service.Inatparam.Type(), &inatparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(inatparamName)

	return readInatparamFunc(ctx, d, meta)
}

func readInatparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readInatparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading inatparam state")
	data, err := client.FindResource(service.Inatparam.Type(), strconv.Itoa(d.Get("td").(int)))
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing inatparam state")
		d.SetId("")
		return nil
	}
	d.Set("nat46fragheader", data["nat46fragheader"])
	d.Set("nat46ignoretos", data["nat46ignoretos"])
	setToInt("nat46v6mtu", d, data["nat46v6mtu"])
	d.Set("nat46v6prefix", data["nat46v6prefix"])
	d.Set("nat46zerochecksum", data["nat46zerochecksum"])
	setToInt("td", d, data["td"])

	return nil

}

func updateInatparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateInatparamFunc")
	client := meta.(*NetScalerNitroClient).client

	inatparam := make(map[string]interface{})

	hasChange := false
	if d.HasChange("nat46fragheader") {
		log.Printf("[DEBUG]  citrixadc-provider: nat46fragheader has changed for inatparam, starting update")
		inatparam["nat46fragheader"] = d.Get("nat46fragheader").(string)
		hasChange = true
	}
	if d.HasChange("nat46ignoretos") {
		log.Printf("[DEBUG]  citrixadc-provider: nat46ignoretos has changed for inatparam, starting update")
		inatparam["nat46ignoretos"] = d.Get("nat46ignoretos").(string)
		hasChange = true
	}
	if d.HasChange("nat46v6mtu") {
		log.Printf("[DEBUG]  citrixadc-provider: nat46v6mtu has changed for inatparam, starting update")
		inatparam["nat46v6mtu"] = d.Get("nat46v6mtu").(int)
		hasChange = true
	}
	if d.HasChange("nat46v6prefix") {
		log.Printf("[DEBUG]  citrixadc-provider: nat46v6prefix has changed for inatparam, starting update")
		inatparam["nat46v6prefix"] = d.Get("nat46v6prefix").(string)
		hasChange = true
	}
	if d.HasChange("nat46zerochecksum") {
		log.Printf("[DEBUG]  citrixadc-provider: nat46zerochecksum has changed for inatparam, starting update")
		inatparam["nat46zerochecksum"] = d.Get("nat46zerochecksum").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for inatparam, starting update")
		inatparam["td"] = d.Get("td").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Inatparam.Type(), &inatparam)
		if err != nil {
			return diag.Errorf("Error updating inatparam")
		}
	}
	return readInatparamFunc(ctx, d, meta)
}

func deleteInatparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteInatparamFunc")
	//inatparam does not support DELETE operation
	d.SetId("")

	return nil
}
