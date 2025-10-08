package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcPolicymap() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPolicymapFunc,
		ReadContext:   readPolicymapFunc,
		DeleteContext: deletePolicymapFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"mappolicyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sd": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"su": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"td": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tu": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicymapFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicymapFunc")
	client := meta.(*NetScalerNitroClient).client
	var policymapName string
	if v, ok := d.GetOk("mappolicyname"); ok {
		policymapName = v.(string)
	} else {
		policymapName = resource.PrefixedUniqueId("tf-policymap-")
		d.Set("mappolicyname", policymapName)
	}
	policymap := policy.Policymap{
		Mappolicyname: d.Get("mappolicyname").(string),
		Sd:            d.Get("sd").(string),
		Su:            d.Get("su").(string),
		Td:            d.Get("td").(string),
		Tu:            d.Get("tu").(string),
	}

	_, err := client.AddResource(service.Policymap.Type(), policymapName, &policymap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policymapName)

	return readPolicymapFunc(ctx, d, meta)
}

func readPolicymapFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicymapFunc")
	client := meta.(*NetScalerNitroClient).client
	policymapName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading policymap state %s", policymapName)
	data, err := client.FindResource(service.Policymap.Type(), policymapName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policymap state %s", policymapName)
		d.SetId("")
		return nil
	}
	d.Set("mappolicyname", data["mappolicyname"])
	d.Set("mappolicyname", data["mappolicyname"])
	d.Set("sd", data["sd"])
	d.Set("su", data["su"])
	d.Set("td", data["td"])
	d.Set("tu", data["tu"])

	return nil

}

func deletePolicymapFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicymapFunc")
	client := meta.(*NetScalerNitroClient).client
	policymapName := d.Id()
	err := client.DeleteResource(service.Policymap.Type(), policymapName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
