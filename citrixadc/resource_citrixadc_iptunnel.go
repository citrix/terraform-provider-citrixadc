package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcIptunnel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIptunnelFunc,
		Read:          readIptunnelFunc,
		Delete:        deleteIptunnelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"grepayload": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ipsecprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"local": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
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
			"remote": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"remotesubnetmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createIptunnelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIptunnelFunc")
	client := meta.(*NetScalerNitroClient).client
	iptunnelName := d.Get("name").(string)

	iptunnel := network.Iptunnel{
		Grepayload:       d.Get("grepayload").(string),
		Ipsecprofilename: d.Get("ipsecprofilename").(string),
		Local:            d.Get("local").(string),
		Name:             d.Get("name").(string),
		Ownergroup:       d.Get("ownergroup").(string),
		Protocol:         d.Get("protocol").(string),
		Remote:           d.Get("remote").(string),
		Remotesubnetmask: d.Get("remotesubnetmask").(string),
		Vlan:             d.Get("vlan").(int),
	}

	_, err := client.AddResource(service.Iptunnel.Type(), iptunnelName, &iptunnel)
	if err != nil {
		return err
	}

	d.SetId(iptunnelName)

	err = readIptunnelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this iptunnel but we can't read it ?? %s", iptunnelName)
		return nil
	}
	return nil
}

func readIptunnelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIptunnelFunc")
	client := meta.(*NetScalerNitroClient).client
	iptunnelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading iptunnel state %s", iptunnelName)
	data, err := client.FindResource(service.Iptunnel.Type(), iptunnelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing iptunnel state %s", iptunnelName)
		d.SetId("")
		return nil
	}
	d.Set("grepayload", data["grepayload"])
	d.Set("ipsecprofilename", data["ipsecprofilename"])
	d.Set("local", data["local"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("protocol", data["protocol"])
	d.Set("remote", data["remote"])
	d.Set("remotesubnetmask", data["remotesubnetmask"])
	d.Set("vlan", data["vlan"])

	return nil

}

func deleteIptunnelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIptunnelFunc")
	client := meta.(*NetScalerNitroClient).client
	iptunnelName := d.Id()
	err := client.DeleteResource(service.Iptunnel.Type(), iptunnelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
