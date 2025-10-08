package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/citrix/adc-nitro-go/service"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcHanode() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createHanodeFunc,
		ReadContext:   readHanodeFunc,
		UpdateContext: updateHanodeFunc,
		DeleteContext: deleteHanodeFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"hanode_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"deadinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"failsafe": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"haprop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hastatus": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hasync": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hellointerval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"inc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxflips": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxfliptime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"syncstatusstrictmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syncvlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createHanodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createHanodeFunc")
	client := meta.(*NetScalerNitroClient).client
	hanodeName := strconv.Itoa(d.Get("hanode_id").(int))

	hanode := ha.Hanode{
		Id:                   d.Get("hanode_id").(int),
		Deadinterval:         d.Get("deadinterval").(int),
		Inc:                  d.Get("inc").(string),
		Failsafe:             d.Get("failsafe").(string),
		Haprop:               d.Get("haprop").(string),
		Hastatus:             d.Get("hastatus").(string),
		Hasync:               d.Get("hasync").(string),
		Hellointerval:        d.Get("hellointerval").(int),
		Maxflips:             d.Get("maxflips").(int),
		Maxfliptime:          d.Get("maxfliptime").(int),
		Syncstatusstrictmode: d.Get("syncstatusstrictmode").(string),
		Syncvlan:             d.Get("syncvlan").(int),
		Ipaddress:            d.Get("ipaddress").(string),
	}
	var err error
	if d.Get("hanode_id").(int) != 0 {
		_, err = client.AddResource(service.Hanode.Type(), hanodeName, &hanode)
	} else {
		err = client.UpdateUnnamedResource(service.Hanode.Type(), &hanode)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(hanodeName)

	return readHanodeFunc(ctx, d, meta)
}

func readHanodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	// We recieve "UP" as value from the NetScaler, if the user has given "ENABLED"
	// FIXME: Revisit once the API is fixed
	if data["hastatus"] == "UP" {
		d.Set("hastatus", "ENABLED")
	}
	setToInt("hanode_id", d, data["id"])
	setToInt("deadinterval", d, data["deadinterval"])
	d.Set("failsafe", data["failsafe"])
	d.Set("haprop", data["haprop"])
	// d.Set("hastatus", data["hastatus"]) // We recieve "UP" as value from the NetScaler, if the user has given "ENABLED"
	d.Set("hasync", data["hasync"])
	setToInt("hellointerval", d, data["hellointerval"])
	d.Set("inc", data["inc"])
	setToInt("maxflips", d, data["maxflips"])
	setToInt("maxfliptime", d, data["maxfliptime"])
	d.Set("syncstatusstrictmode", data["syncstatusstrictmode"])
	setToInt("syncvlan", d, data["syncvlan"])

	return nil

}

func updateHanodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("Error updating hanode %s", hanodeName)
		}
	}
	return readHanodeFunc(ctx, d, meta)
}

func deleteHanodeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteHanodeFunc")
	client := meta.(*NetScalerNitroClient).client
	hanodeName := d.Id()
	err := client.DeleteResource(service.Hanode.Type(), hanodeName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
