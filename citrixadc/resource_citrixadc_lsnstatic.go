package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcLsnstatic() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnstaticFunc,
		Read:          readLsnstaticFunc,
		Delete:        deleteLsnstaticFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dsttd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"natip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"natport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"nattype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"network6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"subscrip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"subscrport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"transportprotocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnstaticFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(lsnstaticName)

	err = readLsnstaticFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnstatic but we can't read it ?? %s", lsnstaticName)
		return nil
	}
	return nil
}

func readLsnstaticFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("subscrport", data["subscrport"])
	d.Set("td", data["td"])
	d.Set("transportprotocol", data["transportprotocol"])

	return nil

}
func deleteLsnstaticFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnstaticFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnstaticName := d.Id()
	err := client.DeleteResource("lsnstatic", lsnstaticName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}