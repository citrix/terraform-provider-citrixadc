package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strconv"
)

func resourceCitrixAdcDnsmxrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsmxrecFunc,
		Read:          readDnsmxrecFunc,
		Update:        updateDnsmxrecFunc,
		Delete:        deleteDnsmxrecFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mx": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pref": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
			"ecssubnet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnsmxrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsmxrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsmxrecName := d.Get("domain").(string)
	dnsmxrec := dns.Dnsmxrec{
		Domain:    d.Get("domain").(string),
		Ecssubnet: d.Get("ecssubnet").(string),
		Mx:        d.Get("mx").(string),
		Nodeid:    d.Get("nodeid").(int),
		Pref:      d.Get("pref").(int),
		Ttl:       d.Get("ttl").(int),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource(service.Dnsmxrec.Type(), dnsmxrecName, &dnsmxrec)
	if err != nil {
		return err
	}

	d.SetId(dnsmxrecName)

	err = readDnsmxrecFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnsmxrec but we can't read it ?? %s", dnsmxrecName)
		return nil
	}
	return nil
}

func readDnsmxrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsmxrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsmxrecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsmxrec state %s", dnsmxrecName)
	data, err := client.FindResource(service.Dnsmxrec.Type(), dnsmxrecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsmxrec state %s", dnsmxrecName)
		d.SetId("")
		return nil
	}

	pref1, _ := strconv.Atoi(data["pref"].(string))
	d.Set("domain", data["domain"])
	d.Set("ecssubnet", data["ecssubnet"])
	d.Set("mx", data["mx"])
	d.Set("nodeid", data["nodeid"])
	d.Set("pref", pref1)
	d.Set("ttl", data["ttl"])
	d.Set("type", data["type"])
	log.Printf("set functionality:  %v", data)
	return nil

}

func updateDnsmxrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnsmxrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsmxrecName := d.Get("domain").(string)
	dnsmxrec := dns.Dnsmxrec{
		Domain: dnsmxrecName,
		Mx:     d.Get("mx").(string),
	}
	hasChange := false
	if d.HasChange("ecssubnet") {
		log.Printf("[DEBUG]  citrixadc-provider: Ecssubnet has changed for dnsmxrec %s, starting update", dnsmxrecName)
		dnsmxrec.Ecssubnet = d.Get("ecssubnet").(string)
		hasChange = true
	}
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for dnsmxrec %s, starting update", dnsmxrecName)
		dnsmxrec.Nodeid = d.Get("nodeid").(int)
		hasChange = true
	}
	if d.HasChange("pref") {
		log.Printf("[DEBUG]  citrixadc-provider: Pref has changed for dnsmxrec %s, starting update", dnsmxrecName)
		dnsmxrec.Pref = d.Get("pref").(int)
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for dnsmxrec %s, starting update", dnsmxrecName)
		dnsmxrec.Ttl = d.Get("ttl").(int)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for dnsmxrec %s, starting update", dnsmxrecName)
		dnsmxrec.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnsmxrec.Type(), dnsmxrecName, &dnsmxrec)
		if err != nil {
			return fmt.Errorf("Error updating dnsmxrec %s", dnsmxrecName)
		}
	}
	return readDnsmxrecFunc(d, meta)
}

func deleteDnsmxrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsmxrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsmxrecName := d.Id()
	argsMap := make(map[string]string)
	argsMap["mx"] = url.QueryEscape(d.Get("mx").(string))

	if ecscheck, ok := d.GetOk("ecssubnet"); ok {
		argsMap["ecssubnet"] = url.QueryEscape(ecscheck.(string))
	}

	//argsMap["domain"] = dnsmxrecName
	err := client.DeleteResourceWithArgsMap(service.Dnsmxrec.Type(), dnsmxrecName, argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
