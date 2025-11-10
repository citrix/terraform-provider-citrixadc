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

func resourceCitrixAdcPolicydataset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPolicydatasetFunc,
		ReadContext:   readPolicydatasetFunc,
		UpdateContext: updatePolicydatasetFunc,
		DeleteContext: deletePolicydatasetFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"patsetfile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dynamic": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicydatasetFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicydatasetFunc")
	client := meta.(*NetScalerNitroClient).client
	var policydatasetName string
	if v, ok := d.GetOk("name"); ok {
		policydatasetName = v.(string)
	} else {
		policydatasetName = resource.PrefixedUniqueId("tf-policydataset-")
		d.Set("name", policydatasetName)
	}
	policydataset := policy.Policydataset{
		Comment:    d.Get("comment").(string),
		Name:       d.Get("name").(string),
		Type:       d.Get("type").(string),
		Dynamic:    d.Get("dynamic").(string),
		Patsetfile: d.Get("patsetfile").(string),
	}

	_, err := client.AddResource(service.Policydataset.Type(), policydatasetName, &policydataset)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policydatasetName)

	return readPolicydatasetFunc(ctx, d, meta)
}

func readPolicydatasetFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicydatasetFunc")
	client := meta.(*NetScalerNitroClient).client
	policydatasetName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading policydataset state %s", policydatasetName)
	data, err := client.FindResource(service.Policydataset.Type(), policydatasetName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policydataset state %s", policydatasetName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("patsetfile", data["patsetfile"])
	d.Set("dynamic", data["dynamic"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("type", data["type"])

	return nil

}

func updatePolicydatasetFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePolicydatasetFunc")
	client := meta.(*NetScalerNitroClient).client
	policydatasetName := d.Id()
	policydataset := policy.Policydataset{
		Name: policydatasetName,
	}
	hasChange := false
	if d.HasChange("dynamic") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamic has changed for policydataset %s, starting update", policydatasetName)
		policydataset.Dynamic = d.Get("dynamic").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Policydataset.Type(), policydatasetName, &policydataset)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return readPolicydatasetFunc(ctx, d, meta)
}

func deletePolicydatasetFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicydatasetFunc")
	client := meta.(*NetScalerNitroClient).client
	policydatasetName := d.Id()
	err := client.DeleteResource(service.Policydataset.Type(), policydatasetName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
