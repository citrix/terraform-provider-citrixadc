package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcNssimpleacl6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNssimpleacl6Func,
		Read:          readNssimpleacl6Func,
		Delete:        deleteNssimpleacl6Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"aclname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"aclaction": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"srcipv6": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"destport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"estsessions": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNssimpleacl6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNssimpleacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nssimpleacl6Name := d.Get("aclname").(string)
	nssimpleacl6 := ns.Nssimpleacl6{
		Aclaction:   d.Get("aclaction").(string),
		Aclname:     d.Get("aclname").(string),
		Destport:    d.Get("destport").(int),
		Estsessions: d.Get("estsessions").(bool),
		Protocol:    d.Get("protocol").(string),
		Srcipv6:     d.Get("srcipv6").(string),
		Td:          d.Get("td").(int),
		Ttl:         d.Get("ttl").(int),
	}

	_, err := client.AddResource(service.Nssimpleacl6.Type(), nssimpleacl6Name, &nssimpleacl6)
	if err != nil {
		return err
	}

	d.SetId(nssimpleacl6Name)

	err = readNssimpleacl6Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nssimpleacl6 but we can't read it ?? %s", nssimpleacl6Name)
		return nil
	}
	return nil
}

func readNssimpleacl6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNssimpleacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nssimpleacl6Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nssimpleacl6 state %s", nssimpleacl6Name)
	data, err := client.FindResource(service.Nssimpleacl6.Type(), nssimpleacl6Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nssimpleacl6 state %s", nssimpleacl6Name)
		d.SetId("")
		return nil
	}
	d.Set("aclname", data["aclname"])
	d.Set("aclaction", data["aclaction"])
	d.Set("aclname", data["aclname"])
	d.Set("destport", data["destport"])
	d.Set("estsessions", data["estsessions"])
	d.Set("protocol", data["protocol"])
	d.Set("srcipv6", data["srcipv6"])
	d.Set("td", data["td"])
	d.Set("ttl", data["ttl"])

	return nil

}

func deleteNssimpleacl6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNssimpleacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nssimpleacl6Name := d.Id()
	err := client.DeleteResource(service.Nssimpleacl6.Type(), nssimpleacl6Name)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
