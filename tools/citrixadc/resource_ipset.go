package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/network"
	
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)


func resourceCitrixAdcIpset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIpsetFunc,
		Read:          readIpsetFunc,
		Update:        updateIpsetFunc,
		Delete:        deleteIpsetFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
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

func createIpsetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpsetFunc")
	client := meta.(*NetScalerNitroClient).client
	var ipsetName string
	if v, ok := d.GetOk("name"); ok {
		ipsetName = v.(string)
	} else {
		ipsetName= resource.PrefixedUniqueId("tf-ipset-")
		d.Set("name", ipsetName)
	}
	ipset := network.Ipset{
		Name:           d.Get("name").(string),
		Td:           d.Get("td").(int),
		
	}

	_, err := client.AddResource(netscaler.Ipset.Type(), ipsetName, &ipset)
	if err != nil {
		return err
	}

	d.SetId(ipsetName)
	
	err = readIpsetFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ipset but we can't read it ?? %s", ipsetName)
		return nil
	}
	return nil
}

func readIpsetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpsetFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsetName:= d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading ipset state %s", ipsetName)
	data, err := client.FindResource(netscaler.Ipset.Type(), ipsetName)
	if err != nil {
	log.Printf("[WARN] citrixadc-provider: Clearing ipset state %s", ipsetName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("name", data["name"])
	d.Set("td", data["td"])
	

	return nil

}

func updateIpsetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIpsetFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsetName := d.Get("name").(string)

	ipset := network.Ipset{
		Name : d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for ipset %s, starting update", ipsetName)
		ipset.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for ipset %s, starting update", ipsetName)
		ipset.Td = d.Get("td").(int)
		hasChange = true
	}
	

	if hasChange {
		_, err := client.UpdateResource(netscaler.Ipset.Type(), ipsetName, &ipset)
		if err != nil {
			return fmt.Errorf("Error updating ipset %s", ipsetName)
		}
	}
	return readIpsetFunc(d, meta)
}

func deleteIpsetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpsetFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsetName := d.Id()
	err := client.DeleteResource(netscaler.Ipset.Type(), ipsetName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
