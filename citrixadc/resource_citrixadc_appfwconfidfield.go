package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
)

func resourceCitrixAdcAppfwconfidfield() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwconfidfieldFunc,
		Read:          readAppfwconfidfieldFunc,
		Update:        updateAppfwconfidfieldFunc,
		Delete:        deleteAppfwconfidfieldFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"fieldname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"url": {
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
			"isregex": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppfwconfidfieldFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwconfidfieldFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwconfidfieldName := d.Get("fieldname").(string)
	appfwconfidfieldUrl := d.Get("url").(string)
	bindingId := fmt.Sprintf("%s,%s", appfwconfidfieldName, appfwconfidfieldUrl)
	appfwconfidfield := appfw.Appfwconfidfield{
		Comment:   d.Get("comment").(string),
		Fieldname: d.Get("fieldname").(string),
		Isregex:   d.Get("isregex").(string),
		State:     d.Get("state").(string),
		Url:       d.Get("url").(string),
	}

	_, err := client.AddResource(service.Appfwconfidfield.Type(), appfwconfidfieldName, &appfwconfidfield)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwconfidfieldFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwconfidfield but we can't read it ?? %s", appfwconfidfieldName)
		return nil
	}
	return nil
}

func readAppfwconfidfieldFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwconfidfieldFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwconfidfieldName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwconfidfield state %s", appfwconfidfieldName)
	findParams := service.FindParams{
		ResourceType: service.Appfwconfidfield.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwconfidfield state %s", appfwconfidfieldName)
		d.SetId("")
		return nil
	}

	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: appfw confidfield does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, appfwconfidfield := range dataArray {
		match := true
		if appfwconfidfield["fieldname"] != d.Get("fieldname").(string) {
			match = false
		}
		if appfwconfidfield["url"] != d.Get("url").(string) {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams appfwconfidfield not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwconfidfield state %s", appfwconfidfieldName)
		d.SetId("")
		return nil
	}
	data := dataArray[foundIndex]
	d.Set("comment", data["comment"])
	d.Set("fieldname", data["fieldname"])
	d.Set("isregex", data["isregex"])
	d.Set("state", data["state"])
	d.Set("url", data["url"])

	return nil
}

func updateAppfwconfidfieldFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppfwconfidfieldFunc")
	client := meta.(*NetScalerNitroClient).client

	appfwconfidfieldName := d.Get("fieldname").(string)
	appfwconfidfield := appfw.Appfwconfidfield{}
	log.Println(appfwconfidfield)
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for appfwconfidfield %s, starting update", appfwconfidfieldName)
		appfwconfidfield.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("isregex") {
		log.Printf("[DEBUG]  citrixadc-provider: Isregex has changed for appfwconfidfield %s, starting update", appfwconfidfieldName)
		appfwconfidfield.Isregex = d.Get("isregex").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for appfwconfidfield %s, starting update", appfwconfidfieldName)
		appfwconfidfield.State = d.Get("state").(string)
		hasChange = true
	}

	if hasChange {
		appfwconfidfield.Fieldname = d.Get("fieldname").(string)
		appfwconfidfield.Url = d.Get("url").(string)

		err := client.UpdateUnnamedResource(service.Appfwconfidfield.Type(), &appfwconfidfield)
		if err != nil {
			return fmt.Errorf("Error updating appfwconfidfield %s", appfwconfidfieldName)
		}
	}
	return readAppfwconfidfieldFunc(d, meta)
}

func deleteAppfwconfidfieldFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwconfidfieldFunc")
	client := meta.(*NetScalerNitroClient).client
	argsMap := make(map[string]string)
	// Only the fieldname and url properties are required for deletion
	argsMap["fieldname"] = url.QueryEscape(d.Get("fieldname").(string))
	argsMap["url"] = url.QueryEscape(d.Get("url").(string))

	err := client.DeleteResourceWithArgsMap(service.Appfwconfidfield.Type(), "", argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
