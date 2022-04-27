package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
	"net/url"
)

func resourceCitrixAdcAppfwprofile_cmdinjection_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_cmdinjection_bindingFunc,
		Read:          readAppfwprofile_cmdinjection_bindingFunc,
		Delete:        deleteAppfwprofile_cmdinjection_bindingFunc,
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
			"cmdinjection": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"formactionurl_cmd": &schema.Schema{
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
			"as_scan_location_cmd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_expr_cmd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_type_cmd": &schema.Schema{
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
			"isregex_cmd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isvalueregex_cmd": &schema.Schema{
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

func createAppfwprofile_cmdinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_cmdinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	cmdinjection := d.Get("cmdinjection")
	formactionurl_cmd := d.Get("formactionurl_cmd")
	bindingId := fmt.Sprintf("%s,%s,%s", name, cmdinjection, formactionurl_cmd)
	appfwprofile_cmdinjection_binding := appfw.Appfwprofilecmdinjectionbinding{
		Alertonly:         d.Get("alertonly").(string),
		Asscanlocationcmd: d.Get("as_scan_location_cmd").(string),
		Asvalueexprcmd:    d.Get("as_value_expr_cmd").(string),
		Asvaluetypecmd:    d.Get("as_value_type_cmd").(string),
		Cmdinjection:      d.Get("cmdinjection").(string),
		Comment:           d.Get("comment").(string),
		Formactionurlcmd:  d.Get("formactionurl_cmd").(string),
		Isautodeployed:    d.Get("isautodeployed").(string),
		Isregexcmd:        d.Get("isregex_cmd").(string),
		Isvalueregexcmd:   d.Get("isvalueregex_cmd").(string),
		Name:              d.Get("name").(string),
		Resourceid:        d.Get("resourceid").(string),
		Ruletype:          d.Get("ruletype").(string),
		State:             d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource("appfwprofile_cmdinjection_binding", &appfwprofile_cmdinjection_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_cmdinjection_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_cmdinjection_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_cmdinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_cmdinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	cmdinjection := idSlice[1]
	formactionurl_cmd := idSlice[2]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_cmdinjection_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_cmdinjection_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_cmdinjection_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true
		if v["cmdinjection"].(string) != cmdinjection {
			match = false
		}
		if v["formactionurl_cmd"].(string) != formactionurl_cmd {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_cmdinjection_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("alertonly", data["alertonly"])
	d.Set("as_scan_location_cmd", data["as_scan_location_cmd"])
	d.Set("as_value_expr_cmd", data["as_value_expr_cmd"])
	d.Set("as_value_type_cmd", data["as_value_type_cmd"])
	d.Set("cmdinjection", data["cmdinjection"])
	d.Set("comment", data["comment"])
	d.Set("formactionurl_cmd", data["formactionurl_cmd"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex_cmd", data["isregex_cmd"])
	d.Set("isvalueregex_cmd", data["isvalueregex_cmd"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_cmdinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_cmdinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	cmdinjection := idSlice[1]
	formactionurl_cmd := idSlice[2]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("cmdinjection:%s", cmdinjection))
	args = append(args, fmt.Sprintf("formactionurl_cmd:%s", url.QueryEscape(formactionurl_cmd)))
	
	if val, ok := d.GetOk("as_scan_location_cmd"); ok {
		args = append(args, fmt.Sprintf("as_scan_location_cmd:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("as_value_type_cmd"); ok {
		args = append(args, fmt.Sprintf("as_value_type_cmd:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("as_value_expr_cmd"); ok {
		args = append(args, fmt.Sprintf("as_value_expr_cmd:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs("appfwprofile_cmdinjection_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
