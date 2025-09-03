package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcDnsprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsprofileFunc,
		Read:          readDnsprofileFunc,
		Update:        updateDnsprofileFunc,
		Delete:        deleteDnsprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cacheecsresponses": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachenegativeresponses": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacherecords": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsanswerseclogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnserrorlogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsextendedlogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dnsquerylogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropmultiqueryrequest": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recursiveresolution": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertecs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replaceecs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxcacheableecsprefixlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxcacheableecsprefixlength6": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsprofileName := d.Get("dnsprofilename").(string)
	dnsprofile := dns.Dnsprofile{
		Cacheecsresponses:            d.Get("cacheecsresponses").(string),
		Cachenegativeresponses:       d.Get("cachenegativeresponses").(string),
		Cacherecords:                 d.Get("cacherecords").(string),
		Dnsanswerseclogging:          d.Get("dnsanswerseclogging").(string),
		Dnserrorlogging:              d.Get("dnserrorlogging").(string),
		Dnsextendedlogging:           d.Get("dnsextendedlogging").(string),
		Dnsprofilename:               d.Get("dnsprofilename").(string),
		Dnsquerylogging:              d.Get("dnsquerylogging").(string),
		Dropmultiqueryrequest:        d.Get("dropmultiqueryrequest").(string),
		Recursiveresolution:          d.Get("recursiveresolution").(string),
		Insertecs:                    d.Get("insertecs").(string),
		Replaceecs:                   d.Get("replaceecs").(string),
		Maxcacheableecsprefixlength:  d.Get("maxcacheableecsprefixlength").(int),
		Maxcacheableecsprefixlength6: d.Get("maxcacheableecsprefixlength6").(int),
	}

	_, err := client.AddResource(service.Dnsprofile.Type(), dnsprofileName, &dnsprofile)
	if err != nil {
		return err
	}

	d.SetId(dnsprofileName)

	err = readDnsprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnsprofile but we can't read it ?? %s", dnsprofileName)
		return nil
	}
	return nil
}

func readDnsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsprofile state %s", dnsprofileName)
	data, err := client.FindResource(service.Dnsprofile.Type(), dnsprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsprofile state %s", dnsprofileName)
		d.SetId("")
		return nil
	}
	d.Set("dnsprofilename", data["dnsprofilename"])
	d.Set("cacheecsresponses", data["cacheecsresponses"])
	d.Set("cachenegativeresponses", data["cachenegativeresponses"])
	d.Set("cacherecords", data["cacherecords"])
	d.Set("dnsanswerseclogging", data["dnsanswerseclogging"])
	d.Set("dnserrorlogging", data["dnserrorlogging"])
	d.Set("dnsextendedlogging", data["dnsextendedlogging"])
	d.Set("dnsprofilename", data["dnsprofilename"])
	d.Set("dnsquerylogging", data["dnsquerylogging"])
	d.Set("dropmultiqueryrequest", data["dropmultiqueryrequest"])
	d.Set("recursiveresolution", data["recursiveresolution"])
	d.Set("insertecs", data["insertecs"])
	d.Set("replaceecs", data["replaceecs"])
	d.Set("maxcacheableecsprefixlength", data["maxcacheableecsprefixlength"])
	d.Set("maxcacheableecsprefixlength6", data["maxcacheableecsprefixlength6"])

	return nil

}

func updateDnsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsprofileName := d.Get("dnsprofilename").(string)

	dnsprofile := dns.Dnsprofile{
		Dnsprofilename: d.Get("dnsprofilename").(string),
	}
	hasChange := false
	if d.HasChange("cacheecsresponses") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacheecsresponses has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Cacheecsresponses = d.Get("cacheecsresponses").(string)
		hasChange = true
	}
	if d.HasChange("cachenegativeresponses") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachenegativeresponses has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Cachenegativeresponses = d.Get("cachenegativeresponses").(string)
		hasChange = true
	}
	if d.HasChange("cacherecords") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacherecords has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Cacherecords = d.Get("cacherecords").(string)
		hasChange = true
	}
	if d.HasChange("dnsanswerseclogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsanswerseclogging has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Dnsanswerseclogging = d.Get("dnsanswerseclogging").(string)
		hasChange = true
	}
	if d.HasChange("dnserrorlogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnserrorlogging has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Dnserrorlogging = d.Get("dnserrorlogging").(string)
		hasChange = true
	}
	if d.HasChange("dnsextendedlogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsextendedlogging has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Dnsextendedlogging = d.Get("dnsextendedlogging").(string)
		hasChange = true
	}
	if d.HasChange("dnsquerylogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsquerylogging has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Dnsquerylogging = d.Get("dnsquerylogging").(string)
		hasChange = true
	}
	if d.HasChange("dropmultiqueryrequest") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropmultiqueryrequest has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Dropmultiqueryrequest = d.Get("dropmultiqueryrequest").(string)
		hasChange = true
	}
	if d.HasChange("recursiveresolution") {
		log.Printf("[DEBUG]  citrixadc-provider: Recursiveresolution has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Recursiveresolution = d.Get("recursiveresolution").(string)
		hasChange = true
	}
	if d.HasChange("insertecs") {
		log.Printf("[DEBUG]  citrixadc-provider: Insertecs has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Insertecs = d.Get("insertecs").(string)
		hasChange = true
	}
	if d.HasChange("replaceecs") {
		log.Printf("[DEBUG]  citrixadc-provider: Replaceecs has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Replaceecs = d.Get("replaceecs").(string)
		hasChange = true
	}
	if d.HasChange("maxcacheableecsprefixlength") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxcacheableecsprefixlength has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Maxcacheableecsprefixlength = d.Get("maxcacheableecsprefixlength").(int)
		hasChange = true
	}
	if d.HasChange("maxcacheableecsprefixlength6") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxcacheableecsprefixlength6 has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Maxcacheableecsprefixlength6 = d.Get("maxcacheableecsprefixlength6").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnsprofile.Type(), dnsprofileName, &dnsprofile)
		if err != nil {
			return fmt.Errorf("Error updating dnsprofile %s", dnsprofileName)
		}
	}
	return readDnsprofileFunc(d, meta)
}

func deleteDnsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsprofileName := d.Id()
	err := client.DeleteResource(service.Dnsprofile.Type(), dnsprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
