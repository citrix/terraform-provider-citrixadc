package citrixadc

import (
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcDnscnamerec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnscnamerecFunc,
		Read:          readDnscnamerecFunc,
		Delete:        deleteDnscnamerecFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"aliasname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"canonicalname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ecssubnet": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createDnscnamerecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnscnamerecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnscnamerecName := d.Get("aliasname").(string)
	dnscnamerec := dns.Dnscnamerec{
		Aliasname:     dnscnamerecName,
		Canonicalname: d.Get("canonicalname").(string),
		Ecssubnet:     d.Get("ecssubnet").(string),
		Nodeid:        d.Get("nodeid").(int),
		Ttl:           d.Get("ttl").(int),
		Type:          d.Get("type").(string),
	}

	_, err := client.AddResource(service.Dnscnamerec.Type(), dnscnamerecName, &dnscnamerec)
	if err != nil {
		return err
	}

	d.SetId(dnscnamerecName)

	err = readDnscnamerecFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnscnamerec but we can't read it ?? %s", dnscnamerecName)
		return nil
	}
	return nil
}

func readDnscnamerecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnscnamerecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnscnamerecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnscnamerec state %s", dnscnamerecName)
	data, err := client.FindResource(service.Dnscnamerec.Type(), dnscnamerecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnscnamerec state %s", dnscnamerecName)
		d.SetId("")
		return nil
	}
	d.Set("aliasname", data["aliasname"])
	d.Set("canonicalname", data["canonicalname"])
	d.Set("ecssubnet", data["ecssubnet"])
	d.Set("nodeid", data["nodeid"])
	d.Set("ttl", data["ttl"])
	d.Set("type", data["type"])

	return nil

}

func deleteDnscnamerecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnscnamerecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnscnamerecName := d.Id()
	argsMap := make(map[string]string)
	if ecs, ok := d.GetOk("ecssubnet"); ok {
		argsMap["ecssubnet"] = url.QueryEscape(ecs.(string))
	}

	err := client.DeleteResourceWithArgsMap(service.Dnscnamerec.Type(), dnscnamerecName, argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
