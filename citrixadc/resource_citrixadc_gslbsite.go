package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcGslbsite() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbsiteFunc,
		Read:          readGslbsiteFunc,
		Update:        updateGslbsiteFunc,
		Delete:        deleteGslbsiteFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"clip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metricexchange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"naptrreplacementsuffix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nwmetricexchange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parentsite": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publicclip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publicip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionexchange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"siteipaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"triggermonitor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backupparentlist": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sitepassword": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createGslbsiteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createGslbsiteFunc")
	client := meta.(*NetScalerNitroClient).client
	var gslbsiteName string
	if v, ok := d.GetOk("sitename"); ok {
		gslbsiteName = v.(string)
	} else {
		gslbsiteName = resource.PrefixedUniqueId("tf-gslbsite-")
		d.Set("sitename", gslbsiteName)
	}
	gslbsite := gslb.Gslbsite{
		Clip:                   d.Get("clip").(string),
		Metricexchange:         d.Get("metricexchange").(string),
		Naptrreplacementsuffix: d.Get("naptrreplacementsuffix").(string),
		Nwmetricexchange:       d.Get("nwmetricexchange").(string),
		Parentsite:             d.Get("parentsite").(string),
		Publicclip:             d.Get("publicclip").(string),
		Publicip:               d.Get("publicip").(string),
		Sessionexchange:        d.Get("sessionexchange").(string),
		Siteipaddress:          d.Get("siteipaddress").(string),
		Sitename:               d.Get("sitename").(string),
		Sitetype:               d.Get("sitetype").(string),
		Triggermonitor:         d.Get("triggermonitor").(string),
		Sitepassword:           d.Get("sitepassword").(string),
	}
	if listVal, ok := d.Get("backupparentlist").([]interface{}); ok {
		gslbsite.Backupparentlist = toStringList(listVal)
	}
	_, err := client.AddResource(service.Gslbsite.Type(), gslbsiteName, &gslbsite)
	if err != nil {
		return err
	}

	d.SetId(gslbsiteName)

	err = readGslbsiteFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbsite but we can't read it ?? %s", gslbsiteName)
		return nil
	}
	return nil
}

func readGslbsiteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readGslbsiteFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbsiteName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading gslbsite state %s", gslbsiteName)
	data, err := client.FindResource(service.Gslbsite.Type(), gslbsiteName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing gslbsite state %s", gslbsiteName)
		d.SetId("")
		return nil
	}
	d.Set("sitename", data["sitename"])
	d.Set("clip", data["clip"])
	d.Set("metricexchange", data["metricexchange"])
	d.Set("naptrreplacementsuffix", data["naptrreplacementsuffix"])
	d.Set("nwmetricexchange", data["nwmetricexchange"])
	d.Set("parentsite", data["parentsite"])
	d.Set("publicclip", data["publicclip"])
	d.Set("publicip", data["publicip"])
	d.Set("sessionexchange", data["sessionexchange"])
	d.Set("siteipaddress", data["siteipaddress"])
	d.Set("sitetype", data["sitetype"])
	d.Set("triggermonitor", data["triggermonitor"])
	d.Set("sitepassword", d.Get("sitepassword").(string))
	d.Set("status", data["status"])
	d.Set("persistencemepstatus", data["persistencemepstatus"])
	d.Set("version", data["version"])
	d.Set("curbackupparentip", data["curbackupparentip"])
	d.Set("sitestate", data["sitestate"])
	d.Set("oldname", data["oldname"])
	d.Set("nextgenapiresource", data["_nextgenapiresource"])
	if val, ok := data["backupparentlist"]; ok {
		if list, ok := val.([]interface{}); ok {
			d.Set("backupparentlist", toStringList(list))
		}
	} else {
		d.Set("backupparentlist", nil)
	}

	return nil

}

func updateGslbsiteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateGslbsiteFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbsiteName := d.Get("sitename").(string)

	gslbsite := gslb.Gslbsite{
		Sitename: gslbsiteName,
	}
	hasRename := false
	hasChange := false
	if d.HasChange("metricexchange") {
		log.Printf("[DEBUG]  netscaler-provider: Metricexchange has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Metricexchange = d.Get("metricexchange").(string)
		hasChange = true
	}
	if d.HasChange("naptrreplacementsuffix") {
		log.Printf("[DEBUG]  netscaler-provider: Naptrreplacementsuffix has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Naptrreplacementsuffix = d.Get("naptrreplacementsuffix").(string)
		hasChange = true
	}
	if d.HasChange("nwmetricexchange") {
		log.Printf("[DEBUG]  netscaler-provider: Nwmetricexchange has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Nwmetricexchange = d.Get("nwmetricexchange").(string)
		hasChange = true
	}
	if d.HasChange("parentsite") {
		log.Printf("[DEBUG]  netscaler-provider: Parentsite has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Parentsite = d.Get("parentsite").(string)
		hasChange = true
	}
	if d.HasChange("publicip") {
		log.Printf("[DEBUG]  netscaler-provider: Publicip has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Publicip = d.Get("publicip").(string)
		hasChange = true
	}
	if d.HasChange("sessionexchange") {
		log.Printf("[DEBUG]  netscaler-provider: Sessionexchange has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Sessionexchange = d.Get("sessionexchange").(string)
		hasChange = true
	}
	if d.HasChange("siteipaddress") {
		log.Printf("[DEBUG]  netscaler-provider: Siteipaddress has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Siteipaddress = d.Get("siteipaddress").(string)
		hasChange = true
	}
	if d.HasChange("sitename") {
		log.Printf("[DEBUG]  netscaler-provider: Sitename has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Sitename = gslbsiteName
		hasChange = true
	}
	if d.HasChange("triggermonitor") {
		log.Printf("[DEBUG]  netscaler-provider: Triggermonitor has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Triggermonitor = d.Get("triggermonitor").(string)
		hasChange = true
	}
	if d.HasChange("backupparentlist") {
		log.Printf("[DEBUG]  netscaler-provider: Backupparentlist has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Backupparentlist = toStringList(d.Get("backupparentlist").([]interface{}))
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  netscaler-provider: Newname has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Newname = d.Get("newname").(string)
		hasRename = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Gslbsite.Type(), gslbsiteName, &gslbsite)
		if err != nil {
			return err
		}
	}

	if hasRename {
		err := client.ActOnResource(service.Gslbsite.Type(), &gslbsite, "rename")
		if err != nil {
			return err
		}
		d.SetId(gslbsite.Newname)
	}

	return readGslbsiteFunc(d, meta)
}

func deleteGslbsiteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteGslbsiteFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbsiteName := d.Id()
	err := client.DeleteResource(service.Gslbsite.Type(), gslbsiteName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
