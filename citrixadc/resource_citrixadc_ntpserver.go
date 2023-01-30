package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ntp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNtpserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNtpserverFunc,
		Read:          readNtpserverFunc,
		Update:        updateNtpserverFunc,
		Delete:        deleteNtpserverFunc,
		Schema: map[string]*schema.Schema{
			"serverip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"autokey": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxpoll": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minpoll": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"preferredntpserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNtpserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNtpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var ntpserverName string
	if v, ok := d.GetOk("serverip"); ok {
		ntpserverName = v.(string)
	} else  if v, ok := d.GetOk("servername"); ok {
		ntpserverName = v.(string)
	}
	ntpserver := ntp.Ntpserver{
		Autokey:            d.Get("autokey").(bool),
		Key:                d.Get("key").(int),
		Maxpoll:            d.Get("maxpoll").(int),
		Minpoll:            d.Get("minpoll").(int),
		Preferredntpserver: d.Get("preferredntpserver").(string),
		Serverip:           d.Get("serverip").(string),
		Servername:         d.Get("servername").(string),
	}

	_, err := client.AddResource(service.Ntpserver.Type(), ntpserverName, &ntpserver)
	if err != nil {
		return err
	}

	d.SetId(ntpserverName)

	err = readNtpserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ntpserver but we can't read it ?? %s", ntpserverName)
		return nil
	}
	return nil
}

func readNtpserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNtpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	ntpserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading ntpserver state %s", ntpserverName)
	dataArr, err := client.FindAllResources(service.Ntpserver.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ntpserver state %s", ntpserverName)
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["serverip"] == ntpserverName || v["servername"] == ntpserverName {
			foundIndex = i
			break
		}
	}
	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindAllResources Ntpserver not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing Ntpserver state %s", ntpserverName)
		d.SetId("")
		return nil
	}
	
	data := dataArr[foundIndex]
	d.Set("autokey", data["autokey"])
	d.Set("key", data["key"])
	d.Set("maxpoll", data["maxpoll"])
	d.Set("minpoll", data["minpoll"])
	d.Set("preferredntpserver", data["preferredntpserver"])
	//d.Set("serverip", data["serverip"])
	//d.Set("servername", data["servername"])

	return nil

}

func updateNtpserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNtpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	ntpserverName := d.Id()
	ntpserver := ntp.Ntpserver{}
	
	if v, ok := d.GetOk("serverip"); ok {
		ntpserver.Serverip = v.(string)
	} else  if v, ok := d.GetOk("servername"); ok {
		ntpserver.Servername = v.(string)
	}
	
	hasChange := false
	if d.HasChange("autokey") {
		log.Printf("[DEBUG]  citrixadc-provider: Autokey has changed for ntpserver %s, starting update", ntpserverName)
		ntpserver.Autokey = d.Get("autokey").(bool)
		hasChange = true
	}
	if d.HasChange("key") {
		log.Printf("[DEBUG]  citrixadc-provider: Key has changed for ntpserver %s, starting update", ntpserverName)
		ntpserver.Key = d.Get("key").(int)
		hasChange = true
	}
	if d.HasChange("maxpoll") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxpoll has changed for ntpserver %s, starting update", ntpserverName)
		ntpserver.Maxpoll = d.Get("maxpoll").(int)
		hasChange = true
	}
	if d.HasChange("minpoll") {
		log.Printf("[DEBUG]  citrixadc-provider: Minpoll has changed for ntpserver %s, starting update", ntpserverName)
		ntpserver.Minpoll = d.Get("minpoll").(int)
		hasChange = true
	}
	if d.HasChange("preferredntpserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Preferredntpserver has changed for ntpserver %s, starting update", ntpserverName)
		ntpserver.Preferredntpserver = d.Get("preferredntpserver").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Ntpserver.Type(), &ntpserver)
		if err != nil {
			return fmt.Errorf("Error updating ntpserver %s", ntpserverName)
		}
	}
	return readNtpserverFunc(d, meta)
}

func deleteNtpserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNtpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	ntpserverName := d.Id()
	err := client.DeleteResource(service.Ntpserver.Type(), ntpserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
