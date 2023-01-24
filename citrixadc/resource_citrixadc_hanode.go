package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"strconv"
	"log"
)

func resourceCitrixAdcHanode() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createHanodeFunc,
		Read:          readHanodeFunc,
		Update:        updateHanodeFunc,
		Delete:        deleteHanodeFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"hanode_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"deadinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"failsafe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"haprop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hastatus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hasync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hellointerval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"inc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxflips": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxfliptime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"syncstatusstrictmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syncvlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createHanodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createHanodeFunc")
	client := meta.(*NetScalerNitroClient).client
	hanodeName := strconv.Itoa(d.Get("hanode_id").(int))
	
	hanode := ha.Hanode{
        Id:                   d.Get("hanode_id").(int),
        Deadinterval:         d.Get("deadinterval").(int),
		Inc:				  d.Get("inc").(string),
        Failsafe:             d.Get("failsafe").(string),
        Haprop:               d.Get("haprop").(string),
        Hastatus:             d.Get("hastatus").(string),
        Hasync:               d.Get("hasync").(string),
        Hellointerval:        d.Get("hellointerval").(int),
        Maxflips:             d.Get("maxflips").(int),
        Maxfliptime:          d.Get("maxfliptime").(int),
        Syncstatusstrictmode: d.Get("syncstatusstrictmode").(string),
        Syncvlan:             d.Get("syncvlan").(int),
		Ipaddress: 			  d.Get("ipaddress").(string),
    }
	var err error 
	if d.Get("hanode_id").(int) != 0 {
	_, err = client.AddResource(service.Hanode.Type(), hanodeName, &hanode)
	} else {
		err = client.UpdateUnnamedResource(service.Hanode.Type(), &hanode)
	}
	if err != nil {
		return err
	}

	d.SetId(hanodeName)


	err = readHanodeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this hanode but we can't read it ?? %s", hanodeName)
		return nil
	}
	return nil
}

func readHanodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readHanodeFunc")
	client := meta.(*NetScalerNitroClient).client
	hanodeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading hanode state %s", hanodeName)
	data, err := client.FindResource(service.Hanode.Type(), hanodeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing hanode state %s", hanodeName)
		d.SetId("")
		return nil
	}
	d.Set("hanode_id", data["id"])
	d.Set("deadinterval", data["deadinterval"])
	d.Set("failsafe", data["failsafe"])
	d.Set("haprop", data["haprop"])
	d.Set("hastatus", data["hastatus"])
	d.Set("hasync", data["hasync"])
	d.Set("hellointerval", data["hellointerval"])
	d.Set("inc", data["inc"])
	d.Set("maxflips", data["maxflips"])
	d.Set("maxfliptime", data["maxfliptime"])
	d.Set("syncstatusstrictmode", data["syncstatusstrictmode"])
	d.Set("syncvlan", data["syncvlan"])

	return nil

}

func updateHanodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateHanodeFunc")
	client := meta.(*NetScalerNitroClient).client
	hanodeName := d.Id()

	hanode := ha.Hanode{
		Id: d.Get("hanode_id").(int),
	}
	hasChange := false
	if d.HasChange("deadinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Deadinterval has changed for hanode %s, starting update", hanodeName)
		hanode.Deadinterval = d.Get("deadinterval").(int)
		hasChange = true
	}
	if d.HasChange("failsafe") {
		log.Printf("[DEBUG]  citrixadc-provider: Failsafe has changed for hanode %s, starting update", hanodeName)
		hanode.Failsafe = d.Get("failsafe").(string)
		hasChange = true
	}
	if d.HasChange("haprop") {
		log.Printf("[DEBUG]  citrixadc-provider: Haprop has changed for hanode %s, starting update", hanodeName)
		hanode.Haprop = d.Get("haprop").(string)
		hasChange = true
	}
	if d.HasChange("hastatus") {
		log.Printf("[DEBUG]  citrixadc-provider: Hastatus has changed for hanode %s, starting update", hanodeName)
		hanode.Hastatus = d.Get("hastatus").(string)
		hasChange = true
	}
	if d.HasChange("hasync") {
		log.Printf("[DEBUG]  citrixadc-provider: Hasync has changed for hanode %s, starting update", hanodeName)
		hanode.Hasync = d.Get("hasync").(string)
		hasChange = true
	}
	if d.HasChange("hellointerval") {
		log.Printf("[DEBUG]  citrixadc-provider: Hellointerval has changed for hanode %s, starting update", hanodeName)
		hanode.Hellointerval = d.Get("hellointerval").(int)
		hasChange = true
	}
	if d.HasChange("inc") {
		log.Printf("[DEBUG]  citrixadc-provider: Inc has changed for hanode %s, starting update", hanodeName)
		hanode.Inc = d.Get("inc").(string)
		hasChange = true
	}
	if d.HasChange("maxflips") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxflips has changed for hanode %s, starting update", hanodeName)
		hanode.Maxflips = d.Get("maxflips").(int)
		hasChange = true
	}
	if d.HasChange("maxfliptime") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxfliptime has changed for hanode %s, starting update", hanodeName)
		hanode.Maxfliptime = d.Get("maxfliptime").(int)
		hasChange = true
	}
	if d.HasChange("syncstatusstrictmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Syncstatusstrictmode has changed for hanode %s, starting update", hanodeName)
		hanode.Syncstatusstrictmode = d.Get("syncstatusstrictmode").(string)
		hasChange = true
	}
	if d.HasChange("syncvlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Syncvlan has changed for hanode %s, starting update", hanodeName)
		hanode.Syncvlan = d.Get("syncvlan").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Hanode.Type(), &hanode)
		if err != nil {
			return fmt.Errorf("Error updating hanode %s", hanodeName)
		}
	}
	return readHanodeFunc(d, meta)
}

func deleteHanodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteHanodeFunc")
	client := meta.(*NetScalerNitroClient).client
	hanodeName := d.Id()
	err := client.DeleteResource(service.Hanode.Type(), hanodeName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
