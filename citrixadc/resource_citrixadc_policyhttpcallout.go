package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcPolicyhttpcallout() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPolicyhttpcalloutFunc,
		ReadContext:   readPolicyhttpcalloutFunc,
		UpdateContext: updatePolicyhttpcalloutFunc,
		DeleteContext: deletePolicyhttpcalloutFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"bodyexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacheforsecs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fullreqexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"headers": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"hostexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameters": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"resultexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"returntype": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scheme": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urlstemexpr": {
				Type:     schema.TypeString,
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

func createPolicyhttpcalloutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicyhttpcalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	var policyhttpcalloutName string
	if v, ok := d.GetOk("name"); ok {
		policyhttpcalloutName = v.(string)
	} else {
		policyhttpcalloutName = resource.PrefixedUniqueId("tf-policyhttpcallout-")
		d.Set("name", policyhttpcalloutName)
	}
	policyhttpcallout := policy.Policyhttpcallout{
		Bodyexpr:     d.Get("bodyexpr").(string),
		Cacheforsecs: d.Get("cacheforsecs").(int),
		Comment:      d.Get("comment").(string),
		Fullreqexpr:  d.Get("fullreqexpr").(string),
		Headers:      toStringList(d.Get("headers").([]interface{})),
		Hostexpr:     d.Get("hostexpr").(string),
		Httpmethod:   d.Get("httpmethod").(string),
		Ipaddress:    d.Get("ipaddress").(string),
		Name:         d.Get("name").(string),
		Parameters:   toStringList(d.Get("parameters").([]interface{})),
		Port:         d.Get("port").(int),
		Resultexpr:   d.Get("resultexpr").(string),
		Returntype:   d.Get("returntype").(string),
		Scheme:       d.Get("scheme").(string),
		Urlstemexpr:  d.Get("urlstemexpr").(string),
		Vserver:      d.Get("vserver").(string),
	}

	_, err := client.AddResource(service.Policyhttpcallout.Type(), policyhttpcalloutName, &policyhttpcallout)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policyhttpcalloutName)

	return readPolicyhttpcalloutFunc(ctx, d, meta)
}

func readPolicyhttpcalloutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicyhttpcalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	policyhttpcalloutName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading policyhttpcallout state %s", policyhttpcalloutName)
	data, err := client.FindResource(service.Policyhttpcallout.Type(), policyhttpcalloutName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policyhttpcallout state %s", policyhttpcalloutName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("bodyexpr", data["bodyexpr"])
	setToInt("cacheforsecs", d, data["cacheforsecs"])
	d.Set("comment", data["comment"])
	d.Set("fullreqexpr", data["fullreqexpr"])
	d.Set("headers", data["headers"])
	d.Set("hostexpr", data["hostexpr"])
	d.Set("httpmethod", data["httpmethod"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("name", data["name"])
	d.Set("parameters", data["parameters"])
	setToInt("port", d, data["port"])
	d.Set("resultexpr", data["resultexpr"])
	d.Set("returntype", data["returntype"])
	d.Set("scheme", data["scheme"])
	d.Set("urlstemexpr", data["urlstemexpr"])
	d.Set("vserver", data["vserver"])

	return nil

}

func updatePolicyhttpcalloutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePolicyhttpcalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	policyhttpcalloutName := d.Get("name").(string)

	policyhttpcallout := policy.Policyhttpcallout{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("bodyexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Bodyexpr has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Bodyexpr = d.Get("bodyexpr").(string)
		hasChange = true
	}
	if d.HasChange("cacheforsecs") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacheforsecs has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Cacheforsecs = d.Get("cacheforsecs").(int)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("fullreqexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Fullreqexpr has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Fullreqexpr = d.Get("fullreqexpr").(string)
		hasChange = true
	}
	if d.HasChange("headers") {
		log.Printf("[DEBUG]  citrixadc-provider: Headers has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Headers = toStringList(d.Get("headers").([]interface{}))
		hasChange = true
	}
	if d.HasChange("hostexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Hostexpr has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Hostexpr = d.Get("hostexpr").(string)
		hasChange = true
	}
	if d.HasChange("httpmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpmethod has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Httpmethod = d.Get("httpmethod").(string)
		hasChange = true
	}
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipaddress has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Ipaddress = d.Get("ipaddress").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG]  citrixadc-provider: Parameters has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Parameters = toStringList(d.Get("parameters").([]interface{}))
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("resultexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Resultexpr has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Resultexpr = d.Get("resultexpr").(string)
		hasChange = true
	}
	if d.HasChange("returntype") {
		log.Printf("[DEBUG]  citrixadc-provider: Returntype has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Returntype = d.Get("returntype").(string)
		hasChange = true
	}
	if d.HasChange("scheme") {
		log.Printf("[DEBUG]  citrixadc-provider: Scheme has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Scheme = d.Get("scheme").(string)
		hasChange = true
	}
	if d.HasChange("urlstemexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Urlstemexpr has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Urlstemexpr = d.Get("urlstemexpr").(string)
		hasChange = true
	}
	if d.HasChange("vserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserver has changed for policyhttpcallout %s, starting update", policyhttpcalloutName)
		policyhttpcallout.Vserver = d.Get("vserver").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Policyhttpcallout.Type(), policyhttpcalloutName, &policyhttpcallout)
		if err != nil {
			return diag.Errorf("Error updating policyhttpcallout %s", policyhttpcalloutName)
		}
	}
	return readPolicyhttpcalloutFunc(ctx, d, meta)
}

func deletePolicyhttpcalloutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicyhttpcalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	policyhttpcalloutName := d.Id()
	err := client.DeleteResource(service.Policyhttpcallout.Type(), policyhttpcalloutName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
