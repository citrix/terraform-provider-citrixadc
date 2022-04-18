package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/bot"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
	"strconv"
	"net/url"
)

func resourceCitrixAdcBotpolicylabel_botpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBotpolicylabel_botpolicy_bindingFunc,
		Read:          readBotpolicylabel_botpolicy_bindingFunc,
		Delete:        deleteBotpolicylabel_botpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"invoke": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"invokelabelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBotpolicylabel_botpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotpolicylabel_botpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	labelname := d.Get("labelname")
	policyname := d.Get("policyname")
	bindingId := fmt.Sprintf("%s,%s", labelname, policyname)
	botpolicylabel_botpolicy_binding := bot.Botpolicylabelbotpolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Invokelabelname:        d.Get("invokelabelname").(string),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource("botpolicylabel_botpolicy_binding", &botpolicylabel_botpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readBotpolicylabel_botpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this botpolicylabel_botpolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readBotpolicylabel_botpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotpolicylabel_botpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	labelname := idSlice[0]
	policyname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading botpolicylabel_botpolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "botpolicylabel_botpolicy_binding",
		ResourceName:             labelname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing botpolicylabel_botpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == policyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing botpolicylabel_botpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("invokelabelname", data["invokelabelname"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("policyname", data["policyname"])
	d.Set("priority", data["priority"])

	return nil

}

func deleteBotpolicylabel_botpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotpolicylabel_botpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	labelname := idSlice[0]
	policyname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))
	if val, ok := d.GetOk("td"); ok {
		args = append(args, fmt.Sprintf("td:%s", url.QueryEscape(strconv.Itoa(val.(int)))))
	}

	err := client.DeleteResourceWithArgs("botpolicylabel_botpolicy_binding", labelname, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
