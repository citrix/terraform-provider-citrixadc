package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcSslfipskey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslfipskeyFunc,
		Read:          readSslfipskeyFunc,
		Delete:        deleteSslfipskeyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"curve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"exponent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"fipskeyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"inform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"iv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"keytype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"modulus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"wrapkeyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslfipskeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslfipskeyFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslfipskeyName = d.Get("fipskeyname").(string)

	sslfipskey := ssl.Sslfipskey{
		Curve:       d.Get("curve").(string),
		Exponent:    d.Get("exponent").(string),
		Fipskeyname: sslfipskeyName,
		Inform:      d.Get("inform").(string),
		Iv:          d.Get("iv").(string),
		Key:         d.Get("key").(string),
		Keytype:     d.Get("keytype").(string),
		Modulus:     d.Get("modulus").(int),
		Wrapkeyname: d.Get("wrapkeyname").(string),
	}

	err := client.ActOnResource(service.Sslfipskey.Type(), &sslfipskey, "create")
	if err != nil {
		return err
	}

	d.SetId(sslfipskeyName)

	err = readSslfipskeyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslfipskey but we can't read it ?? %s", sslfipskeyName)
		return nil
	}
	return nil
}

func readSslfipskeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslfipskeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslfipskeyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslfipskey state %s", sslfipskeyName)
	data, err := client.FindResource(service.Sslfipskey.Type(), sslfipskeyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslfipskey state %s", sslfipskeyName)
		d.SetId("")
		return nil
	}
	d.Set("fipskeyname", data["fipskeyname"])
	d.Set("curve", data["curve"])
	d.Set("exponent", data["exponent"])
	d.Set("fipskeyname", data["fipskeyname"])
	d.Set("inform", data["inform"])
	d.Set("iv", data["iv"])
	d.Set("key", data["key"])
	d.Set("keytype", data["keytype"])
	d.Set("modulus", data["modulus"])
	d.Set("wrapkeyname", data["wrapkeyname"])

	return nil

}

func deleteSslfipskeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslfipskeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslfipskeyName := d.Id()
	err := client.DeleteResource(service.Sslfipskey.Type(), sslfipskeyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
