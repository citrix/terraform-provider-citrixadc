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
			"cacheecsresponses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachenegativeresponses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacherecords": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsanswerseclogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnserrorlogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsextendedlogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dnsquerylogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropmultiqueryrequest": &schema.Schema{
				Type:     schema.TypeString,
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
		Cacheecsresponses:      d.Get("cacheecsresponses").(string),
		Cachenegativeresponses: d.Get("cachenegativeresponses").(string),
		Cacherecords:           d.Get("cacherecords").(string),
		Dnsanswerseclogging:    d.Get("dnsanswerseclogging").(string),
		Dnserrorlogging:        d.Get("dnserrorlogging").(string),
		Dnsextendedlogging:     d.Get("dnsextendedlogging").(string),
		Dnsprofilename:         d.Get("dnsprofilename").(string),
		Dnsquerylogging:        d.Get("dnsquerylogging").(string),
		Dropmultiqueryrequest:  d.Get("dropmultiqueryrequest").(string),
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
