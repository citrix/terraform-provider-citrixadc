package netscaler

import (
	"github.com/chiradeep/go-nitro/config/gslb"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerGslbsite() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbsiteFunc,
		Read:          readGslbsiteFunc,
		Update:        updateGslbsiteFunc,
		Delete:        deleteGslbsiteFunc,
		Schema: map[string]*schema.Schema{
			"clip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metricexchange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"naptrreplacementsuffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nwmetricexchange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parentsite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publicclip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publicip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sessionexchange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"siteipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sitename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"triggermonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
	}

	_, err := client.AddResource(netscaler.Gslbsite.Type(), gslbsiteName, &gslbsite)
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
	data, err := client.FindResource(netscaler.Gslbsite.Type(), gslbsiteName)
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
	d.Set("sitename", data["sitename"])
	d.Set("sitetype", data["sitetype"])
	d.Set("triggermonitor", data["triggermonitor"])

	return nil

}

func updateGslbsiteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateGslbsiteFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbsiteName := d.Get("sitename").(string)

	gslbsite := gslb.Gslbsite{
		Sitename: d.Get("sitename").(string),
	}
	hasChange := false
	if d.HasChange("clip") {
		log.Printf("[DEBUG]  netscaler-provider: Clip has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Clip = d.Get("clip").(string)
		hasChange = true
	}
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
	if d.HasChange("publicclip") {
		log.Printf("[DEBUG]  netscaler-provider: Publicclip has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Publicclip = d.Get("publicclip").(string)
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
		gslbsite.Sitename = d.Get("sitename").(string)
		hasChange = true
	}
	if d.HasChange("sitetype") {
		log.Printf("[DEBUG]  netscaler-provider: Sitetype has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Sitetype = d.Get("sitetype").(string)
		hasChange = true
	}
	if d.HasChange("triggermonitor") {
		log.Printf("[DEBUG]  netscaler-provider: Triggermonitor has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Triggermonitor = d.Get("triggermonitor").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Gslbsite.Type(), gslbsiteName, &gslbsite)
		if err != nil {
			return fmt.Errorf("Error updating gslbsite %s", gslbsiteName)
		}
	}
	return readGslbsiteFunc(d, meta)
}

func deleteGslbsiteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteGslbsiteFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbsiteName := d.Id()
	err := client.DeleteResource(netscaler.Gslbsite.Type(), gslbsiteName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
