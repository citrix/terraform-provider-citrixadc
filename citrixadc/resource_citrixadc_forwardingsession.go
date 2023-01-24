package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcForwardingsession() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createForwardingsessionFunc,
		Read:          readForwardingsessionFunc,
		Update:        updateForwardingsessionFunc,
		Delete:        deleteForwardingsessionFunc,
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
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"acl6name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aclname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connfailover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"processlocal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sourceroutecache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createForwardingsessionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createForwardingsessionFunc")
	client := meta.(*NetScalerNitroClient).client
	forwardingsessionName := d.Get("name").(string)
	
	forwardingsession := network.Forwardingsession{
		Acl6name:         d.Get("acl6name").(string),
		Aclname:          d.Get("aclname").(string),
		Connfailover:     d.Get("connfailover").(string),
		Name:             d.Get("name").(string),
		Netmask:          d.Get("netmask").(string),
		Network:          d.Get("network").(string),
		Processlocal:     d.Get("processlocal").(string),
		Sourceroutecache: d.Get("sourceroutecache").(string),
		Td:               d.Get("td").(int),
	}

	_, err := client.AddResource(service.Forwardingsession.Type(), forwardingsessionName, &forwardingsession)
	if err != nil {
		return err
	}

	d.SetId(forwardingsessionName)

	err = readForwardingsessionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this forwardingsession but we can't read it ?? %s", forwardingsessionName)
		return nil
	}
	return nil
}

func readForwardingsessionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readForwardingsessionFunc")
	client := meta.(*NetScalerNitroClient).client
	forwardingsessionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading forwardingsession state %s", forwardingsessionName)
	data, err := client.FindResource(service.Forwardingsession.Type(), forwardingsessionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing forwardingsession state %s", forwardingsessionName)
		d.SetId("")
		return nil
	}
	d.Set("acl6name", data["acl6name"])
	d.Set("aclname", data["aclname"])
	d.Set("connfailover", data["connfailover"])
	d.Set("name", data["name"])
	d.Set("netmask", data["netmask"])
	d.Set("network", data["network"])
	d.Set("processlocal", data["processlocal"])
	d.Set("sourceroutecache", data["sourceroutecache"])
	d.Set("td", data["td"])

	return nil

}

func updateForwardingsessionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateForwardingsessionFunc")
	client := meta.(*NetScalerNitroClient).client
	forwardingsessionName := d.Get("name").(string)

	forwardingsession := network.Forwardingsession{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("acl6name") {
		log.Printf("[DEBUG]  citrixadc-provider: Acl6name has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Acl6name = d.Get("acl6name").(string)
		hasChange = true
	}
	if d.HasChange("aclname") {
		log.Printf("[DEBUG]  citrixadc-provider: Aclname has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Aclname = d.Get("aclname").(string)
		hasChange = true
	}
	if d.HasChange("connfailover") {
		log.Printf("[DEBUG]  citrixadc-provider: Connfailover has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Connfailover = d.Get("connfailover").(string)
		hasChange = true
	}
	if d.HasChange("processlocal") {
		log.Printf("[DEBUG]  citrixadc-provider: Processlocal has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Processlocal = d.Get("processlocal").(string)
		hasChange = true
	}
	if d.HasChange("sourceroutecache") {
		log.Printf("[DEBUG]  citrixadc-provider: Sourceroutecache has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Sourceroutecache = d.Get("sourceroutecache").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for forwardingsession %s, starting update", forwardingsessionName)
		forwardingsession.Td = d.Get("td").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Forwardingsession.Type(), forwardingsessionName, &forwardingsession)
		if err != nil {
			return fmt.Errorf("Error updating forwardingsession %s", forwardingsessionName)
		}
	}
	return readForwardingsessionFunc(d, meta)
}

func deleteForwardingsessionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteForwardingsessionFunc")
	client := meta.(*NetScalerNitroClient).client
	forwardingsessionName := d.Id()
	err := client.DeleteResource(service.Forwardingsession.Type(), forwardingsessionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
