package citrixadc

import (
	"log"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceCitrixAdcHanode() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCitrixAdcHanodeRead,
		Schema: map[string]*schema.Schema{
			"hanode_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"curflips": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"completedfliptime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"deadinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enaifaces": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"failsafe": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"haprop": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hastatus": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hasync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hellointerval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"inc": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"masterstatetime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"maxflips": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"maxfliptime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"routemonitor": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"routemonitorstate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"syncstatusstrictmode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"syncvlan": &schema.Schema{
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
