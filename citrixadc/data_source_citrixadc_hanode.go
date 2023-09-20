package citrixadc

import (
	"log"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCitrixAdcHanode() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCitrixAdcHanodeRead,
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

func dataSourceCitrixAdcHanodeRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In dataSourceCitrixAdcHanodeRead")
	client := meta.(*NetScalerNitroClient).client
	hanodeName := d.Get("hanode_id").(int)

	data, err := client.FindResource(service.Hanode.Type(), strconv.Itoa(hanodeName))
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing hanode state")
		return nil
	}
	d.SetId(data["id"].(string))
	d.Set("hanode_id", data["id"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("curflips", data["curflips"])
	d.Set("completedfliptime", data["completedfliptime"])
	d.Set("deadinterval", data["deadinterval"])
	d.Set("enaifaces", data["enaifaces"])
	d.Set("failsafe", data["failsafe"])
	d.Set("haprop", data["haprop"])
	d.Set("hastatus", data["hastatus"])
	d.Set("hasync", data["hasync"])
	d.Set("hellointerval", data["hellointerval"])
	d.Set("inc", data["inc"])
	d.Set("masterstatetime", data["masterstatetime"])
	d.Set("maxflips", data["maxflips"])
	d.Set("maxfliptime", data["maxfliptime"])
	d.Set("netmask", data["netmask"])
	d.Set("routemonitor", data["routemonitor"])
	d.Set("routemonitorstate", data["routemonitorstate"])
	d.Set("ssl2", data["ssl2"])
	d.Set("state", data["state"])
	d.Set("syncstatusstrictmode", data["syncstatusstrictmode"])
	d.Set("syncvlan", data["syncvlan"])

	return nil
}
