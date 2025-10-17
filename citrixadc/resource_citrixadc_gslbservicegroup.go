package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcGslbservicegroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createGslbservicegroupFunc,
		ReadContext:   readGslbservicegroupFunc,
		UpdateContext: updateGslbservicegroupFunc,
		DeleteContext: deleteGslbservicegroupFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"servicegroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"appflowlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autoscale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cipheader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clttimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"downstateflush": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dupweight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"graceful": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hashid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healthmonitor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"includemembers": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitornamesvc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"publicip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publicport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitepersistence": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"siteprefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svrtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createGslbservicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbservicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbservicegroupName := d.Get("servicegroupname").(string)

	gslbservicegroup := gslb.Gslbservicegroup{
		Appflowlog:       d.Get("appflowlog").(string),
		Autoscale:        d.Get("autoscale").(string),
		Cip:              d.Get("cip").(string),
		Cipheader:        d.Get("cipheader").(string),
		Comment:          d.Get("comment").(string),
		Downstateflush:   d.Get("downstateflush").(string),
		Graceful:         d.Get("graceful").(string),
		Healthmonitor:    d.Get("healthmonitor").(string),
		Includemembers:   d.Get("includemembers").(bool),
		Monitornamesvc:   d.Get("monitornamesvc").(string),
		Publicip:         d.Get("publicip").(string),
		Servername:       d.Get("servername").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
		Servicetype:      d.Get("servicetype").(string),
		Sitename:         d.Get("sitename").(string),
		Sitepersistence:  d.Get("sitepersistence").(string),
		Siteprefix:       d.Get("siteprefix").(string),
		State:            d.Get("state").(string),
	}

	if raw := d.GetRawConfig().GetAttr("clttimeout"); !raw.IsNull() {
		gslbservicegroup.Clttimeout = intPtr(d.Get("clttimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("delay"); !raw.IsNull() {
		gslbservicegroup.Delay = intPtr(d.Get("delay").(int))
	}
	if raw := d.GetRawConfig().GetAttr("dupweight"); !raw.IsNull() {
		gslbservicegroup.Dupweight = intPtr(d.Get("dupweight").(int))
	}
	if raw := d.GetRawConfig().GetAttr("hashid"); !raw.IsNull() {
		gslbservicegroup.Hashid = intPtr(d.Get("hashid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxbandwidth"); !raw.IsNull() {
		gslbservicegroup.Maxbandwidth = intPtr(d.Get("maxbandwidth").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxclient"); !raw.IsNull() {
		gslbservicegroup.Maxclient = intPtr(d.Get("maxclient").(int))
	}
	if raw := d.GetRawConfig().GetAttr("monthreshold"); !raw.IsNull() {
		gslbservicegroup.Monthreshold = intPtr(d.Get("monthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		gslbservicegroup.Port = intPtr(d.Get("port").(int))
	}
	if raw := d.GetRawConfig().GetAttr("publicport"); !raw.IsNull() {
		gslbservicegroup.Publicport = intPtr(d.Get("publicport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("svrtimeout"); !raw.IsNull() {
		gslbservicegroup.Svrtimeout = intPtr(d.Get("svrtimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("weight"); !raw.IsNull() {
		gslbservicegroup.Weight = intPtr(d.Get("weight").(int))
	}

	_, err := client.AddResource("gslbservicegroup", gslbservicegroupName, &gslbservicegroup)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(gslbservicegroupName)

	return readGslbservicegroupFunc(ctx, d, meta)
}

func readGslbservicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbservicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbservicegroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading gslbservicegroup state %s", gslbservicegroupName)
	data, err := client.FindResource("gslbservicegroup", gslbservicegroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservicegroup state %s", gslbservicegroupName)
		d.SetId("")
		return nil
	}
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("autoscale", data["autoscale"])
	d.Set("cip", data["cip"])
	d.Set("cipheader", data["cipheader"])
	setToInt("clttimeout", d, data["clttimeout"])
	d.Set("comment", data["comment"])
	setToInt("delay", d, data["delay"])
	d.Set("downstateflush", data["downstateflush"])
	setToInt("dupweight", d, data["dupweight"])
	d.Set("graceful", data["graceful"])
	setToInt("hashid", d, data["hashid"])
	d.Set("healthmonitor", data["healthmonitor"])
	d.Set("includemembers", data["includemembers"])
	setToInt("maxbandwidth", d, data["maxbandwidth"])
	setToInt("maxclient", d, data["maxclient"])
	d.Set("monitornamesvc", data["monitornamesvc"])
	setToInt("monthreshold", d, data["monthreshold"])
	setToInt("port", d, data["port"])
	d.Set("publicip", data["publicip"])
	setToInt("publicport", d, data["publicport"])
	d.Set("servername", data["servername"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sitename", data["sitename"])
	d.Set("sitepersistence", data["sitepersistence"])
	d.Set("siteprefix", data["siteprefix"])
	d.Set("state", data["state"])
	setToInt("svrtimeout", d, data["svrtimeout"])
	setToInt("weight", d, data["weight"])

	return nil

}

func updateGslbservicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateGslbservicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbservicegroupName := d.Get("servicegroupname").(string)

	gslbservicegroup := gslb.Gslbservicegroup{
		Servicegroupname: d.Get("servicegroupname").(string),
	}
	hasChange := false
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowlog has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("autoscale") {
		log.Printf("[DEBUG]  citrixadc-provider: Autoscale has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Autoscale = d.Get("autoscale").(string)
		hasChange = true
	}
	if d.HasChange("cip") {
		log.Printf("[DEBUG]  citrixadc-provider: Cip has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Cip = d.Get("cip").(string)
		hasChange = true
	}
	if d.HasChange("cipheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Cipheader has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Cipheader = d.Get("cipheader").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Clttimeout has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Clttimeout = intPtr(d.Get("clttimeout").(int))
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("delay") {
		log.Printf("[DEBUG]  citrixadc-provider: Delay has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Delay = intPtr(d.Get("delay").(int))
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG]  citrixadc-provider: Downstateflush has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("dupweight") {
		log.Printf("[DEBUG]  citrixadc-provider: Dupweight has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Dupweight = intPtr(d.Get("dupweight").(int))
		hasChange = true
	}
	if d.HasChange("graceful") {
		log.Printf("[DEBUG]  citrixadc-provider: Graceful has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Graceful = d.Get("graceful").(string)
		hasChange = true
	}
	if d.HasChange("hashid") {
		log.Printf("[DEBUG]  citrixadc-provider: Hashid has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Hashid = intPtr(d.Get("hashid").(int))
		hasChange = true
	}
	if d.HasChange("healthmonitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Healthmonitor has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Healthmonitor = d.Get("healthmonitor").(string)
		hasChange = true
	}
	if d.HasChange("includemembers") {
		log.Printf("[DEBUG]  citrixadc-provider: Includemembers has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Includemembers = d.Get("includemembers").(bool)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxbandwidth has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Maxbandwidth = intPtr(d.Get("maxbandwidth").(int))
		hasChange = true
	}
	if d.HasChange("maxclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxclient has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Maxclient = intPtr(d.Get("maxclient").(int))
		hasChange = true
	}
	if d.HasChange("monitornamesvc") {
		log.Printf("[DEBUG]  citrixadc-provider: Monitornamesvc has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Monitornamesvc = d.Get("monitornamesvc").(string)
		hasChange = true
	}
	if d.HasChange("monthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Monthreshold has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Monthreshold = intPtr(d.Get("monthreshold").(int))
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Port = intPtr(d.Get("port").(int))
		hasChange = true
	}
	if d.HasChange("publicip") {
		log.Printf("[DEBUG]  citrixadc-provider: Publicip has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Publicip = d.Get("publicip").(string)
		hasChange = true
	}
	if d.HasChange("publicport") {
		log.Printf("[DEBUG]  citrixadc-provider: Publicport has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Publicport = intPtr(d.Get("publicport").(int))
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicetype has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sitename") {
		log.Printf("[DEBUG]  citrixadc-provider: Sitename has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Sitename = d.Get("sitename").(string)
		hasChange = true
	}
	if d.HasChange("sitepersistence") {
		log.Printf("[DEBUG]  citrixadc-provider: Sitepersistence has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Sitepersistence = d.Get("sitepersistence").(string)
		hasChange = true
	}
	if d.HasChange("siteprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Siteprefix has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Siteprefix = d.Get("siteprefix").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("svrtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Svrtimeout has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Svrtimeout = intPtr(d.Get("svrtimeout").(int))
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  citrixadc-provider: Weight has changed for gslbservicegroup %s, starting update", gslbservicegroupName)
		gslbservicegroup.Weight = intPtr(d.Get("weight").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("gslbservicegroup", gslbservicegroupName, &gslbservicegroup)
		if err != nil {
			return diag.Errorf("Error updating gslbservicegroup %s", gslbservicegroupName)
		}
	}
	return readGslbservicegroupFunc(ctx, d, meta)
}

func deleteGslbservicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbservicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbservicegroupName := d.Id()
	err := client.DeleteResource("gslbservicegroup", gslbservicegroupName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
