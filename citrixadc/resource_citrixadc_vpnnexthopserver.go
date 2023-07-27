package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcVpnnexthopserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnnexthopserverFunc,
		Read:          readVpnnexthopserverFunc,
		Delete:        deleteVpnnexthopserverFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"nexthopport": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"nexthopfqdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nexthopip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resaddresstype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"secure": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnnexthopserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnnexthopserverFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnnexthopserverName := d.Get("name").(string)
	vpnnexthopserver := vpn.Vpnnexthopserver{
		Name:           d.Get("name").(string),
		Nexthopfqdn:    d.Get("nexthopfqdn").(string),
		Nexthopip:      d.Get("nexthopip").(string),
		Nexthopport:    d.Get("nexthopport").(int),
		Resaddresstype: d.Get("resaddresstype").(string),
		Secure:         d.Get("secure").(string),
	}

	_, err := client.AddResource("vpnnexthopserver", vpnnexthopserverName, &vpnnexthopserver)
	if err != nil {
		return err
	}

	d.SetId(vpnnexthopserverName)

	err = readVpnnexthopserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnnexthopserver but we can't read it ?? %s", vpnnexthopserverName)
		return nil
	}
	return nil
}

func readVpnnexthopserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnnexthopserverFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnnexthopserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnnexthopserver state %s", vpnnexthopserverName)
	data, err := client.FindResource(service.Vpnnexthopserver.Type(), vpnnexthopserverName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnnexthopserver state %s", vpnnexthopserverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("nexthopfqdn", data["nexthopfqdn"])
	d.Set("nexthopip", data["nexthopip"])
	d.Set("nexthopport", data["nexthopport"])
	d.Set("resaddresstype", data["resaddresstype"])
	d.Set("secure", data["secure"])

	return nil

}

func deleteVpnnexthopserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnnexthopserverFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnnexthopserverName := d.Id()
	err := client.DeleteResource(service.Vpnnexthopserver.Type(), vpnnexthopserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
