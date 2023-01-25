package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcVpnintranetapplication() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnintranetapplicationFunc,
		Read:          readVpnintranetapplicationFunc,
		Delete:        deleteVpnintranetapplicationFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"intranetapplication": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"clientapplication": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"destip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"destport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"interception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"iprange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"spoofiip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"srcport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnintranetapplicationFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnintranetapplicationFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnintranetapplicationName := d.Get("intranetapplication").(string)
	vpnintranetapplication := vpn.Vpnintranetapplication{
		Clientapplication:   toStringList(d.Get("clientapplication").([]interface{})),
		Destip:              d.Get("destip").(string),
		Destport:            d.Get("destport").(string),
		Hostname:            d.Get("hostname").(string),
		Interception:        d.Get("interception").(string),
		Intranetapplication: d.Get("intranetapplication").(string),
		Iprange:             d.Get("iprange").(string),
		Netmask:             d.Get("netmask").(string),
		Protocol:            d.Get("protocol").(string),
		Spoofiip:            d.Get("spoofiip").(string),
		Srcip:               d.Get("srcip").(string),
		Srcport:             d.Get("srcport").(int),
	}

	_, err := client.AddResource(service.Vpnintranetapplication.Type(), vpnintranetapplicationName, &vpnintranetapplication)
	if err != nil {
		return err
	}

	d.SetId(vpnintranetapplicationName)

	err = readVpnintranetapplicationFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnintranetapplication but we can't read it ?? %s", vpnintranetapplicationName)
		return nil
	}
	return nil
}

func readVpnintranetapplicationFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnintranetapplicationFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnintranetapplicationName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnintranetapplication state %s", vpnintranetapplicationName)
	data, err := client.FindResource(service.Vpnintranetapplication.Type(), vpnintranetapplicationName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnintranetapplication state %s", vpnintranetapplicationName)
		d.SetId("")
		return nil
	}
	d.Set("intranetapplication", data["intranetapplication"])
	d.Set("clientapplication", data["clientapplication"])
	d.Set("destip", data["destip"])
	d.Set("destport", data["destport"])
	d.Set("hostname", data["hostname"])
	d.Set("interception", data["interception"])
	d.Set("intranetapplication", data["intranetapplication"])
	d.Set("iprange", data["iprange"])
	d.Set("netmask", data["netmask"])
	d.Set("protocol", data["protocol"])
	d.Set("spoofiip", data["spoofiip"])
	d.Set("srcip", data["srcip"])
	d.Set("srcport", data["srcport"])

	return nil

}

func deleteVpnintranetapplicationFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnintranetapplicationFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnintranetapplicationName := d.Id()
	err := client.DeleteResource(service.Vpnintranetapplication.Type(), vpnintranetapplicationName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
