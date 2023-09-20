package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcRnat() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRnatFunc,
		Read:          readRnatFunc,
		Update:        updateRnatFunc,
		Delete:        deleteRnatFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"aclname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"connfailover": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"natip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srcippersistency": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"useproxyport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRnatFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRnatFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatName := d.Get("name").(string)

	rnat := network.Rnat{
		Aclname:          d.Get("aclname").(string),
		Connfailover:     d.Get("connfailover").(string),
		Name:             d.Get("name").(string),
		Netmask:          d.Get("netmask").(string),
		Network:          d.Get("network").(string),
		Newname:          d.Get("newname").(string),
		Ownergroup:       d.Get("ownergroup").(string),
		Redirectport:     d.Get("redirectport").(int),
		Srcippersistency: d.Get("srcippersistency").(string),
		Td:               d.Get("td").(int),
		Useproxyport:     d.Get("useproxyport").(string),
	}

	_, err := client.AddResource(service.Rnat.Type(), rnatName, &rnat)
	if err != nil {
		return err
	}

	d.SetId(rnatName)

	err = readRnatFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rnat but we can't read it ?? %s", rnatName)
		return nil
	}
	return nil
}

func readRnatFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRnatFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rnat state %s", rnatName)
	data, err := client.FindResource(service.Rnat.Type(), rnatName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rnat state %s", rnatName)
		d.SetId("")
		return nil
	}
	d.Set("aclname", data["aclname"])
	d.Set("connfailover", data["connfailover"])
	d.Set("name", data["name"])
	d.Set("natip", data["natip"])
	d.Set("netmask", data["netmask"])
	d.Set("network", data["network"])
	d.Set("newname", data["newname"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("redirectport", data["redirectport"])
	d.Set("srcippersistency", data["srcippersistency"])
	d.Set("td", data["td"])
	d.Set("useproxyport", data["useproxyport"])

	return nil

}

func updateRnatFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRnatFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatName := d.Get("name").(string)

	rnat := network.Rnat{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("connfailover") {
		log.Printf("[DEBUG]  citrixadc-provider: Connfailover has changed for rnat %s, starting update", rnatName)
		rnat.Connfailover = d.Get("connfailover").(string)
		hasChange = true
	}
	if d.HasChange("natip") {
		log.Printf("[DEBUG]  citrixadc-provider: Natip has changed for rnat %s, starting update", rnatName)
		rnat.Natip = d.Get("natip").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for rnat %s, starting update", rnatName)
		rnat.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("ownergroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownergroup has changed for rnat %s, starting update", rnatName)
		rnat.Ownergroup = d.Get("ownergroup").(string)
		hasChange = true
	}
	if d.HasChange("redirectport") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectport has changed for rnat %s, starting update", rnatName)
		rnat.Redirectport = d.Get("redirectport").(int)
		hasChange = true
	}
	if d.HasChange("srcippersistency") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcippersistency has changed for rnat %s, starting update", rnatName)
		rnat.Srcippersistency = d.Get("srcippersistency").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for rnat %s, starting update", rnatName)
		rnat.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("useproxyport") {
		log.Printf("[DEBUG]  citrixadc-provider: Useproxyport has changed for rnat %s, starting update", rnatName)
		rnat.Useproxyport = d.Get("useproxyport").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Rnat.Type(), rnatName, &rnat)
		if err != nil {
			return fmt.Errorf("Error updating rnat %s", rnatName)
		}
	}
	return readRnatFunc(d, meta)
}

func deleteRnatFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRnatFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatName := d.Id()
	err := client.DeleteResource(service.Rnat.Type(), rnatName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
