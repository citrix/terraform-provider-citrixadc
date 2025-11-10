package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcSnmpcommunity() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSnmpcommunityFunc,
		ReadContext:   readSnmpcommunityFunc,
		DeleteContext: deleteSnmpcommunityFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"communityname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permissions": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSnmpcommunityFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpcommunityFunc")
	client := meta.(*NetScalerNitroClient).client
	communityname := d.Get("communityname").(string)
	snmpcommunity := snmp.Snmpcommunity{
		Communityname: d.Get("communityname").(string),
		Permissions:   d.Get("permissions").(string),
	}

	_, err := client.AddResource(service.Snmpcommunity.Type(), communityname, &snmpcommunity)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(communityname)

	return readSnmpcommunityFunc(ctx, d, meta)
}

func readSnmpcommunityFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpcommunityFunc")
	client := meta.(*NetScalerNitroClient).client
	communityname := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpcommunity state %s", communityname)
	data, err := client.FindResource(service.Snmpcommunity.Type(), communityname)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpcommunity state %s", communityname)
		d.SetId("")
		return nil
	}
	d.Set("communityname", data["communityname"])
	d.Set("permissions", data["permissions"])

	return nil

}

func deleteSnmpcommunityFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpcommunityFunc")
	client := meta.(*NetScalerNitroClient).client
	communityname := d.Id()
	err := client.DeleteResource(service.Snmpcommunity.Type(), communityname)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
