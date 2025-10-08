package citrixadc

import (
	"context"
	"log"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCitrixAdcHanode() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCitrixAdcHanodeRead,
		Schema: map[string]*schema.Schema{
			"hanode_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"curflips": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"completedfliptime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"deadinterval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enaifaces": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"failsafe": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"haprop": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hastatus": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hasync": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hellointerval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"inc": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"masterstatetime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"maxflips": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"maxfliptime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"routemonitor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"routemonitorstate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"syncstatusstrictmode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"syncvlan": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceCitrixAdcHanodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	log.Printf("[DEBUG] citrixadc-provider:  In dataSourceCitrixAdcHanodeRead")
	client := meta.(*NetScalerNitroClient).client
	hanodeName := d.Get("hanode_id").(int)

	data, err := client.FindResource(service.Hanode.Type(), strconv.Itoa(hanodeName))
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing hanode state")
		return diags
	}
	d.SetId(data["id"].(string))
	setToInt("hanode_id", d, data["id"])
	d.Set("ipaddress", data["ipaddress"])
	setToInt("curflips", d, data["curflips"])
	setToInt("completedfliptime", d, data["completedfliptime"])
	setToInt("deadinterval", d, data["deadinterval"])
	setToInt("enaifaces", d, data["enaifaces"])
	d.Set("failsafe", data["failsafe"])
	d.Set("haprop", data["haprop"])
	d.Set("hastatus", data["hastatus"])
	d.Set("hasync", data["hasync"])
	setToInt("hellointerval", d, data["hellointerval"])
	d.Set("inc", data["inc"])
	setToInt("masterstatetime", d, data["masterstatetime"])
	setToInt("maxflips", d, data["maxflips"])
	setToInt("maxfliptime", d, data["maxfliptime"])
	d.Set("netmask", data["netmask"])
	d.Set("routemonitor", data["routemonitor"])
	d.Set("routemonitorstate", data["routemonitorstate"])
	d.Set("ssl2", data["ssl2"])
	d.Set("state", data["state"])
	d.Set("syncstatusstrictmode", data["syncstatusstrictmode"])
	setToInt("syncvlan", d, data["syncvlan"])

	return diags
}
