package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ipsec"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcIpsecparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIpsecparameterFunc,
		Read:          readIpsecparameterFunc,
		Update:        updateIpsecparameterFunc,
		Delete:        deleteIpsecparameterFunc,
		Schema: map[string]*schema.Schema{
			"encalgo": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hashalgo": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ikeretryinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ikeversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"livenesscheckinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"perfectforwardsecrecy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replaywindowsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retransmissiontime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIpsecparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpsecparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecparameterName := resource.PrefixedUniqueId("tf-ipsecparameter-")
	
	ipsecparameter := ipsec.Ipsecparameter{
		Encalgo:               toStringList(d.Get("encalgo").([]interface{})),
		Hashalgo:              toStringList(d.Get("hashalgo").([]interface{})),
		Ikeretryinterval:      d.Get("ikeretryinterval").(int),
		Ikeversion:            d.Get("ikeversion").(string),
		Lifetime:              d.Get("lifetime").(int),
		Livenesscheckinterval: d.Get("livenesscheckinterval").(int),
		Perfectforwardsecrecy: d.Get("perfectforwardsecrecy").(string),
		Replaywindowsize:      d.Get("replaywindowsize").(int),
		Retransmissiontime:    d.Get("retransmissiontime").(int),
	}

	err := client.UpdateUnnamedResource(service.Ipsecparameter.Type(), &ipsecparameter)
	if err != nil {
		return err
	}

	d.SetId(ipsecparameterName)

	err = readIpsecparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ipsecparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readIpsecparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpsecparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading ipsecparameter state")
	data, err := client.FindResource(service.Ipsecparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ipsecparameter state")
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("encalgo", data["encalgo"])
	d.Set("hashalgo", data["hashalgo"])
	d.Set("ikeretryinterval", data["ikeretryinterval"])
	d.Set("ikeversion", data["ikeversion"])
	d.Set("lifetime", data["lifetime"])
	d.Set("livenesscheckinterval", data["livenesscheckinterval"])
	d.Set("perfectforwardsecrecy", data["perfectforwardsecrecy"])
	d.Set("replaywindowsize", data["replaywindowsize"])
	d.Set("retransmissiontime", data["retransmissiontime"])

	return nil

}

func updateIpsecparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIpsecparameterFunc")
	client := meta.(*NetScalerNitroClient).client


	ipsecparameter := ipsec.Ipsecparameter{}
	hasChange := false
	if d.HasChange("encalgo") {
		log.Printf("[DEBUG]  citrixadc-provider: Encalgo has changed for ipsecparameter, starting update")
		ipsecparameter.Encalgo = toStringList(d.Get("encalgo").([]interface{}))
		hasChange = true
	}
	if d.HasChange("hashalgo") {
		log.Printf("[DEBUG]  citrixadc-provider: Hashalgo has changed for ipsecparameter, starting update")
		ipsecparameter.Hashalgo = toStringList(d.Get("hashalgo").([]interface{}))
		hasChange = true
	}
	if d.HasChange("ikeretryinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Ikeretryinterval has changed for ipsecparameter, starting update")
		ipsecparameter.Ikeretryinterval = d.Get("ikeretryinterval").(int)
		hasChange = true
	}
	if d.HasChange("ikeversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Ikeversion has changed for ipsecparameter, starting update")
		ipsecparameter.Ikeversion = d.Get("ikeversion").(string)
		hasChange = true
	}
	if d.HasChange("lifetime") {
		log.Printf("[DEBUG]  citrixadc-provider: Lifetime has changed for ipsecparameter, starting update")
		ipsecparameter.Lifetime = d.Get("lifetime").(int)
		hasChange = true
	}
	if d.HasChange("livenesscheckinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Livenesscheckinterval has changed for ipsecparameter, starting update")
		ipsecparameter.Livenesscheckinterval = d.Get("livenesscheckinterval").(int)
		hasChange = true
	}
	if d.HasChange("perfectforwardsecrecy") {
		log.Printf("[DEBUG]  citrixadc-provider: Perfectforwardsecrecy has changed for ipsecparameter, starting update")
		ipsecparameter.Perfectforwardsecrecy = d.Get("perfectforwardsecrecy").(string)
		hasChange = true
	}
	if d.HasChange("replaywindowsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Replaywindowsize has changed for ipsecparameter, starting update")
		ipsecparameter.Replaywindowsize = d.Get("replaywindowsize").(int)
		hasChange = true
	}
	if d.HasChange("retransmissiontime") {
		log.Printf("[DEBUG]  citrixadc-provider: Retransmissiontime has changed for ipsecparameter, starting update")
		ipsecparameter.Retransmissiontime = d.Get("retransmissiontime").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Ipsecparameter.Type(), &ipsecparameter)
		if err != nil {
			return fmt.Errorf("Error updating ipsecparameter")
		}
	}
	return readIpsecparameterFunc(d, meta)
}

func deleteIpsecparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpsecparameterFunc")
	//ipsecparameter does not support DELETE operation
	d.SetId("")

	return nil
}
