package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcDnszone() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnszoneFunc,
		Read:          readDnszoneFunc,
		Update:        updateDnszoneFunc,
		Delete:        deleteDnszoneFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zonename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"proxymode": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"dnssecoffload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyname": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			
		},
	}
}

func createDnszoneFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnszoneFunc")
	client := meta.(*NetScalerNitroClient).client
	dnszoneName:= d.Get("zonename").(string)
	dnszone := dns.Dnszone{
		Dnssecoffload: d.Get("dnssecoffload").(string),
		Keyname:       toStringList(d.Get("keyname").([]interface{})),
		Nsec:          d.Get("nsec").(string),
		Proxymode:     d.Get("proxymode").(string),
		Type:          d.Get("type").(string),
		Zonename:      d.Get("zonename").(string),
	}

	_, err := client.AddResource(service.Dnszone.Type(), dnszoneName, &dnszone)
	if err != nil {
		return err
	}

	d.SetId(dnszoneName)

	err = readDnszoneFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnszone but we can't read it ?? %s", dnszoneName)
		return nil
	}
	return nil
}

func readDnszoneFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnszoneFunc")
	client := meta.(*NetScalerNitroClient).client
	dnszoneName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnszone state %s", dnszoneName)
	data, err := client.FindResource(service.Dnszone.Type(), dnszoneName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnszone state %s", dnszoneName)
		d.SetId("")
		return nil
	}
	d.Set("zonename", data["zonename"])
	d.Set("dnssecoffload", data["dnssecoffload"])
	d.Set("keyname", data["keyname"])
	d.Set("nsec", data["nsec"])
	d.Set("proxymode", data["proxymode"])
	d.Set("type", data["type"])
	return nil

}

func updateDnszoneFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnszoneFunc")
	client := meta.(*NetScalerNitroClient).client
	dnszoneName := d.Get("zonename").(string)

	dnszone := dns.Dnszone{
		Zonename: dnszoneName,
	}
	hasChange := false
	if d.HasChange("dnssecoffload") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnssecoffload has changed for dnszone %s, starting update", dnszoneName)
		dnszone.Dnssecoffload = d.Get("dnssecoffload").(string)
		hasChange = true
	}
	if d.HasChange("keyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Keyname has changed for dnszone %s, starting update", dnszoneName)
		dnszone.Keyname = toStringList(d.Get("keyname").([]interface{}))
		hasChange = true
	}
	if d.HasChange("nsec") {
		log.Printf("[DEBUG]  citrixadc-provider: Nsec has changed for dnszone %s, starting update", dnszoneName)
		dnszone.Nsec = d.Get("nsec").(string)
		hasChange = true
	}
	if d.HasChange("proxymode") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxymode has changed for dnszone %s, starting update", dnszoneName)
		dnszone.Proxymode = d.Get("proxymode").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for dnszone %s, starting update", dnszoneName)
		dnszone.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnszone.Type(), dnszoneName, &dnszone)
		if err != nil {
			return fmt.Errorf("Error updating dnszone %s", dnszoneName)
		}
	}
	return readDnszoneFunc(d, meta)
}

func deleteDnszoneFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnszoneFunc")
	client := meta.(*NetScalerNitroClient).client
	dnszoneName := d.Id()
	err := client.DeleteResource(service.Dnszone.Type(), dnszoneName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
