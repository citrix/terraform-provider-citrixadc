package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/subscriber"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSubscribergxinterface() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSubscribergxinterfaceFunc,
		Read:          readSubscribergxinterfaceFunc,
		Update:        updateSubscribergxinterfaceFunc,
		Delete:        deleteSubscribergxinterfaceFunc,
		Schema: map[string]*schema.Schema{
			"cerrequesttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"healthcheckttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"holdonsubscriberabsence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idlettl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"negativettl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"negativettllimitedsuccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pcrfrealm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"purgesdbongxfailure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requestretryattempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"requesttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"revalidationtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicepathavp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},

			},
			"servicepathvendorid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSubscribergxinterfaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSubscribergxinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	subscribergxinterfaceName := resource.PrefixedUniqueId("tf-subscribergxinterface-")
	
	subscribergxinterface := subscriber.Subscribergxinterface{
		Cerrequesttimeout:         d.Get("cerrequesttimeout").(int),
		Healthcheck:               d.Get("healthcheck").(string),
		Healthcheckttl:            d.Get("healthcheckttl").(int),
		Holdonsubscriberabsence:   d.Get("holdonsubscriberabsence").(string),
		Idlettl:                   d.Get("idlettl").(int),
		Negativettl:               d.Get("negativettl").(int),
		Negativettllimitedsuccess: d.Get("negativettllimitedsuccess").(string),
		Nodeid:                    d.Get("nodeid").(int),
		Pcrfrealm:                 d.Get("pcrfrealm").(string),
		Purgesdbongxfailure:       d.Get("purgesdbongxfailure").(string),
		Requestretryattempts:      d.Get("requestretryattempts").(int),
		Requesttimeout:            d.Get("requesttimeout").(int),
		Revalidationtimeout:       d.Get("revalidationtimeout").(int),
		Service:                   d.Get("service").(string),
		Servicepathavp:            toIntegerList(d.Get("servicepathavp").([]interface{})),
		Servicepathvendorid:       d.Get("servicepathvendorid").(int),
		Vserver:                   d.Get("vserver").(string),
	}

	err := client.UpdateUnnamedResource("subscribergxinterface", &subscribergxinterface)
	if err != nil {
		return err
	}

	d.SetId(subscribergxinterfaceName)

	err = readSubscribergxinterfaceFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this subscribergxinterface but we can't read it ??")
		return nil
	}
	return nil
}

func readSubscribergxinterfaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSubscribergxinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading subscribergxinterface state")
	data, err := client.FindResource("subscribergxinterface", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing subscribergxinterface state")
		d.SetId("")
		return nil
	}
	d.Set("cerrequesttimeout", data["cerrequesttimeout"])
	d.Set("healthcheck", data["healthcheck"])
	d.Set("healthcheckttl", data["healthcheckttl"])
	d.Set("holdonsubscriberabsence", data["holdonsubscriberabsence"])
	d.Set("idlettl", data["idlettl"])
	d.Set("negativettl", data["negativettl"])
	d.Set("negativettllimitedsuccess", data["negativettllimitedsuccess"])
	d.Set("nodeid", data["nodeid"])
	d.Set("pcrfrealm", data["pcrfrealm"])
	d.Set("purgesdbongxfailure", data["purgesdbongxfailure"])
	d.Set("requestretryattempts", data["requestretryattempts"])
	d.Set("requesttimeout", data["requesttimeout"])
	d.Set("revalidationtimeout", data["revalidationtimeout"])
	d.Set("service", data["service"])
	d.Set("servicepathavp", data["servicepathavp"])
	d.Set("servicepathvendorid", data["servicepathvendorid"])
	d.Set("vserver", data["vserver"])

	return nil

}

func updateSubscribergxinterfaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSubscribergxinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client

	subscribergxinterface := subscriber.Subscribergxinterface{}
	hasChange := false
	if d.HasChange("cerrequesttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Cerrequesttimeout has changed for subscribergxinterface, starting update")
		subscribergxinterface.Cerrequesttimeout = d.Get("cerrequesttimeout").(int)
		hasChange = true
	}
	if d.HasChange("healthcheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Healthcheck has changed for subscribergxinterface, starting update")
		subscribergxinterface.Healthcheck = d.Get("healthcheck").(string)
		hasChange = true
	}
	if d.HasChange("healthcheckttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Healthcheckttl has changed for subscribergxinterface, starting update")
		subscribergxinterface.Healthcheckttl = d.Get("healthcheckttl").(int)
		hasChange = true
	}
	if d.HasChange("holdonsubscriberabsence") {
		log.Printf("[DEBUG]  citrixadc-provider: Holdonsubscriberabsence has changed for subscribergxinterface, starting update")
		subscribergxinterface.Holdonsubscriberabsence = d.Get("holdonsubscriberabsence").(string)
		hasChange = true
	}
	if d.HasChange("idlettl") {
		log.Printf("[DEBUG]  citrixadc-provider: Idlettl has changed for subscribergxinterface, starting update")
		subscribergxinterface.Idlettl = d.Get("idlettl").(int)
		hasChange = true
	}
	if d.HasChange("negativettl") {
		log.Printf("[DEBUG]  citrixadc-provider: Negativettl has changed for subscribergxinterface, starting update")
		subscribergxinterface.Negativettl = d.Get("negativettl").(int)
		hasChange = true
	}
	if d.HasChange("negativettllimitedsuccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Negativettllimitedsuccess has changed for subscribergxinterface, starting update")
		subscribergxinterface.Negativettllimitedsuccess = d.Get("negativettllimitedsuccess").(string)
		hasChange = true
	}
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for subscribergxinterface, starting update")
		subscribergxinterface.Nodeid = d.Get("nodeid").(int)
		hasChange = true
	}
	if d.HasChange("pcrfrealm") {
		log.Printf("[DEBUG]  citrixadc-provider: Pcrfrealm has changed for subscribergxinterface, starting update")
		subscribergxinterface.Pcrfrealm = d.Get("pcrfrealm").(string)
		hasChange = true
	}
	if d.HasChange("purgesdbongxfailure") {
		log.Printf("[DEBUG]  citrixadc-provider: Purgesdbongxfailure has changed for subscribergxinterface, starting update")
		subscribergxinterface.Purgesdbongxfailure = d.Get("purgesdbongxfailure").(string)
		hasChange = true
	}
	if d.HasChange("requestretryattempts") {
		log.Printf("[DEBUG]  citrixadc-provider: Requestretryattempts has changed for subscribergxinterface, starting update")
		subscribergxinterface.Requestretryattempts = d.Get("requestretryattempts").(int)
		hasChange = true
	}
	if d.HasChange("requesttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Requesttimeout has changed for subscribergxinterface, starting update")
		subscribergxinterface.Requesttimeout = d.Get("requesttimeout").(int)
		hasChange = true
	}
	if d.HasChange("revalidationtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Revalidationtimeout has changed for subscribergxinterface, starting update")
		subscribergxinterface.Revalidationtimeout = d.Get("revalidationtimeout").(int)
		hasChange = true
	}
	if d.HasChange("service") {
		log.Printf("[DEBUG]  citrixadc-provider: Service has changed for subscribergxinterface, starting update")
		subscribergxinterface.Service = d.Get("service").(string)
		hasChange = true
	}
	if d.HasChange("servicepathavp") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicepathavp has changed for subscribergxinterface, starting update")
		subscribergxinterface.Servicepathavp = toIntegerList(d.Get("servicepathavp").([]interface{}))
		hasChange = true
	}
	if d.HasChange("servicepathvendorid") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicepathvendorid has changed for subscribergxinterface, starting update")
		subscribergxinterface.Servicepathvendorid = d.Get("servicepathvendorid").(int)
		hasChange = true
	}
	if d.HasChange("vserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserver has changed for subscribergxinterface, starting update")
		subscribergxinterface.Vserver = d.Get("vserver").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("subscribergxinterface", &subscribergxinterface)
		if err != nil {
			return fmt.Errorf("Error updating subscribergxinterface")
		}
	}
	return readSubscribergxinterfaceFunc(d, meta)
}

func deleteSubscribergxinterfaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSubscribergxinterfaceFunc")
	//subscribergxinterface does not support DELETE operation
	d.SetId("")

	return nil
}
