package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAaagroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaagroupFunc,
		ReadContext:   readAaagroupFunc,
		DeleteContext: deleteAaagroupFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"loggedin": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createAaagroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaagroupFunc")
	client := meta.(*NetScalerNitroClient).client
	aaagroupName := d.Get("groupname").(string)
	aaagroup := aaa.Aaagroup{
		Groupname: d.Get("groupname").(string),
		Weight:    d.Get("weight").(int),
		Loggedin:  d.Get("loggedin").(bool),
	}

	_, err := client.AddResource(service.Aaagroup.Type(), aaagroupName, &aaagroup)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaagroupName)

	return readAaagroupFunc(ctx, d, meta)
}

func readAaagroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaagroupFunc")
	client := meta.(*NetScalerNitroClient).client
	aaagroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading aaagroup state %s", aaagroupName)
	data, err := client.FindResource(service.Aaagroup.Type(), aaagroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup state %s", aaagroupName)
		d.SetId("")
		return nil
	}
	d.Set("groupname", data["groupname"])
	d.Set("loggedin", data["loggedin"])
	setToInt("weight", d, data["weight"])

	return nil

}
func deleteAaagroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaagroupFunc")
	client := meta.(*NetScalerNitroClient).client
	aaagroupName := d.Id()
	err := client.DeleteResource(service.Aaagroup.Type(), aaagroupName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
