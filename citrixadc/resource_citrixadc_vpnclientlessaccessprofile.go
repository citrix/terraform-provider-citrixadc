package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcVpnclientlessaccessprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnclientlessaccessprofileFunc,
		ReadContext:   readVpnclientlessaccessprofileFunc,
		UpdateContext: updateVpnclientlessaccessprofileFunc,
		DeleteContext: deleteVpnclientlessaccessprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"clientconsumedcookies": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"javascriptrewritepolicylabel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"regexforfindingcustomurls": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"regexforfindingurlincss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"regexforfindingurlinjavascript": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"regexforfindingurlinxcomponent": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"regexforfindingurlinxml": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reqhdrrewritepolicylabel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requirepersistentcookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reshdrrewritepolicylabel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urlrewritepolicylabel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpnclientlessaccessprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnclientlessaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	var vpnclientlessaccessprofileName string
	if v, ok := d.GetOk("profilename"); ok {
		vpnclientlessaccessprofileName = v.(string)
	} else {
		vpnclientlessaccessprofileName = resource.PrefixedUniqueId("tf-vpnclientlessaccessprofile-")
		d.Set("profilename", vpnclientlessaccessprofileName)
	}
	vpnclientlessaccessprofile := vpn.Vpnclientlessaccessprofile{
		Clientconsumedcookies:          d.Get("clientconsumedcookies").(string),
		Javascriptrewritepolicylabel:   d.Get("javascriptrewritepolicylabel").(string),
		Profilename:                    d.Get("profilename").(string),
		Regexforfindingcustomurls:      d.Get("regexforfindingcustomurls").(string),
		Regexforfindingurlincss:        d.Get("regexforfindingurlincss").(string),
		Regexforfindingurlinjavascript: d.Get("regexforfindingurlinjavascript").(string),
		Regexforfindingurlinxcomponent: d.Get("regexforfindingurlinxcomponent").(string),
		Regexforfindingurlinxml:        d.Get("regexforfindingurlinxml").(string),
		Reqhdrrewritepolicylabel:       d.Get("reqhdrrewritepolicylabel").(string),
		Requirepersistentcookie:        d.Get("requirepersistentcookie").(string),
		Reshdrrewritepolicylabel:       d.Get("reshdrrewritepolicylabel").(string),
		Urlrewritepolicylabel:          d.Get("urlrewritepolicylabel").(string),
	}

	_, err := client.AddResource(service.Vpnclientlessaccessprofile.Type(), vpnclientlessaccessprofileName, &vpnclientlessaccessprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vpnclientlessaccessprofileName)

	return readVpnclientlessaccessprofileFunc(ctx, d, meta)
}

func readVpnclientlessaccessprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnclientlessaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnclientlessaccessprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnclientlessaccessprofile state %s", vpnclientlessaccessprofileName)
	data, err := client.FindResource(service.Vpnclientlessaccessprofile.Type(), vpnclientlessaccessprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnclientlessaccessprofile state %s", vpnclientlessaccessprofileName)
		d.SetId("")
		return nil
	}
	d.Set("profilename", data["profilename"])
	d.Set("clientconsumedcookies", data["clientconsumedcookies"])
	d.Set("javascriptrewritepolicylabel", data["javascriptrewritepolicylabel"])
	d.Set("profilename", data["profilename"])
	d.Set("regexforfindingcustomurls", data["regexforfindingcustomurls"])
	d.Set("regexforfindingurlincss", data["regexforfindingurlincss"])
	d.Set("regexforfindingurlinjavascript", data["regexforfindingurlinjavascript"])
	d.Set("regexforfindingurlinxcomponent", data["regexforfindingurlinxcomponent"])
	d.Set("regexforfindingurlinxml", data["regexforfindingurlinxml"])
	d.Set("reqhdrrewritepolicylabel", data["reqhdrrewritepolicylabel"])
	d.Set("requirepersistentcookie", data["requirepersistentcookie"])
	d.Set("reshdrrewritepolicylabel", data["reshdrrewritepolicylabel"])
	d.Set("urlrewritepolicylabel", data["urlrewritepolicylabel"])

	return nil

}

func updateVpnclientlessaccessprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnclientlessaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnclientlessaccessprofileName := d.Get("profilename").(string)

	vpnclientlessaccessprofile := vpn.Vpnclientlessaccessprofile{
		Profilename: d.Get("profilename").(string),
	}
	hasChange := false
	if d.HasChange("clientconsumedcookies") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientconsumedcookies has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Clientconsumedcookies = d.Get("clientconsumedcookies").(string)
		hasChange = true
	}
	if d.HasChange("javascriptrewritepolicylabel") {
		log.Printf("[DEBUG]  citrixadc-provider: Javascriptrewritepolicylabel has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Javascriptrewritepolicylabel = d.Get("javascriptrewritepolicylabel").(string)
		hasChange = true
	}
	if d.HasChange("profilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Profilename has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Profilename = d.Get("profilename").(string)
		hasChange = true
	}
	if d.HasChange("regexforfindingcustomurls") {
		log.Printf("[DEBUG]  citrixadc-provider: Regexforfindingcustomurls has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Regexforfindingcustomurls = d.Get("regexforfindingcustomurls").(string)
		hasChange = true
	}
	if d.HasChange("regexforfindingurlincss") {
		log.Printf("[DEBUG]  citrixadc-provider: Regexforfindingurlincss has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Regexforfindingurlincss = d.Get("regexforfindingurlincss").(string)
		hasChange = true
	}
	if d.HasChange("regexforfindingurlinjavascript") {
		log.Printf("[DEBUG]  citrixadc-provider: Regexforfindingurlinjavascript has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Regexforfindingurlinjavascript = d.Get("regexforfindingurlinjavascript").(string)
		hasChange = true
	}
	if d.HasChange("regexforfindingurlinxcomponent") {
		log.Printf("[DEBUG]  citrixadc-provider: Regexforfindingurlinxcomponent has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Regexforfindingurlinxcomponent = d.Get("regexforfindingurlinxcomponent").(string)
		hasChange = true
	}
	if d.HasChange("regexforfindingurlinxml") {
		log.Printf("[DEBUG]  citrixadc-provider: Regexforfindingurlinxml has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Regexforfindingurlinxml = d.Get("regexforfindingurlinxml").(string)
		hasChange = true
	}
	if d.HasChange("reqhdrrewritepolicylabel") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqhdrrewritepolicylabel has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Reqhdrrewritepolicylabel = d.Get("reqhdrrewritepolicylabel").(string)
		hasChange = true
	}
	if d.HasChange("requirepersistentcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Requirepersistentcookie has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Requirepersistentcookie = d.Get("requirepersistentcookie").(string)
		hasChange = true
	}
	if d.HasChange("reshdrrewritepolicylabel") {
		log.Printf("[DEBUG]  citrixadc-provider: Reshdrrewritepolicylabel has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Reshdrrewritepolicylabel = d.Get("reshdrrewritepolicylabel").(string)
		hasChange = true
	}
	if d.HasChange("urlrewritepolicylabel") {
		log.Printf("[DEBUG]  citrixadc-provider: Urlrewritepolicylabel has changed for vpnclientlessaccessprofile %s, starting update", vpnclientlessaccessprofileName)
		vpnclientlessaccessprofile.Urlrewritepolicylabel = d.Get("urlrewritepolicylabel").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Vpnclientlessaccessprofile.Type(), vpnclientlessaccessprofileName, &vpnclientlessaccessprofile)
		if err != nil {
			return diag.Errorf("Error updating vpnclientlessaccessprofile %s", vpnclientlessaccessprofileName)
		}
	}
	return readVpnclientlessaccessprofileFunc(ctx, d, meta)
}

func deleteVpnclientlessaccessprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnclientlessaccessprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnclientlessaccessprofileName := d.Id()
	err := client.DeleteResource(service.Vpnclientlessaccessprofile.Type(), vpnclientlessaccessprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
