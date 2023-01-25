package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ipsecalg"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcIpsecalgprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIpsecalgprofileFunc,
		Read:          readIpsecalgprofileFunc,
		Update:        updateIpsecalgprofileFunc,
		Delete:        deleteIpsecalgprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"connfailover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"espgatetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"espsessiontimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ikesessiontimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIpsecalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpsecalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecalgprofileName := d.Get("name").(string)
	ipsecalgprofile := ipsecalg.Ipsecalgprofile{
		Connfailover:      d.Get("connfailover").(string),
		Espgatetimeout:    d.Get("espgatetimeout").(int),
		Espsessiontimeout: d.Get("espsessiontimeout").(int),
		Ikesessiontimeout: d.Get("ikesessiontimeout").(int),
		Name:              d.Get("name").(string),
	}

	_, err := client.AddResource("ipsecalgprofile", ipsecalgprofileName, &ipsecalgprofile)
	if err != nil {
		return err
	}

	d.SetId(ipsecalgprofileName)

	err = readIpsecalgprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ipsecalgprofile but we can't read it ?? %s", ipsecalgprofileName)
		return nil
	}
	return nil
}

func readIpsecalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpsecalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecalgprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading ipsecalgprofile state %s", ipsecalgprofileName)
	data, err := client.FindResource("ipsecalgprofile", ipsecalgprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ipsecalgprofile state %s", ipsecalgprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("connfailover", data["connfailover"])
	d.Set("espgatetimeout", data["espgatetimeout"])
	d.Set("espsessiontimeout", data["espsessiontimeout"])
	d.Set("ikesessiontimeout", data["ikesessiontimeout"])

	return nil

}

func updateIpsecalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIpsecalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecalgprofileName := d.Get("name").(string)

	ipsecalgprofile := ipsecalg.Ipsecalgprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("connfailover") {
		log.Printf("[DEBUG]  citrixadc-provider: Connfailover has changed for ipsecalgprofile %s, starting update", ipsecalgprofileName)
		ipsecalgprofile.Connfailover = d.Get("connfailover").(string)
		hasChange = true
	}
	if d.HasChange("espgatetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Espgatetimeout has changed for ipsecalgprofile %s, starting update", ipsecalgprofileName)
		ipsecalgprofile.Espgatetimeout = d.Get("espgatetimeout").(int)
		hasChange = true
	}
	if d.HasChange("espsessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Espsessiontimeout has changed for ipsecalgprofile %s, starting update", ipsecalgprofileName)
		ipsecalgprofile.Espsessiontimeout = d.Get("espsessiontimeout").(int)
		hasChange = true
	}
	if d.HasChange("ikesessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Ikesessiontimeout has changed for ipsecalgprofile %s, starting update", ipsecalgprofileName)
		ipsecalgprofile.Ikesessiontimeout = d.Get("ikesessiontimeout").(int)
		hasChange = true
	}

	if hasChange {
			err := client.UpdateUnnamedResource("ipsecalgprofile", &ipsecalgprofile)
		if err != nil {
			return fmt.Errorf("Error updating ipsecalgprofile %s", ipsecalgprofileName)
		}
	}
	return readIpsecalgprofileFunc(d, meta)
}

func deleteIpsecalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpsecalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecalgprofileName := d.Id()
	err := client.DeleteResource("ipsecalgprofile", ipsecalgprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
