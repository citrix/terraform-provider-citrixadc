package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/network"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNetprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNetprofileFunc,
		Read:          readNetprofileFunc,
		Update:        updateNetprofileFunc,
		Delete:        deleteNetprofileFunc,
		Schema: map[string]*schema.Schema{
			"mbf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"overridelsn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxyprotocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxyprotocoltxversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcippersistency": &schema.Schema{
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

func createNetprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNetprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	var netprofileName string
	if v, ok := d.GetOk("name"); ok {
		netprofileName = v.(string)
	} else {
		netprofileName = resource.PrefixedUniqueId("tf-netprofile-")
		d.Set("name", netprofileName)
	}
	netprofile := network.Netprofile{
		Mbf:                    d.Get("mbf").(string),
		Name:                   d.Get("name").(string),
		Overridelsn:            d.Get("overridelsn").(string),
		Proxyprotocol:          d.Get("proxyprotocol").(string),
		Proxyprotocoltxversion: d.Get("proxyprotocoltxversion").(string),
		Srcip:                  d.Get("srcip").(string),
		Srcippersistency:       d.Get("srcippersistency").(string),
		Td:                     d.Get("td").(int),
	}

	_, err := client.AddResource(netscaler.Netprofile.Type(), netprofileName, &netprofile)
	if err != nil {
		return err
	}

	d.SetId(netprofileName)

	err = readNetprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this netprofile but we can't read it ?? %s", netprofileName)
		return nil
	}
	return nil
}

func readNetprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNetprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	netprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading netprofile state %s", netprofileName)
	data, err := client.FindResource(netscaler.Netprofile.Type(), netprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing netprofile state %s", netprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("mbf", data["mbf"])
	d.Set("name", data["name"])
	d.Set("overridelsn", data["overridelsn"])
	d.Set("proxyprotocol", data["proxyprotocol"])
	d.Set("proxyprotocoltxversion", data["proxyprotocoltxversion"])
	d.Set("srcip", data["srcip"])
	d.Set("srcippersistency", data["srcippersistency"])
	d.Set("td", data["td"])

	return nil

}

func updateNetprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNetprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	netprofileName := d.Get("name").(string)

	netprofile := network.Netprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("mbf") {
		log.Printf("[DEBUG]  citrixadc-provider: Mbf has changed for netprofile %s, starting update", netprofileName)
		netprofile.Mbf = d.Get("mbf").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for netprofile %s, starting update", netprofileName)
		netprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("overridelsn") {
		log.Printf("[DEBUG]  citrixadc-provider: Overridelsn has changed for netprofile %s, starting update", netprofileName)
		netprofile.Overridelsn = d.Get("overridelsn").(string)
		hasChange = true
	}
	if d.HasChange("proxyprotocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyprotocol has changed for netprofile %s, starting update", netprofileName)
		netprofile.Proxyprotocol = d.Get("proxyprotocol").(string)
		hasChange = true
	}
	if d.HasChange("proxyprotocoltxversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyprotocoltxversion has changed for netprofile %s, starting update", netprofileName)
		netprofile.Proxyprotocoltxversion = d.Get("proxyprotocoltxversion").(string)
		hasChange = true
	}
	if d.HasChange("srcip") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcip has changed for netprofile %s, starting update", netprofileName)
		netprofile.Srcip = d.Get("srcip").(string)
		hasChange = true
	}
	if d.HasChange("srcippersistency") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcippersistency has changed for netprofile %s, starting update", netprofileName)
		netprofile.Srcippersistency = d.Get("srcippersistency").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for netprofile %s, starting update", netprofileName)
		netprofile.Td = d.Get("td").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Netprofile.Type(), netprofileName, &netprofile)
		if err != nil {
			return fmt.Errorf("Error updating netprofile %s", netprofileName)
		}
	}
	return readNetprofileFunc(d, meta)
}

func deleteNetprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNetprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	netprofileName := d.Id()
	err := client.DeleteResource(netscaler.Netprofile.Type(), netprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
