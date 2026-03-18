package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/subscriber"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSubscribergxinterface() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSubscribergxinterfaceFunc,
		ReadContext:   readSubscribergxinterfaceFunc,
		UpdateContext: updateSubscribergxinterfaceFunc,
		DeleteContext: deleteSubscribergxinterfaceFunc,
		Schema: map[string]*schema.Schema{
			"cerrequesttimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healthcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"healthcheckttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"holdonsubscriberabsence": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idlettl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"negativettl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"negativettllimitedsuccess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pcrfrealm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"purgesdbongxfailure": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requestretryattempts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"requesttimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"revalidationtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"service": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicepathavp": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"servicepathvendorid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSubscribergxinterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSubscribergxinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	subscribergxinterfaceName := resource.PrefixedUniqueId("tf-subscribergxinterface-")

	subscribergxinterface := subscriber.Subscribergxinterface{
		Healthcheck:               d.Get("healthcheck").(string),
		Holdonsubscriberabsence:   d.Get("holdonsubscriberabsence").(string),
		Negativettllimitedsuccess: d.Get("negativettllimitedsuccess").(string),
		Pcrfrealm:                 d.Get("pcrfrealm").(string),
		Purgesdbongxfailure:       d.Get("purgesdbongxfailure").(string),
		Service:                   d.Get("service").(string),
		Servicepathavp:            toIntegerList(d.Get("servicepathavp").([]interface{})),
		Vserver:                   d.Get("vserver").(string),
	}

	if raw := d.GetRawConfig().GetAttr("cerrequesttimeout"); !raw.IsNull() {
		subscribergxinterface.Cerrequesttimeout = intPtr(d.Get("cerrequesttimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("healthcheckttl"); !raw.IsNull() {
		subscribergxinterface.Healthcheckttl = intPtr(d.Get("healthcheckttl").(int))
	}
	if raw := d.GetRawConfig().GetAttr("idlettl"); !raw.IsNull() {
		subscribergxinterface.Idlettl = intPtr(d.Get("idlettl").(int))
	}
	if raw := d.GetRawConfig().GetAttr("negativettl"); !raw.IsNull() {
		subscribergxinterface.Negativettl = intPtr(d.Get("negativettl").(int))
	}
	if raw := d.GetRawConfig().GetAttr("nodeid"); !raw.IsNull() {
		subscribergxinterface.Nodeid = intPtr(d.Get("nodeid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("requestretryattempts"); !raw.IsNull() {
		subscribergxinterface.Requestretryattempts = intPtr(d.Get("requestretryattempts").(int))
	}
	if raw := d.GetRawConfig().GetAttr("requesttimeout"); !raw.IsNull() {
		subscribergxinterface.Requesttimeout = intPtr(d.Get("requesttimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("revalidationtimeout"); !raw.IsNull() {
		subscribergxinterface.Revalidationtimeout = intPtr(d.Get("revalidationtimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("servicepathvendorid"); !raw.IsNull() {
		subscribergxinterface.Servicepathvendorid = intPtr(d.Get("servicepathvendorid").(int))
	}

	err := client.UpdateUnnamedResource("subscribergxinterface", &subscribergxinterface)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(subscribergxinterfaceName)

	return readSubscribergxinterfaceFunc(ctx, d, meta)
}

func readSubscribergxinterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSubscribergxinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading subscribergxinterface state")
	data, err := client.FindResource("subscribergxinterface", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing subscribergxinterface state")
		d.SetId("")
		return nil
	}
	setToInt("cerrequesttimeout", d, data["cerrequesttimeout"])
	d.Set("healthcheck", data["healthcheck"])
	setToInt("healthcheckttl", d, data["healthcheckttl"])
	d.Set("holdonsubscriberabsence", data["holdonsubscriberabsence"])
	setToInt("idlettl", d, data["idlettl"])
	setToInt("negativettl", d, data["negativettl"])
	d.Set("negativettllimitedsuccess", data["negativettllimitedsuccess"])
	setToInt("nodeid", d, data["nodeid"])
	d.Set("pcrfrealm", data["pcrfrealm"])
	d.Set("purgesdbongxfailure", data["purgesdbongxfailure"])
	setToInt("requestretryattempts", d, data["requestretryattempts"])
	setToInt("requesttimeout", d, data["requesttimeout"])
	setToInt("revalidationtimeout", d, data["revalidationtimeout"])
	d.Set("service", data["service"])
	if servicepathavp, ok := data["servicepathavp"]; ok && servicepathavp != nil {
		d.Set("servicepathavp", stringListToIntList(servicepathavp.([]interface{})))
	}
	setToInt("servicepathvendorid", d, data["servicepathvendorid"])
	d.Set("vserver", data["vserver"])

	return nil

}

func updateSubscribergxinterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSubscribergxinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client

	subscribergxinterface := subscriber.Subscribergxinterface{}
	hasChange := false
	if d.HasChange("cerrequesttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Cerrequesttimeout has changed for subscribergxinterface, starting update")
		subscribergxinterface.Cerrequesttimeout = intPtr(d.Get("cerrequesttimeout").(int))
		hasChange = true
	}
	if d.HasChange("healthcheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Healthcheck has changed for subscribergxinterface, starting update")
		subscribergxinterface.Healthcheck = d.Get("healthcheck").(string)
		hasChange = true
	}
	if d.HasChange("healthcheckttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Healthcheckttl has changed for subscribergxinterface, starting update")
		subscribergxinterface.Healthcheckttl = intPtr(d.Get("healthcheckttl").(int))
		hasChange = true
	}
	if d.HasChange("holdonsubscriberabsence") {
		log.Printf("[DEBUG]  citrixadc-provider: Holdonsubscriberabsence has changed for subscribergxinterface, starting update")
		subscribergxinterface.Holdonsubscriberabsence = d.Get("holdonsubscriberabsence").(string)
		hasChange = true
	}
	if d.HasChange("idlettl") {
		log.Printf("[DEBUG]  citrixadc-provider: Idlettl has changed for subscribergxinterface, starting update")
		subscribergxinterface.Idlettl = intPtr(d.Get("idlettl").(int))
		hasChange = true
	}
	if d.HasChange("negativettl") {
		log.Printf("[DEBUG]  citrixadc-provider: Negativettl has changed for subscribergxinterface, starting update")
		subscribergxinterface.Negativettl = intPtr(d.Get("negativettl").(int))
		hasChange = true
	}
	if d.HasChange("negativettllimitedsuccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Negativettllimitedsuccess has changed for subscribergxinterface, starting update")
		subscribergxinterface.Negativettllimitedsuccess = d.Get("negativettllimitedsuccess").(string)
		hasChange = true
	}
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for subscribergxinterface, starting update")
		subscribergxinterface.Nodeid = intPtr(d.Get("nodeid").(int))
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
		subscribergxinterface.Requestretryattempts = intPtr(d.Get("requestretryattempts").(int))
		hasChange = true
	}
	if d.HasChange("requesttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Requesttimeout has changed for subscribergxinterface, starting update")
		subscribergxinterface.Requesttimeout = intPtr(d.Get("requesttimeout").(int))
		hasChange = true
	}
	if d.HasChange("revalidationtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Revalidationtimeout has changed for subscribergxinterface, starting update")
		subscribergxinterface.Revalidationtimeout = intPtr(d.Get("revalidationtimeout").(int))
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
		subscribergxinterface.Servicepathvendorid = intPtr(d.Get("servicepathvendorid").(int))
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
			return diag.Errorf("Error updating subscribergxinterface")
		}
	}
	return readSubscribergxinterfaceFunc(ctx, d, meta)
}

func deleteSubscribergxinterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSubscribergxinterfaceFunc")
	//subscribergxinterface does not support DELETE operation
	d.SetId("")

	return nil
}
