package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcAppfwprofile_excluderescontenttype_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_excluderescontenttype_bindingFunc,
		Read:          readAppfwprofile_excluderescontenttype_bindingFunc,
		Delete:        deleteAppfwprofile_excluderescontenttype_bindingFunc,
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
			"excluderescontenttype": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"alertonly": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isautodeployed": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resourceid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ruletype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_excluderescontenttype_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_excluderescontenttype_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	excluderescontenttype := d.Get("excluderescontenttype")
	bindingId := fmt.Sprintf("%s,%s", name, excluderescontenttype)
	appfwprofile_excluderescontenttype_binding := appfw.Appfwprofileexcluderescontenttypebinding{
		Alertonly:             d.Get("alertonly").(string),
		Comment:               d.Get("comment").(string),
		Excluderescontenttype: d.Get("excluderescontenttype").(string),
		Isautodeployed:        d.Get("isautodeployed").(string),
		Name:                  d.Get("name").(string),
		Resourceid:            d.Get("resourceid").(string),
		Ruletype:              d.Get("ruletype").(string),
		State:                 d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_excluderescontenttype_binding.Type(), &appfwprofile_excluderescontenttype_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_excluderescontenttype_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_excluderescontenttype_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_excluderescontenttype_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_excluderescontenttype_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	excluderescontenttype := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_excluderescontenttype_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_excluderescontenttype_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_excluderescontenttype_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["excluderescontenttype"].(string) == excluderescontenttype {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_excluderescontenttype_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	d.Set("excluderescontenttype", data["excluderescontenttype"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_excluderescontenttype_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_excluderescontenttype_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	excluderescontenttype := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("excluderescontenttype:%s", excluderescontenttype))
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_excluderescontenttype_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
