package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcNssimpleacl() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNssimpleaclFunc,
		Read:          readNssimpleaclFunc,
		Delete:        deleteNssimpleaclFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"aclname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"aclaction": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"destport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"estsessions": &schema.Schema{
				Type:     schema.TypeBool,
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
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNssimpleaclFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNssimpleaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nssimpleaclName := d.Get("aclname").(string)
	nssimpleacl := ns.Nssimpleacl{
		Aclaction:   d.Get("aclaction").(string),
		Aclname:     d.Get("aclname").(string),
		Destport:    d.Get("destport").(int),
		Estsessions: d.Get("estsessions").(bool),
		Protocol:    d.Get("protocol").(string),
		Srcip:       d.Get("srcip").(string),
		Td:          d.Get("td").(int),
		Ttl:         d.Get("ttl").(int),
	}

	_, err := client.AddResource(service.Nssimpleacl.Type(), nssimpleaclName, &nssimpleacl)
	if err != nil {
		return err
	}

	d.SetId(nssimpleaclName)

	err = readNssimpleaclFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nssimpleacl but we can't read it ?? %s", nssimpleaclName)
		return nil
	}
	return nil
}

func readNssimpleaclFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNssimpleaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nssimpleaclName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nssimpleacl state %s", nssimpleaclName)
	data, err := client.FindResource(service.Nssimpleacl.Type(), nssimpleaclName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nssimpleacl state %s", nssimpleaclName)
		d.SetId("")
		return nil
	}
	d.Set("aclname", data["aclname"])
	d.Set("aclaction", data["aclaction"])
	d.Set("aclname", data["aclname"])
	d.Set("destport", data["destport"])
	d.Set("estsessions", data["estsessions"])
	d.Set("protocol", data["protocol"])
	d.Set("srcip", data["srcip"])
	d.Set("td", data["td"])
	d.Set("ttl", data["ttl"])

	return nil

}


func deleteNssimpleaclFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNssimpleaclFunc")
	client := meta.(*NetScalerNitroClient).client
	nssimpleaclName := d.Id()
	err := client.DeleteResource(service.Nssimpleacl.Type(), nssimpleaclName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
