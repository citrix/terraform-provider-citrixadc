package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcLsnstatic() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnstaticFunc,
		ReadContext:   readLsnstaticFunc,
		DeleteContext: deleteLsnstaticFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dsttd": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"natip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"natport": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"nattype": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"network6": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"subscrip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"subscrport": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"transportprotocol": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnstaticFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnstaticFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnstaticName := d.Get("name").(string)
	lsnstatic := lsn.Lsnstatic{
		Destip:            d.Get("destip").(string),
		Dsttd:             d.Get("dsttd").(int),
		Name:              d.Get("name").(string),
		Natip:             d.Get("natip").(string),
		Natport:           d.Get("natport").(int),
		Nattype:           d.Get("nattype").(string),
		Network6:          d.Get("network6").(string),
		Subscrip:          d.Get("subscrip").(string),
		Subscrport:        d.Get("subscrport").(int),
		Td:                d.Get("td").(int),
		Transportprotocol: d.Get("transportprotocol").(string),
	}

	_, err := client.AddResource("lsnstatic", lsnstaticName, &lsnstatic)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsnstaticName)

	return readLsnstaticFunc(ctx, d, meta)
}

func readLsnstaticFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnstaticFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnstaticName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnstatic state %s", lsnstaticName)
	data, err := client.FindResource("lsnstatic", lsnstaticName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnstatic state %s", lsnstaticName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("nattype", data["nattype"])
	d.Set("network6", data["network6"])
	d.Set("subscrip", data["subscrip"])
	setToInt("subscrport", d, data["subscrport"])
	setToInt("td", d, data["td"])
	d.Set("transportprotocol", data["transportprotocol"])

	return nil

}
func deleteLsnstaticFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnstaticFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnstaticName := d.Id()
	err := client.DeleteResource("lsnstatic", lsnstaticName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
