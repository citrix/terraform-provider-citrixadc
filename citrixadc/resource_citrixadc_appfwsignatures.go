package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwsignatures() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwsignaturesFunc,
		Read:          readAppfwsignaturesFunc,
		Update:        updateAppfwsignaturesFunc,
		Delete:        deleteAppfwsignaturesFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"merge": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"mergedefault": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"preservedefactions": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sha1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vendortype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xslt": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autoenablenewsignatures": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ruleid": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func createAppfwsignaturesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwsignaturesFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwsignaturesName := d.Get("name").(string)

	appfwsignatures := appfw.Appfwsignatures{
		Comment:                 d.Get("comment").(string),
		Merge:                   d.Get("merge").(bool),
		Mergedefault:            d.Get("mergedefault").(bool),
		Name:                    d.Get("name").(string),
		Overwrite:               d.Get("overwrite").(bool),
		Preservedefactions:      d.Get("preservedefactions").(bool),
		Sha1:                    d.Get("sha1").(string),
		Src:                     d.Get("src").(string),
		Vendortype:              d.Get("vendortype").(string),
		Xslt:                    d.Get("xslt").(string),
		Autoenablenewsignatures: d.Get("autoenablenewsignatures").(string),
		Category:                d.Get("category").(string),
		Enabled:                 d.Get("enabled").(string),
		Action:                  toStringList(d.Get("action").([]interface{})),
	}

	appfwsignatures_update_obj := appfw.Appfwsignatures{
		Name:         appfwsignaturesName,
		Mergedefault: d.Get("mergedefault").(bool),
	}

	err := client.ActOnResource(service.Appfwsignatures.Type(), &appfwsignatures, "Import")
	if err != nil {
		return err
	}

	if _, ok := d.GetOk("ruleid"); ok {
		appfwsignatures.Ruleid = toIntegerList(d.Get("ruleid").([]interface{}))
		err := client.ActOnResource(service.Appfwsignatures.Type(), &appfwsignatures, "Import")
		if err != nil {
			return err
		}
	}

	err = client.ActOnResource(service.Appfwsignatures.Type(), &appfwsignatures_update_obj, "update")
	if err != nil {
		return err
	}

	d.SetId(appfwsignaturesName)

	err = readAppfwsignaturesFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwsignatures but we can't read it ?? %s", appfwsignaturesName)
		return nil
	}
	return nil
}

func updateAppfwsignaturesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppfwsignaturesFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwsignaturesName := d.Id()

	appfwsignatures := appfw.Appfwsignatures{
		Name:      appfwsignaturesName,
		Src:       d.Get("src").(string),
		Overwrite: d.Get("overwrite").(bool),
	}

	appfwsignatures_update_obj := appfw.Appfwsignatures{
		Name:         appfwsignaturesName,
		Mergedefault: d.Get("mergedefault").(bool),
	}

	hasChange := false

	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("merge") {
		log.Printf("[DEBUG]  citrixadc-provider: Merge has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Merge = d.Get("merge").(bool)
		hasChange = true
	}
	if d.HasChange("mergedefault") {
		log.Printf("[DEBUG]  citrixadc-provider: Mergedefault has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Mergedefault = d.Get("mergedefault").(bool)
		hasChange = true
	}
	if d.HasChange("preservedefactions") {
		log.Printf("[DEBUG]  citrixadc-provider: Preservedefactions has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Preservedefactions = d.Get("preservedefactions").(bool)
		hasChange = true
	}
	if d.HasChange("sha1") {
		log.Printf("[DEBUG]  citrixadc-provider: Sha1 has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Sha1 = d.Get("sha1").(string)
		hasChange = true
	}
	if d.HasChange("vendortype") {
		log.Printf("[DEBUG]  citrixadc-provider: Vendortype has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Vendortype = d.Get("vendortype").(string)
		hasChange = true
	}
	if d.HasChange("xslt") {
		log.Printf("[DEBUG]  citrixadc-provider: Xslt has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Xslt = d.Get("xslt").(string)
		hasChange = true
	}
	if d.HasChange("autoenablenewsignatures") {
		log.Printf("[DEBUG]  citrixadc-provider: Autoenablenewsignatures has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Autoenablenewsignatures = d.Get("autoenablenewsignatures").(string)
		hasChange = true
	}
	if d.HasChange("ruleid") {
		log.Printf("[DEBUG]  citrixadc-provider: Ruleid has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Ruleid = toIntegerList(d.Get("ruleid").([]interface{}))
		hasChange = true
	}
	if d.HasChange("category") {
		log.Printf("[DEBUG]  citrixadc-provider: Category has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Category = d.Get("category").(string)
		hasChange = true
	}
	if d.HasChange("enabled") {
		log.Printf("[DEBUG]  citrixadc-provider: Enabled has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Enabled = d.Get("enabled").(string)
		hasChange = true
	}
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for appfwsignatures %s, starting update", appfwsignaturesName)
		appfwsignatures.Action = toStringList(d.Get("action").([]interface{}))
		hasChange = true
	}

	if hasChange {
		err := client.ActOnResource(service.Appfwsignatures.Type(), &appfwsignatures, "Import")
		if err != nil {
			return err
		}

		err = client.ActOnResource(service.Appfwsignatures.Type(), &appfwsignatures_update_obj, "update")
		if err != nil {
			return err
		}
	}

	err := readAppfwsignaturesFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just updated this appfwsignatures but we can't read it ?? %s", appfwsignaturesName)
		return nil
	}
	return nil
}

func readAppfwsignaturesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwsignaturesFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwsignaturesName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwsignatures state %s", appfwsignaturesName)
	data, err := client.FindResource(service.Appfwsignatures.Type(), appfwsignaturesName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwsignatures state %s", appfwsignaturesName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteAppfwsignaturesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwsignaturesFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwsignaturesName := d.Id()
	err := client.DeleteResource(service.Appfwsignatures.Type(), appfwsignaturesName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
