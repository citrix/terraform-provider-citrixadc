package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcGslbservicegroup_gslbservicegroupmember_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbservicegroup_gslbservicegroupmember_bindingFunc,
		Read:          readGslbservicegroup_gslbservicegroupmember_bindingFunc,
		Delete:        deleteGslbservicegroup_gslbservicegroupmember_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publicport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servername": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"ip"},
			},
			"siteprefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ip": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"servername"},
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"hashid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publicip": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
			},
			"disable_read": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
		},
	}
}
func createGslbservicegroup_gslbservicegroupmember_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbservicegroup_gslbservicegroupmember_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingIdSlice := make([]string, 0, 3)

	bindingIdSlice = append(bindingIdSlice, d.Get("servicegroupname").(string))

	// Second id component will either be servername or ip
	// ConflictsWith restriction ensures that only one of them is defined
	if v, ok := d.GetOk("servername"); ok {
		bindingIdSlice = append(bindingIdSlice, v.(string))
	}
	if v, ok := d.GetOk("ip"); ok {
		bindingIdSlice = append(bindingIdSlice, v.(string))
	}

	// Third component will be the port if defined
	if v, ok := d.GetOk("port"); ok {
		val := strconv.Itoa(v.(int))
		bindingIdSlice = append(bindingIdSlice, val)
	}

	bindingId := strings.Join(bindingIdSlice, ",")

	gslbservicegroup_gslbservicegroupmember_binding := gslb.Gslbservicegroupgslbservicegroupmemberbinding{
		State:   		  d.Get("state").(string),
		Publicport:       d.Get("publicport").(int),
		Port:             d.Get("port").(int),
		Weight:           d.Get("weight").(int),
		Servername:       d.Get("servername").(string),
		Siteprefix:		  d.Get("siteprefix").(string),
		Ip:               d.Get("ip").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
		Hashid:           d.Get("hashid").(int),
		Publicip:         d.Get("publicip").(string),
	}
// service.Gslbservicegroup_gslbservicegroupmember_binding  service.test123.Type()
	err := client.UpdateUnnamedResource(service.Gslbservicegroup_gslbservicegroupmember_binding.Type(), &gslbservicegroup_gslbservicegroupmember_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readGslbservicegroup_gslbservicegroupmember_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbservicegroup_gslbservicegroupmember_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readGslbservicegroup_gslbservicegroupmember_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbservicegroup_gslbservicegroupmember_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	// Skip reading when flag is set
	if d.Get("disable_read").(bool) {
		return nil
	}

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)
	servicegroupname := idSlice[0]

	// When ip is defined ADC will create a server with name equal to the ip address
	// So no matter what the user actually defined servername will always be valid search criterion
	servername := idSlice[1]

	// Port is optional
	// Need to check if it actually exists
	port := 0
	var err error
	if len(idSlice) == 3 {
		if port, err = strconv.Atoi(idSlice[2]); err != nil {
			return err
		}
	}

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbservicegroup_gslbservicegroupmember_binding state %v", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbservicegroup_gslbservicegroupmember_binding",
		ResourceName:             servicegroupname,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	log.Printf("[DEBUG] citrixadc-provider: dataArr: %v", dataArr)
	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservicegroup_gslbservicegroupmember_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if port != 0 && v["port"] != nil {
			portEqual := int(v["port"].(float64)) == port
			servernameEqual := v["servername"] == servername
			if servernameEqual && portEqual {
				foundIndex = i
				break
			}
		} else {
			log.Printf("[WARN] citrixadc-provider: port is zero")
			if v["servername"].(string) == servername {
				foundIndex = i
				break
			}
		}
	}
	log.Printf("[WARN] citrixadc-provider: foundIndex: %v", foundIndex)

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservicegroup_gslbservicegroupmember_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("state", data["state"])
	d.Set("publicport", data["publicport"])
	d.Set("port", data["port"])
	d.Set("weight", data["weight"])
	d.Set("servername", data["servername"])
	d.Set("siteprefix", data["siteprefix"])
	d.Set("ip", data["ip"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("hashid", data["hashid"])
	d.Set("publicip", data["publicip"])

	return nil

}

func deleteGslbservicegroup_gslbservicegroupmember_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbservicegroup_gslbservicegroupmember_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)
	servicegroupname := idSlice[0]
	servername := idSlice[1]
	port := 0
	var err error
	if len(idSlice) == 3 {
		if port, err = strconv.Atoi(idSlice[2]); err != nil {
			return err
		}
	}
	args := make([]string, 0, 3)

	// When ip is defined ADC will create a server with name equal to the ip address
	// So no matter what the user actually defined servername will always be valid search criterion
	args = append(args, fmt.Sprintf("servername:%s", servername))

	// Port is optional
	if port != 0 {
		args = append(args, fmt.Sprintf("port:%v", port))
	}

	err = client.DeleteResourceWithArgs("gslbservicegroup_gslbservicegroupmember_binding", servicegroupname, args)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
