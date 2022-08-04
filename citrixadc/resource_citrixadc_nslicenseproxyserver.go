package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"strings"
	"log"
)

func resourceCitrixAdcNslicenseproxyserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNslicenseproxyserverFunc,
		Read:          readNslicenseproxyserverFunc,
		Update:        updateNslicenseproxyserverFunc,
		Delete:        deleteNslicenseproxyserverFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"serverip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createNslicenseproxyserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNslicenseproxyserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var nslicenseproxyserverName string
	if v, ok := d.GetOk("serverip"); ok {
		nslicenseproxyserverName = v.(string)
	} else if v, ok := d.GetOk("servername"); ok {
		nslicenseproxyserverName = v.(string)
	}
	nslicenseproxyserver := ns.Nslicenseproxyserver{
		Port:       d.Get("port").(int),
		Serverip:   d.Get("serverip").(string),
		Servername: d.Get("servername").(string),
	}

	_, err := client.AddResource(service.Nslicenseproxyserver.Type(), nslicenseproxyserverName, &nslicenseproxyserver)
	if err != nil {
		return err
	}

	d.SetId(nslicenseproxyserverName)

	err = readNslicenseproxyserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nslicenseproxyserver but we can't read it ?? %s", nslicenseproxyserverName)
		return nil
	}
	return nil
}

func readNslicenseproxyserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNslicenseproxyserverFunc")
	client := meta.(*NetScalerNitroClient).client
	nslicenseproxyserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nslicenseproxyserver state %s", nslicenseproxyserverName)
	data, err := client.FindResource(service.Nslicenseproxyserver.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nslicenseproxyserver state %s", nslicenseproxyserverName)
		d.SetId("")
		return nil
	}
	d.Set("port", data["port"])
	d.Set("serverip", data["serverip"])
	d.Set("servername", data["servername"])

	return nil

}

func updateNslicenseproxyserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNslicenseproxyserverFunc")
	client := meta.(*NetScalerNitroClient).client
	nslicenseproxyserverName := d.Id()
	nslicenseproxyserver := ns.Nslicenseproxyserver{}
	
	if v, ok := d.GetOk("serverip"); ok {	
		nslicenseproxyserver.Serverip = v.(string)
	} else if v, ok := d.GetOk("servername"); ok {
		nslicenseproxyserver.Servername = v.(string)
	}
	hasChange := false
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for nslicenseproxyserver %s, starting update", nslicenseproxyserverName)
		nslicenseproxyserver.Port = d.Get("port").(int)
		hasChange = true
	}
	
	if hasChange {
		err := client.UpdateUnnamedResource(service.Nslicenseproxyserver.Type(), &nslicenseproxyserver)
		if err != nil {
			return fmt.Errorf("Error updating nslicenseproxyserver %s", nslicenseproxyserverName)
		}
	}
	return readNslicenseproxyserverFunc(d, meta)
}

func deleteNslicenseproxyserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNslicenseproxyserverFunc")
	client := meta.(*NetScalerNitroClient).client
	nslicenseproxyserverId := d.Id()
	idSlice := strings.SplitN(nslicenseproxyserverId, ",", 2)
	nslicenseproxyserverName := idSlice[0]
	
	err := client.DeleteResource(service.Nslicenseproxyserver.Type(), nslicenseproxyserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
