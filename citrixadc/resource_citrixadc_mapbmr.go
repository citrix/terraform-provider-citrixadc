package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcMapbmr() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createMapbmrFunc,
		Read:          readMapbmrFunc,
		Delete:        deleteMapbmrFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"ruleipv6prefix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"eabitlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"psidlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"psidoffset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createMapbmrFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createMapbmrFunc")
	client := meta.(*NetScalerNitroClient).client
	mapbmrName := d.Get("name").(string)
	mapbmr := network.Mapbmr{
		Eabitlength:    d.Get("eabitlength").(int),
		Name:           d.Get("name").(string),
		Psidlength:     d.Get("psidlength").(int),
		Psidoffset:     d.Get("psidoffset").(int),
		Ruleipv6prefix: d.Get("ruleipv6prefix").(string),
	}

	_, err := client.AddResource("mapbmr", mapbmrName, &mapbmr)
	if err != nil {
		return err
	}

	d.SetId(mapbmrName)

	err = readMapbmrFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this mapbmr but we can't read it ?? %s", mapbmrName)
		return nil
	}
	return nil
}

func readMapbmrFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readMapbmrFunc")
	client := meta.(*NetScalerNitroClient).client
	mapbmrName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading mapbmr state %s", mapbmrName)
	data, err := client.FindResource("mapbmr", mapbmrName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing mapbmr state %s", mapbmrName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("eabitlength", data["eabitlength"])
	d.Set("name", data["name"])
	d.Set("psidlength", data["psidlength"])
	d.Set("psidoffset", data["psidoffset"])
	d.Set("ruleipv6prefix", data["ruleipv6prefix"])

	return nil

}

func deleteMapbmrFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteMapbmrFunc")
	client := meta.(*NetScalerNitroClient).client
	mapbmrName := d.Id()
	err := client.DeleteResource("mapbmr", mapbmrName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
