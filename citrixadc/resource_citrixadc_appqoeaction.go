package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/citrix/adc-nitro-go/service"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAppqoeaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppqoeactionFunc,
		ReadContext:   readAppqoeactionFunc,
		UpdateContext: updateAppqoeactionFunc,
		DeleteContext: deleteAppqoeactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"altcontentpath": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"altcontentsvcname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customfile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dosaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dostrigexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxconn": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"numretries": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polqdepth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"priqdepth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"respondwith": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retryonreset": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retryontimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppqoeactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppqoeactionFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoeactionName := d.Get("name").(string)

	appqoeaction := make(map[string]interface{})

	if v, ok := d.GetOk("tcpprofile"); ok {
		appqoeaction["tcpprofile"] = v.(string)
	}
	if v, ok := d.GetOk("retryontimeout"); ok {
		appqoeaction["retryontimeout"] = v.(int)
	}
	if v, ok := d.GetOk("retryonreset"); ok {
		appqoeaction["retryonreset"] = v.(string)
	}
	if v, ok := d.GetOk("respondwith"); ok {
		appqoeaction["respondwith"] = v.(string)
	}
	if v, ok := d.GetOkExists("priqdepth"); ok {
		appqoeaction["priqdepth"] = v.(int)
	}
	if v, ok := d.GetOk("priority"); ok {
		appqoeaction["priority"] = v.(string)
	}
	if v, ok := d.GetOkExists("polqdepth"); ok {
		appqoeaction["polqdepth"] = v.(int)
	}
	if v, ok := d.GetOkExists("numretries"); ok {
		val, _ := strconv.Atoi(v.(string))
		appqoeaction["numretries"] = val
	}
	if v, ok := d.GetOk("name"); ok {
		appqoeaction["name"] = v.(string)
	}
	if v, ok := d.GetOk("maxconn"); ok {
		appqoeaction["maxconn"] = v.(int)
	}
	if v, ok := d.GetOk("dostrigexpression"); ok {
		appqoeaction["dostrigexpression"] = v.(string)
	}
	if v, ok := d.GetOk("dosaction"); ok {
		appqoeaction["dosaction"] = v.(string)
	}
	if v, ok := d.GetOk("delay"); ok {
		appqoeaction["delay"] = v.(int)
	}
	if v, ok := d.GetOk("customfile"); ok {
		appqoeaction["customfile"] = v.(string)
	}
	if v, ok := d.GetOk("altcontentsvcname"); ok {
		appqoeaction["altcontentsvcname"] = v.(string)
	}
	if v, ok := d.GetOk("altcontentpath"); ok {
		appqoeaction["altcontentpath"] = v.(string)
	}

	_, err := client.AddResource(service.Appqoeaction.Type(), appqoeactionName, &appqoeaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appqoeactionName)

	return readAppqoeactionFunc(ctx, d, meta)
}

func readAppqoeactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppqoeactionFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoeactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appqoeaction state %s", appqoeactionName)
	data, err := client.FindResource(service.Appqoeaction.Type(), appqoeactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appqoeaction state %s", appqoeactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("altcontentpath", data["altcontentpath"])
	d.Set("altcontentsvcname", data["altcontentsvcname"])
	d.Set("customfile", data["customfile"])
	setToInt("delay", d, data["delay"])
	d.Set("dosaction", data["dosaction"])
	d.Set("dostrigexpression", data["dostrigexpression"])
	setToInt("maxconn", d, data["maxconn"])
	d.Set("name", data["name"])
	d.Set("numretries", data["numretries"])
	setToInt("polqdepth", d, data["polqdepth"])
	setToInt("priority", d, data["priority"])
	setToInt("priqdepth", d, data["priqdepth"])
	d.Set("respondwith", data["respondwith"])
	d.Set("retryonreset", data["retryonreset"])
	setToInt("retryontimeout", d, data["retryontimeout"])
	d.Set("tcpprofile", data["tcpprofile"])

	return nil

}

func updateAppqoeactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppqoeactionFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoeactionName := d.Get("name").(string)

	appqoeaction := appqoe.Appqoeaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("altcontentpath") {
		log.Printf("[DEBUG]  citrixadc-provider: Altcontentpath has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Altcontentpath = d.Get("altcontentpath").(string)
		hasChange = true
	}
	if d.HasChange("altcontentsvcname") {
		log.Printf("[DEBUG]  citrixadc-provider: Altcontentsvcname has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Altcontentsvcname = d.Get("altcontentsvcname").(string)
		hasChange = true
	}
	if d.HasChange("customfile") {
		log.Printf("[DEBUG]  citrixadc-provider: Customfile has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Customfile = d.Get("customfile").(string)
		hasChange = true
	}
	if d.HasChange("delay") {
		log.Printf("[DEBUG]  citrixadc-provider: Delay has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Delay = intPtr(d.Get("delay").(int))
		hasChange = true
	}
	if d.HasChange("dosaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Dosaction has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Dosaction = d.Get("dosaction").(string)
		hasChange = true
	}
	if d.HasChange("dostrigexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Dostrigexpression has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Dostrigexpression = d.Get("dostrigexpression").(string)
		hasChange = true
	}
	if d.HasChange("maxconn") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxconn has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Maxconn = intPtr(d.Get("maxconn").(int))
		hasChange = true
	}
	if d.HasChange("numretries") {
		log.Printf("[DEBUG]  citrixadc-provider: Numretries has changed for appqoeaction %s, starting update", appqoeactionName)
		val, _ := strconv.Atoi(d.Get("numretries").(string))
		appqoeaction.Numretries = intPtr(val)
		hasChange = true
	}
	if d.HasChange("polqdepth") {
		log.Printf("[DEBUG]  citrixadc-provider: Polqdepth has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Polqdepth = intPtr(d.Get("polqdepth").(int))
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Priority = d.Get("priority").(string)
		hasChange = true
	}
	if d.HasChange("priqdepth") {
		log.Printf("[DEBUG]  citrixadc-provider: Priqdepth has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Priqdepth = intPtr(d.Get("priqdepth").(int))
		hasChange = true
	}
	if d.HasChange("respondwith") {
		log.Printf("[DEBUG]  citrixadc-provider: Respondwith has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Respondwith = d.Get("respondwith").(string)
		hasChange = true
	}
	if d.HasChange("retryonreset") {
		log.Printf("[DEBUG]  citrixadc-provider: Retryonreset has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Retryonreset = d.Get("retryonreset").(string)
		hasChange = true
	}
	if d.HasChange("retryontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Retryontimeout has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Retryontimeout = intPtr(d.Get("retryontimeout").(int))
		hasChange = true
	}
	if d.HasChange("tcpprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpprofile has changed for appqoeaction %s, starting update", appqoeactionName)
		appqoeaction.Tcpprofile = d.Get("tcpprofile").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Appqoeaction.Type(), appqoeactionName, &appqoeaction)
		if err != nil {
			return diag.Errorf("Error updating appqoeaction %s", appqoeactionName)
		}
	}
	return readAppqoeactionFunc(ctx, d, meta)
}

func deleteAppqoeactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppqoeactionFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoeactionName := d.Id()
	err := client.DeleteResource(service.Appqoeaction.Type(), appqoeactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
