package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
	"net/url"
)

func resourceCitrixAdcAppfwprofile_fieldconsistency_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_fieldconsistency_bindingFunc,
		Read:          readAppfwprofile_fieldconsistency_bindingFunc,
		Delete:        deleteAppfwprofile_fieldconsistency_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"fieldconsistency": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"formactionurl_ffc": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"alertonly": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isautodeployed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isregex_ffc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resourceid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ruletype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_fieldconsistency_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_fieldconsistency_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	fieldconsistency := d.Get("fieldconsistency")
	formactionurl_ffc := d.Get("formactionurl_ffc")
	bindingId := fmt.Sprintf("%s,%s,%s", name, fieldconsistency,formactionurl_ffc)
	appfwprofile_fieldconsistency_binding := appfw.Appfwprofilefieldconsistencybinding{
		Alertonly:        d.Get("alertonly").(string),
		Comment:          d.Get("comment").(string),
		Fieldconsistency: d.Get("fieldconsistency").(string),
		Formactionurlffc: d.Get("formactionurl_ffc").(string),
		Isautodeployed:   d.Get("isautodeployed").(string),
		Isregexffc:       d.Get("isregex_ffc").(string),
		Name:             d.Get("name").(string),
		Resourceid:       d.Get("resourceid").(string),
		Ruletype:         d.Get("ruletype").(string),
		State:            d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_fieldconsistency_binding.Type(), &appfwprofile_fieldconsistency_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_fieldconsistency_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_fieldconsistency_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_fieldconsistency_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_fieldconsistency_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	fieldconsistency := idSlice[1]
	formactionurl_ffc := idSlice[2]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_fieldconsistency_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_fieldconsistency_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_fieldconsistency_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["fieldconsistency"].(string) == fieldconsistency {
			if v["formactionurl_ffc"].(string) == formactionurl_ffc {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_fieldconsistency_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	d.Set("fieldconsistency", data["fieldconsistency"])
	d.Set("formactionurl_ffc", data["formactionurl_ffc"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex_ffc", data["isregex_ffc"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_fieldconsistency_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_fieldconsistency_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	fieldconsistency := idSlice[1]
	formactionurl_ffc := idSlice[2]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("fieldconsistency:%s", fieldconsistency))
	args = append(args, fmt.Sprintf("formactionurl_ffc:%s", url.QueryEscape(formactionurl_ffc)))

	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_fieldconsistency_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
