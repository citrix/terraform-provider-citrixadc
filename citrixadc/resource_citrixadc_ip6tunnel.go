package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcIp6tunnel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIp6tunnelFunc,
		Read:          readIp6tunnelFunc,
		Delete:        deleteIp6tunnelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"local": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"remote": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createIp6tunnelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIp6tunnelFunc")
	client := meta.(*NetScalerNitroClient).client
	ip6tunnelName := d.Get("name").(string)
	ip6tunnel := network.Ip6tunnel{
		Local:      d.Get("local").(string),
		Name:       d.Get("name").(string),
		Ownergroup: d.Get("ownergroup").(string),
		Remote:     d.Get("remote").(string),
	}

	_, err := client.AddResource(service.Ip6tunnel.Type(), ip6tunnelName, &ip6tunnel)
	if err != nil {
		return err
	}

	d.SetId(ip6tunnelName)

	err = readIp6tunnelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ip6tunnel but we can't read it ?? %s", ip6tunnelName)
		return nil
	}
	return nil
}

func readIp6tunnelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIp6tunnelFunc")
	client := meta.(*NetScalerNitroClient).client
	ip6tunnelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading ip6tunnel state %s", ip6tunnelName)
	findParams := service.FindParams{
		ResourceType: service.Ip6tunnel.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ip6tunnel state %s", ip6tunnelName)
		d.SetId("")
		return nil
	}
	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: ip6tunnel does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, ip6tunnel := range dataArray {
		match := true
		if ip6tunnel["remoteip"] != d.Get("remote").(string) {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams ip6tunnel not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing ip6tunnel state %s", ip6tunnelName)
		d.SetId("")
		return nil
	}
	data := dataArray[foundIndex]
	d.Set("local", data["local"])
	d.Set("name", data["name"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("remote", data["remoteip"])
	return nil

}

func deleteIp6tunnelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIp6tunnelFunc")
	client := meta.(*NetScalerNitroClient).client
	ip6tunnelName := d.Id()
	err := client.DeleteResource(service.Ip6tunnel.Type(), ip6tunnelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
