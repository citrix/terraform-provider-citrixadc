package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ntp"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNtpserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNtpserverFunc,
		ReadContext:   readNtpserverFunc,
		UpdateContext: updateNtpserverFunc,
		DeleteContext: deleteNtpserverFunc,
		Schema: map[string]*schema.Schema{
			"serverip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"servername": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"autokey": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxpoll": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minpoll": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"preferredntpserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNtpserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNtpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var ntpserverName string
	if v, ok := d.GetOk("serverip"); ok {
		ntpserverName = v.(string)
	} else if v, ok := d.GetOk("servername"); ok {
		ntpserverName = v.(string)
	}
	ntpserver := ntp.Ntpserver{
		Autokey:            d.Get("autokey").(bool),
		Preferredntpserver: d.Get("preferredntpserver").(string),
		Serverip:           d.Get("serverip").(string),
		Servername:         d.Get("servername").(string),
	}

	if raw := d.GetRawConfig().GetAttr("key"); !raw.IsNull() {
		ntpserver.Key = intPtr(d.Get("key").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxpoll"); !raw.IsNull() {
		ntpserver.Maxpoll = intPtr(d.Get("maxpoll").(int))
	}
	if raw := d.GetRawConfig().GetAttr("minpoll"); !raw.IsNull() {
		ntpserver.Minpoll = intPtr(d.Get("minpoll").(int))
	}

	_, err := client.AddResource(service.Ntpserver.Type(), ntpserverName, &ntpserver)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ntpserverName)

	return readNtpserverFunc(ctx, d, meta)
}

func readNtpserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	setToInt("key", d, data["key"])
	setToInt("maxpoll", d, data["maxpoll"])
	setToInt("minpoll", d, data["minpoll"])
	d.Set("preferredntpserver", data["preferredntpserver"])
	//d.Set("serverip", data["serverip"])
	//d.Set("servername", data["servername"])

	return nil

}

func updateNtpserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNtpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	ntpserverName := d.Id()
	ntpserver := ntp.Ntpserver{}

	if v, ok := d.GetOk("serverip"); ok {
		ntpserver.Serverip = v.(string)
	} else if v, ok := d.GetOk("servername"); ok {
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
		ntpserver.Key = intPtr(d.Get("key").(int))
		hasChange = true
	}
	if d.HasChange("maxpoll") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxpoll has changed for ntpserver %s, starting update", ntpserverName)
		ntpserver.Maxpoll = intPtr(d.Get("maxpoll").(int))
		hasChange = true
	}
	if d.HasChange("minpoll") {
		log.Printf("[DEBUG]  citrixadc-provider: Minpoll has changed for ntpserver %s, starting update", ntpserverName)
		ntpserver.Minpoll = intPtr(d.Get("minpoll").(int))
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
			return diag.Errorf("Error updating ntpserver %s", ntpserverName)
		}
	}
	return readNtpserverFunc(ctx, d, meta)
}

func deleteNtpserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNtpserverFunc")
	client := meta.(*NetScalerNitroClient).client
	ntpserverName := d.Id()
	err := client.DeleteResource(service.Ntpserver.Type(), ntpserverName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
