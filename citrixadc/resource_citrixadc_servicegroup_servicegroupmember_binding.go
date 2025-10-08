package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcServicegroup_servicegroupmember_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createServicegroup_servicegroupmember_bindingFunc,
		ReadContext:   readServicegroup_servicegroupmember_bindingFunc,
		DeleteContext: deleteServicegroup_servicegroupmember_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"customserverid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dbsttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"hashid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ip": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"servername"},
			},
			"nameserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"serverid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servername": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"ip"},
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"disable_read": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
		},
	}
}

func createServicegroup_servicegroupmember_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createServicegroup_servicegroupmember_bindingFunc")
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

	servicegroup_servicegroupmember_binding := basic.Servicegroupservicegroupmemberbinding{
		Customserverid:   d.Get("customserverid").(string),
		Dbsttl:           d.Get("dbsttl").(int),
		Hashid:           d.Get("hashid").(int),
		Ip:               d.Get("ip").(string),
		Nameserver:       d.Get("nameserver").(string),
		Port:             d.Get("port").(int),
		Serverid:         d.Get("serverid").(int),
		Servername:       d.Get("servername").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
		State:            d.Get("state").(string),
		Weight:           d.Get("weight").(int),
	}

	err := client.UpdateUnnamedResource(service.Servicegroup_servicegroupmember_binding.Type(), &servicegroup_servicegroupmember_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readServicegroup_servicegroupmember_bindingFunc(ctx, d, meta)
}

func readServicegroup_servicegroupmember_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readServicegroup_servicegroupmember_bindingFunc")
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
			return diag.FromErr(err)
		}
	}

	log.Printf("[DEBUG] citrixadc-provider: Reading servicegroup_servicegroupmember_binding state %v", bindingId)

	findParams := service.FindParams{
		ResourceType:             "servicegroup_servicegroupmember_binding",
		ResourceName:             servicegroupname,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	log.Printf("[DEBUG] citrixadc-provider: dataArr: %v", dataArr)
	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return diag.FromErr(err)
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing servicegroup_servicegroupmember_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing servicegroup_servicegroupmember_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("customserverid", data["customserverid"])
	setToInt("dbsttl", d, data["dbsttl"])
	setToInt("hashid", d, data["hashid"])
	d.Set("ip", data["ip"])
	d.Set("nameserver", data["nameserver"])
	setToInt("port", d, data["port"])
	setToInt("serverid", d, data["serverid"])
	d.Set("servername", data["servername"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("state", data["state"])
	setToInt("weight", d, data["weight"])

	return nil

}

func deleteServicegroup_servicegroupmember_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteServicegroup_servicegroupmember_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)
	servicegroupname := idSlice[0]
	servername := idSlice[1]
	port := 0
	var err error
	if len(idSlice) == 3 {
		if port, err = strconv.Atoi(idSlice[2]); err != nil {
			return diag.FromErr(err)
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

	err = client.DeleteResourceWithArgs("servicegroup_servicegroupmember_binding", servicegroupname, args)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
