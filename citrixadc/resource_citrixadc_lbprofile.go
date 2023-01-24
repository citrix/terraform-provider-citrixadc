package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLbprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbprofileFunc,
		Read:          readLbprofileFunc,
		Update:        updateLbprofileFunc,
		Delete:        deleteLbprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"lbprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dbslb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"processlocal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httponlycookieflag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookiepassphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usesecuredpersistencecookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useencryptedpersistencecookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"literaladccookieattribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"computedadccookieattribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storemqttclientidandusername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lbhashalgorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lbhashfingers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLbprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lbprofileName := d.Get("lbprofilename").(string)

	lbprofile := lb.Lbprofile{
		Lbprofilename:                 lbprofileName,
		Dbslb:                         d.Get("dbslb").(string),
		Processlocal:                  d.Get("processlocal").(string),
		Httponlycookieflag:            d.Get("httponlycookieflag").(string),
		Cookiepassphrase:              d.Get("cookiepassphrase").(string),
		Usesecuredpersistencecookie:   d.Get("usesecuredpersistencecookie").(string),
		Useencryptedpersistencecookie: d.Get("useencryptedpersistencecookie").(string),
		Literaladccookieattribute:     d.Get("literaladccookieattribute").(string),
		Computedadccookieattribute:    d.Get("computedadccookieattribute").(string),
		Storemqttclientidandusername:  d.Get("storemqttclientidandusername").(string),
		Lbhashalgorithm:               d.Get("lbhashalgorithm").(string),
		Lbhashfingers:                 d.Get("lbhashfingers").(int),
	}

	_, err := client.AddResource("lbprofile", lbprofileName, &lbprofile)
	if err != nil {
		return err
	}

	d.SetId(lbprofileName)

	err = readLbprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbprofile but we can't read it ?? %s", lbprofileName)
		return nil
	}
	return nil
}

func readLbprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lbprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbprofile state %s", lbprofileName)
	data, err := client.FindResource("lbprofile", lbprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbprofile state %s", lbprofileName)
		d.SetId("")
		return nil
	}
	d.Set("lbprofilename", data["lbprofilename"])
	d.Set("dbslb", data["dbslb"])
	d.Set("processlocal", data["processlocal"])
	d.Set("httponlycookieflag", data["httponlycookieflag"])
	d.Set("cookiepassphrase", data["cookiepassphrase"])
	d.Set("usesecuredpersistencecookie", data["usesecuredpersistencecookie"])
	d.Set("useencryptedpersistencecookie", data["useencryptedpersistencecookie"])
	d.Set("literaladccookieattribute", data["literaladccookieattribute"])
	d.Set("computedadccookieattribute", data["computedadccookieattribute"])
	d.Set("storemqttclientidandusername", data["storemqttclientidandusername"])
	d.Set("lbhashalgorithm", data["lbhashalgorithm"])
	d.Set("lbhashfingers", data["lbhashfingers"])

	return nil

}

func updateLbprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lbprofileName := d.Get("lbprofilename").(string)

	lbprofile := lb.Lbprofile{
		Lbprofilename: d.Get("lbprofilename").(string),
	}

	hasChange := false
	if d.HasChange("lbprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Lbprofilename has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Lbprofilename = d.Get("lbprofilename").(string)
		hasChange = true
	}
	if d.HasChange("dbslb") {
		log.Printf("[DEBUG]  citrixadc-provider: Dbslb has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Dbslb = d.Get("dbslb").(string)
		hasChange = true
	}
	if d.HasChange("processlocal") {
		log.Printf("[DEBUG]  citrixadc-provider: Processlocal has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Processlocal = d.Get("processlocal").(string)
		hasChange = true
	}
	if d.HasChange("httponlycookieflag") {
		log.Printf("[DEBUG]  citrixadc-provider: Httponlycookieflag has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Httponlycookieflag = d.Get("httponlycookieflag").(string)
		hasChange = true
	}
	if d.HasChange("cookiepassphrase") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookiepassphrase has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Cookiepassphrase = d.Get("cookiepassphrase").(string)
		hasChange = true
	}
	if d.HasChange("usesecuredpersistencecookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Usesecuredpersistencecookie has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Usesecuredpersistencecookie = d.Get("usesecuredpersistencecookie").(string)
		hasChange = true
	}
	if d.HasChange("useencryptedpersistencecookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Useencryptedpersistencecookie has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Useencryptedpersistencecookie = d.Get("useencryptedpersistencecookie").(string)
		hasChange = true
	}
	if d.HasChange("literaladccookieattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Literaladccookieattribute has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Literaladccookieattribute = d.Get("literaladccookieattribute").(string)
		hasChange = true
	}
	if d.HasChange("computedadccookieattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Computedadccookieattribute has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Computedadccookieattribute = d.Get("computedadccookieattribute").(string)
		hasChange = true
	}
	if d.HasChange("storemqttclientidandusername") {
		log.Printf("[DEBUG]  citrixadc-provider: Storemqttclientidandusername has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Storemqttclientidandusername = d.Get("storemqttclientidandusername").(string)
		hasChange = true
	}
	if d.HasChange("lbhashalgorithm") {
		log.Printf("[DEBUG]  citrixadc-provider: Lbhashalgorithm has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Lbhashalgorithm = d.Get("lbhashalgorithm").(string)
		hasChange = true
	}
	if d.HasChange("lbhashfingers") {
		log.Printf("[DEBUG]  citrixadc-provider: Lbhashfingers has changed for lbprofile %s, starting update", lbprofileName)
		lbprofile.Lbhashfingers = d.Get("lbhashfingers").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("lbprofile", lbprofileName, &lbprofile)
		if err != nil {
			return fmt.Errorf("Error updating lbprofile %s", lbprofileName)
		}
	}
	return readLbprofileFunc(d, meta)
}

func deleteLbprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lbprofileName := d.Id()
	err := client.DeleteResource("lbprofile", lbprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
