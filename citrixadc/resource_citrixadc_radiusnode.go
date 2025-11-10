package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcRadiusnode() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createRadiusnodeFunc,
		ReadContext:   readRadiusnodeFunc,
		UpdateContext: updateRadiusnodeFunc,
		DeleteContext: deleteRadiusnodeFunc,
		Schema: map[string]*schema.Schema{
			"nodeprefix": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"radkey": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createRadiusnodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createRadiusnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	radiusnodeName := d.Get("nodeprefix").(string)
	radiusnode := basic.Radiusnode{
		Nodeprefix: d.Get("nodeprefix").(string),
		Radkey:     d.Get("radkey").(string),
	}

	_, err := client.AddResource("radiusnode", radiusnodeName, &radiusnode)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(radiusnodeName)

	return readRadiusnodeFunc(ctx, d, meta)
}

func readRadiusnodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readRadiusnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	radiusnodeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading radiusnode state %s", radiusnodeName)
	data, err := client.FindResource("radiusnode", radiusnodeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing radiusnode state %s", radiusnodeName)
		d.SetId("")
		return nil
	}
	d.Set("nodeprefix", data["nodeprefix"])
	// d.Set("radkey", data["radkey"])

	return nil

}

func updateRadiusnodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRadiusnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	radiusnodeName := d.Get("nodeprefix").(string)

	radiusnode := basic.Radiusnode{
		Nodeprefix: d.Get("nodeprefix").(string),
	}
	hasChange := false
	if d.HasChange("radkey") {
		log.Printf("[DEBUG]  citrixadc-provider: Radkey has changed for radiusnode %s, starting update", radiusnodeName)
		radiusnode.Radkey = d.Get("radkey").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("radiusnode", radiusnodeName, &radiusnode)
		if err != nil {
			return diag.Errorf("Error updating radiusnode %s", radiusnodeName)
		}
	}
	return readRadiusnodeFunc(ctx, d, meta)
}

func deleteRadiusnodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRadiusnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	radiusnodeName := d.Id()
	err := client.DeleteResource("radiusnode", radiusnodeName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
