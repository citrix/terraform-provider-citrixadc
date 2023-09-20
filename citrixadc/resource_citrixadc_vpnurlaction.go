package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnurlaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnurlactionFunc,
		Read:          readVpnurlactionFunc,
		Update:        updateVpnurlactionFunc,
		Delete:        deleteVpnurlactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"actualurl": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"linkname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"applicationtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientlessaccess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iconurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlssoprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssotype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vservername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpnurlactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnurlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlactionName := d.Get("name").(string)
	vpnurlaction := vpn.Vpnurlaction{
		Actualurl:        d.Get("actualurl").(string),
		Applicationtype:  d.Get("applicationtype").(string),
		Clientlessaccess: d.Get("clientlessaccess").(string),
		Comment:          d.Get("comment").(string),
		Iconurl:          d.Get("iconurl").(string),
		Linkname:         d.Get("linkname").(string),
		Name:             d.Get("name").(string),
		Newname:          d.Get("newname").(string),
		Samlssoprofile:   d.Get("samlssoprofile").(string),
		Ssotype:          d.Get("ssotype").(string),
		Vservername:      d.Get("vservername").(string),
	}

	_, err := client.AddResource("vpnurlaction", vpnurlactionName, &vpnurlaction)
	if err != nil {
		return err
	}

	d.SetId(vpnurlactionName)

	err = readVpnurlactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnurlaction but we can't read it ?? %s", vpnurlactionName)
		return nil
	}
	return nil
}

func readVpnurlactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnurlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnurlaction state %s", vpnurlactionName)
	data, err := client.FindResource("vpnurlaction", vpnurlactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnurlaction state %s", vpnurlactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("actualurl", data["actualurl"])
	d.Set("applicationtype", data["applicationtype"])
	d.Set("clientlessaccess", data["clientlessaccess"])
	d.Set("comment", data["comment"])
	d.Set("iconurl", data["iconurl"])
	d.Set("linkname", data["linkname"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("samlssoprofile", data["samlssoprofile"])
	d.Set("ssotype", data["ssotype"])
	d.Set("vservername", data["vservername"])

	return nil

}

func updateVpnurlactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnurlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlactionName := d.Get("name").(string)

	vpnurlaction := vpn.Vpnurlaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("actualurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Actualurl has changed for vpnurlaction %s, starting update", vpnurlactionName)
		vpnurlaction.Actualurl = d.Get("actualurl").(string)
		hasChange = true
	}
	if d.HasChange("applicationtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Applicationtype has changed for vpnurlaction %s, starting update", vpnurlactionName)
		vpnurlaction.Applicationtype = d.Get("applicationtype").(string)
		hasChange = true
	}
	if d.HasChange("clientlessaccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientlessaccess has changed for vpnurlaction %s, starting update", vpnurlactionName)
		vpnurlaction.Clientlessaccess = d.Get("clientlessaccess").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for vpnurlaction %s, starting update", vpnurlactionName)
		vpnurlaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("iconurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Iconurl has changed for vpnurlaction %s, starting update", vpnurlactionName)
		vpnurlaction.Iconurl = d.Get("iconurl").(string)
		hasChange = true
	}
	if d.HasChange("linkname") {
		log.Printf("[DEBUG]  citrixadc-provider: Linkname has changed for vpnurlaction %s, starting update", vpnurlactionName)
		vpnurlaction.Linkname = d.Get("linkname").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for vpnurlaction %s, starting update", vpnurlactionName)
		vpnurlaction.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("samlssoprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlssoprofile has changed for vpnurlaction %s, starting update", vpnurlactionName)
		vpnurlaction.Samlssoprofile = d.Get("samlssoprofile").(string)
		hasChange = true
	}
	if d.HasChange("ssotype") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssotype has changed for vpnurlaction %s, starting update", vpnurlactionName)
		vpnurlaction.Ssotype = d.Get("ssotype").(string)
		hasChange = true
	}
	if d.HasChange("vservername") {
		log.Printf("[DEBUG]  citrixadc-provider: Vservername has changed for vpnurlaction %s, starting update", vpnurlactionName)
		vpnurlaction.Vservername = d.Get("vservername").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("vpnurlaction", vpnurlactionName, &vpnurlaction)
		if err != nil {
			return fmt.Errorf("Error updating vpnurlaction %s", vpnurlactionName)
		}
	}
	return readVpnurlactionFunc(d, meta)
}

func deleteVpnurlactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnurlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlactionName := d.Id()
	err := client.DeleteResource("vpnurlaction", vpnurlactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
