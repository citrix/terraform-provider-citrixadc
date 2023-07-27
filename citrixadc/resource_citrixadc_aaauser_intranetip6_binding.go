package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
)

func resourceCitrixAdcAaauser_intranetip6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaauser_intranetip6_bindingFunc,
		Read:          readAaauser_intranetip6_bindingFunc,
		Delete:        deleteAaauser_intranetip6_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"intranetip6": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"numaddr": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAaauser_intranetip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaauser_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	username := d.Get("username").(string)
	// checks if user provided CIDR in intranetip6 field
	intranetip6 := d.Get("intranetip6").(string)
	matched, regerr := regexp.MatchString(`.*/`, intranetip6)
	if regerr != nil {
		log.Printf("[ERROR] Regular expression error:%s", regerr)
		return nil
	}
	//remove CIDR if provided
	if matched {
		re := regexp.MustCompile(`.*/`)
		intranetip6 = strings.TrimSuffix(re.FindString(intranetip6), "/")
	}
	//append the correct CIDR
	intranetip6 = intranetip6 + "/128"
	bindingId := fmt.Sprintf("%s,%s", username, intranetip6)
	aaauser_intranetip6_binding := aaa.Aaauserintranetip6binding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Intranetip6:            d.Get("intranetip6").(string),
		Numaddr:                d.Get("numaddr").(int),
		Username:               d.Get("username").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaauser_intranetip6_binding.Type(), &aaauser_intranetip6_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAaauser_intranetip6_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaauser_intranetip6_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAaauser_intranetip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaauser_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	username := idSlice[0]
	intranetip6 := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading aaauser_intranetip6_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "aaauser_intranetip6_binding",
		ResourceName:             username,
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser_intranetip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetip6"].(string) == intranetip6 {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams intranetip6 not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser_intranetip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("intranetip6", data["intranetip6"])
	d.Set("numaddr", data["numaddr"])
	d.Set("username", data["username"])

	return nil

}

func deleteAaauser_intranetip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaauser_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	intranetip6 := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetip6:%s", url.PathEscape(intranetip6)))
	if v, ok := d.GetOk("numaddr"); ok {
		numaddr := v.(int)
		args = append(args, fmt.Sprintf("numaddr:%v", numaddr))
	}

	err := client.DeleteResourceWithArgs(service.Aaauser_intranetip6_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
