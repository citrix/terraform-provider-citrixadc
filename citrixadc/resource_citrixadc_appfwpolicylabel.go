package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwpolicylabelFunc,
		Read:          readAppfwpolicylabelFunc,
		Update:        updateAppfwpolicyLabelFunc,
		Delete:        deleteAppfwpolicylabelFunc,
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policylabeltype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppfwpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicylabelName := d.Get("labelname").(string)
	appfwpolicylabel := appfw.Appfwpolicylabel{
		Labelname:       appfwpolicylabelName,
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	_, err := client.AddResource(service.Appfwpolicylabel.Type(), appfwpolicylabelName, &appfwpolicylabel)
	if err != nil {
		return err
	}

	d.SetId(appfwpolicylabelName)

	err = readAppfwpolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwpolicylabel but we can't read it ?? %s", appfwpolicylabelName)
		return nil
	}
	return nil
}

func readAppfwpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwpolicylabel state %s", appfwpolicylabelName)
	data, err := client.FindResource(service.Appfwpolicylabel.Type(), appfwpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwpolicylabel state %s", appfwpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("policylabeltype", data["policylabeltype"])

	return nil

}

func updateAppfwpolicyLabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppfwpolicyLabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicylabelName := d.Get("labelname").(string)

	appfwpolicylabel := appfw.Appfwpolicylabel{
		Labelname: appfwpolicylabelName,
	}
	hasRename := false
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for appfwpolicylabel %s, starting rename", appfwpolicylabelName)
		appfwpolicylabel.Newname = d.Get("newname").(string)
		hasRename = true
	}

	if hasRename {
		err := client.ActOnResource(service.Appfwpolicylabel.Type(), &appfwpolicylabel, "rename")
		if err != nil {
			return err
		}
		d.SetId(appfwpolicylabel.Newname)
	}
	return readAppfwpolicylabelFunc(d, meta)
}

func deleteAppfwpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicylabelName := d.Id()
	err := client.DeleteResource(service.Appfwpolicylabel.Type(), appfwpolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
