package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsvariable() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsvariableFunc,
		Read:          readNsvariableFunc,
		Update:        updateNsvariableFunc,
		Delete:        deleteNsvariableFunc,
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
			"type": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expires": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iffull": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ifnovalue": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ifvaluetoobig": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"init": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsvariableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsvariableFunc")
	client := meta.(*NetScalerNitroClient).client
	nsvariableName := d.Get("name").(string)
	nsvariable := ns.Nsvariable{
		Comment:       d.Get("comment").(string),
		Expires:       d.Get("expires").(int),
		Iffull:        d.Get("iffull").(string),
		Ifnovalue:     d.Get("ifnovalue").(string),
		Ifvaluetoobig: d.Get("ifvaluetoobig").(string),
		Init:          d.Get("init").(string),
		Name:          d.Get("name").(string),
		Scope:         d.Get("scope").(string),
		Type:          d.Get("type").(string),
	}

	_, err := client.AddResource(service.Nsvariable.Type(), nsvariableName, &nsvariable)
	if err != nil {
		return err
	}

	d.SetId(nsvariableName)

	err = readNsvariableFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsvariable but we can't read it ?? %s", nsvariableName)
		return nil
	}
	return nil
}

func readNsvariableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsvariableFunc")
	client := meta.(*NetScalerNitroClient).client
	nsvariableName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsvariable state %s", nsvariableName)
	data, err := client.FindResource(service.Nsvariable.Type(), nsvariableName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsvariable state %s", nsvariableName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("expires", data["expires"])
	d.Set("iffull", data["iffull"])
	d.Set("ifnovalue", data["ifnovalue"])
	d.Set("ifvaluetoobig", data["ifvaluetoobig"])
	d.Set("init", data["init"])
	d.Set("name", data["name"])
	d.Set("scope", data["scope"])
	d.Set("type", data["type"])

	return nil

}

func updateNsvariableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsvariableFunc")
	client := meta.(*NetScalerNitroClient).client
	nsvariableName := d.Get("name").(string)

	nsvariable := ns.Nsvariable{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for nsvariable %s, starting update", nsvariableName)
		nsvariable.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("expires") {
		log.Printf("[DEBUG]  citrixadc-provider: Expires has changed for nsvariable %s, starting update", nsvariableName)
		nsvariable.Expires = d.Get("expires").(int)
		hasChange = true
	}
	if d.HasChange("iffull") {
		log.Printf("[DEBUG]  citrixadc-provider: Iffull has changed for nsvariable %s, starting update", nsvariableName)
		nsvariable.Iffull = d.Get("iffull").(string)
		hasChange = true
	}
	if d.HasChange("ifnovalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Ifnovalue has changed for nsvariable %s, starting update", nsvariableName)
		nsvariable.Ifnovalue = d.Get("ifnovalue").(string)
		hasChange = true
	}
	if d.HasChange("ifvaluetoobig") {
		log.Printf("[DEBUG]  citrixadc-provider: Ifvaluetoobig has changed for nsvariable %s, starting update", nsvariableName)
		nsvariable.Ifvaluetoobig = d.Get("ifvaluetoobig").(string)
		hasChange = true
	}
	if d.HasChange("init") {
		log.Printf("[DEBUG]  citrixadc-provider: Init has changed for nsvariable %s, starting update", nsvariableName)
		nsvariable.Init = d.Get("init").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nsvariable.Type(), nsvariableName, &nsvariable)
		if err != nil {
			return fmt.Errorf("Error updating nsvariable %s", nsvariableName)
		}
	}
	return readNsvariableFunc(d, meta)
}

func deleteNsvariableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsvariableFunc")
	client := meta.(*NetScalerNitroClient).client
	nsvariableName := d.Id()
	err := client.DeleteResource(service.Nsvariable.Type(), nsvariableName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
