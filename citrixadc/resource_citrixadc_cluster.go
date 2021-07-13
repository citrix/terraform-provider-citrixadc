package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/resource/config/router"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"bytes"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

func resourceCitrixAdcCluster() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusterFunc,
		Read:          readClusterFunc,
		Update:        updateClusterFunc,
		Delete:        deleteClusterFunc,
		Schema: map[string]*schema.Schema{
			"backplanebasedview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"clip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"deadinterval": &schema.Schema{
				Type:     schema.TypeInt,
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
			"nodegroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
			},
			"preemption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"processlocal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quorumtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retainconnectionsoncluster": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// How long to wait before first poll
			"bootstrap_poll_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "60s",
			},
			// Interval between polls
			"bootstrap_poll_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "60s",
			},
			// Timeout for each individual HTTP poll
			"bootstrap_poll_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s",
			},
			// Timeout for the whole operation
			"bootstrap_total_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10m",
			},
			// Delay before first poll
			"clip_migration_poll_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s",
			},
			// Interval between polls
			"clip_migration_poll_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "30s",
			},
			// Timeout for each individual poll HTTP request
			"clip_migration_poll_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s",
			},
			// Timeout for the whole operation
			"clip_migration_total_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10m",
			},
			// Timeouts for the node add operation
			"node_add_poll_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s",
			},
			"node_add_poll_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "30s",
			},
			"node_add_total_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10m",
			},
			"clusternodegroup": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				Set:      clusternodegroupMappingHash,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sticky": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"strict": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"clusternode": {
				Type:     schema.TypeSet,
				Required: true,
				Computed: false,
				Set:      clusternodeMappingHash,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backplane": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"clearnodegroupconfig": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ipaddress": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"nodegroup": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nodeid": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnelmode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						// Overrides for the particular node
						// Needed in bootstraping and adding a new node
						"endpoint": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Computed: false,
						},
						"username": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  false,
							Sensitive: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  false,
							Sensitive: true,
						},
						"insecure_skip_verify": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Ignore validity of endpoint TLS certificate if true",
							Default:     false,
						},
						// Flags for adding node SNIP to CLIP before joining
						"snip_netmask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"snip_ipaddress": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"addsnip": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						// Optional VTYSH commands
						"vtysh_enable": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"vtysh": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func createClusterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusterFunc")
	var err error
	clid := strconv.Itoa(d.Get("clid").(int))

	if err = bootstrapCluster(d, meta); err != nil {
		return err
	}

	d.SetId(clid)

	err = readClusterFunc(d, meta)
	if err != nil {
		return err
	}
	return nil
}

func readClusterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusterFunc")
	client := meta.(*NetScalerNitroClient).client
	clusterId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cluster state %s", clusterId)
	datalist, err := client.FindAllResources(service.Clusterinstance.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cluster state %s", clusterId)
		d.SetId("")
		return err
	}

	if len(datalist) == 0 {
		return fmt.Errorf("[ERROR] could not retrieve cluster instance information.")
	}

	data := datalist[0]
	clid, err := strconv.Atoi(data["clid"].(string))
	if err != nil {
		return err
	}
	d.Set("clid", clid)
	log.Printf("clid %v", clid)

	deadinterval := int(data["deadinterval"].(float64))
	d.Set("deadinterval", deadinterval)
	log.Printf("deadinterval %v", deadinterval)

	hellointerval := int(data["hellointerval"].(float64))
	d.Set("hellointerval", hellointerval)

	log.Printf("hellointerval %v", hellointerval)

	d.Set("backplanebasedview", data["backplanebasedview"])
	d.Set("inc", data["inc"])
	d.Set("nodegroup", data["nodegroup"])
	d.Set("preemption", data["preemption"])
	d.Set("processlocal", data["processlocal"])
	d.Set("quorumtype", data["quorumtype"])
	d.Set("retainconnectionsoncluster", data["retainconnectionsoncluster"])

	err = readClusterNodes(d, meta)
	if err != nil {
		return err
	}
	if isClusterModeL3(d) {
		err = readClusterNodegroups(d, meta)
		if err != nil {
			return err
		}
	}

	return nil

}

func updateClusterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateClusterFunc")
	client := meta.(*NetScalerNitroClient).client

	clid := strconv.Itoa(d.Get("clid").(int))

	clusterinstance := cluster.Clusterinstance{
		Clid: d.Get("clid").(int),
	}
	hasChange := false
	clusterNodegroupChanged := false
	if d.HasChange("backplanebasedview") {
		log.Printf("[DEBUG]  citrixadc-provider: Backplanebasedview has changed for clusterinstance %s, starting update", clid)
		clusterinstance.Backplanebasedview = d.Get("backplanebasedview").(string)
		hasChange = true
	}
	if d.HasChange("clid") {
		log.Printf("[DEBUG]  citrixadc-provider: Clid has changed for clusterinstance %s, starting update", clid)
		clusterinstance.Clid = d.Get("clid").(int)
		hasChange = true
	}
	if d.HasChange("deadinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Deadinterval has changed for clusterinstance %s, starting update", clid)
		clusterinstance.Deadinterval = d.Get("deadinterval").(int)
		hasChange = true
	}
	if d.HasChange("hellointerval") {
		log.Printf("[DEBUG]  citrixadc-provider: Hellointerval has changed for clusterinstance %s, starting update", clid)
		clusterinstance.Hellointerval = d.Get("hellointerval").(int)
		hasChange = true
	}
	if d.HasChange("inc") {
		log.Printf("[DEBUG]  citrixadc-provider: Inc has changed for clusterinstance %s, starting update", clid)
		clusterinstance.Inc = d.Get("inc").(string)
		hasChange = true
	}
	if d.HasChange("nodegroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodegroup has changed for clusterinstance %s, starting update", clid)
		clusterinstance.Nodegroup = d.Get("nodegroup").(string)
		hasChange = true
	}
	if d.HasChange("preemption") {
		log.Printf("[DEBUG]  citrixadc-provider: Preemption has changed for clusterinstance %s, starting update", clid)
		clusterinstance.Preemption = d.Get("preemption").(string)
		hasChange = true
	}
	if d.HasChange("processlocal") {
		log.Printf("[DEBUG]  citrixadc-provider: Processlocal has changed for clusterinstance %s, starting update", clid)
		clusterinstance.Processlocal = d.Get("processlocal").(string)
		hasChange = true
	}
	if d.HasChange("quorumtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Quorumtype has changed for clusterinstance %s, starting update", clid)
		clusterinstance.Quorumtype = d.Get("quorumtype").(string)
		hasChange = true
	}
	if d.HasChange("retainconnectionsoncluster") {
		log.Printf("[DEBUG]  citrixadc-provider: Retainconnectionsoncluster has changed for clusterinstance %s, starting update", clid)
		clusterinstance.Retainconnectionsoncluster = d.Get("retainconnectionsoncluster").(string)
		hasChange = true
	}

	if d.HasChange("clusternodegroup") {
		log.Printf("[DEBUG]  citrixadc-provider: clusternodegroup has changed for clusterinstance %s, starting update", clid)
		clusterNodegroupChanged = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Clusterinstance.Type(), clid, &clusterinstance)
		if err != nil {
			return fmt.Errorf("Error updating clusterinstance %s. %s", clid, err.Error())
		}
	}

	// Add and update nodgroups before nodes
	if isClusterModeL3(d) && clusterNodegroupChanged {
		if err := addClusterNodegroups(d, meta); err != nil {
			return err
		}
		if err := updateClusterNodegroups(d, meta); err != nil {
			return err
		}
	}

	if err := updateClusterNodes(d, meta); err != nil {
		return err
	}

	// Delete nodegroups after nodes
	if isClusterModeL3(d) && clusterNodegroupChanged {
		if err := deleteClusterNodegroups(d, meta); err != nil {
			return err
		}
	}

	return readClusterFunc(d, meta)
}

func deleteClusterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusterFunc")
	//client := meta.(*NetScalerNitroClient).client

	nodeids := getSortedClusternodeIds(d)
	// Delete nodes in sequence excluding CCO
	for _, nodeid := range nodeids[1:] {
		nodeData := getClusterNodeByid(d, nodeid)
		err := deleteSingleClusterNode(d, meta, nodeData, true)
		if err != nil {
			return err
		}
	}

	// Delete CCO last
	nodeData := getClusterNodeByid(d, nodeids[0])
	// Don't wait for CLIP migration on deletion of last node
	err := deleteSingleClusterNode(d, meta, nodeData, false)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func clusternodegroupMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In clusternodegroupMappingHash")
	var buf bytes.Buffer

	m := v.(map[string]interface{})

	if d, ok := m["name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["priority"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}

	if d, ok := m["state"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["sticky"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["strict"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	return hashcode.String(buf.String())
}

func clusternodeMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In clusternodeMappingHash")
	var buf bytes.Buffer

	m := v.(map[string]interface{})

	if d, ok := m["backplane"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["clearnodegroupconfig"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["delay"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}

	if d, ok := m["ipaddress"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["nodegroup"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["nodeid"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}

	if d, ok := m["priority"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}

	if d, ok := m["state"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["tunnelmode"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	return hashcode.String(buf.String())
}

type nodePriority struct {
	nodeid   int
	priority int
}

type nodePriorities []nodePriority

func (n nodePriorities) Len() int {
	return len(n)
}

func (n nodePriorities) Less(i, j int) bool {
	if n[i].priority == n[j].priority {
		return n[i].nodeid < n[j].nodeid
	} else {
		return n[i].priority < n[j].priority
	}
}

func (n nodePriorities) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func getSortedClusternodeIds(d *schema.ResourceData) []int {
	log.Printf("[DEBUG]  citrixadc-provider: In getSortedClusternodeIds")
	// Sort cluster nodes according to priority + node id
	// First node is the CCO
	clusterNodes := d.Get("clusternode").(*schema.Set)
	nodes := make(nodePriorities, 0, clusterNodes.Len())

	for _, v := range clusterNodes.List() {
		val := v.(map[string]interface{})
		nodeid := val["nodeid"].(int)
		priority := val["priority"].(int)
		node := nodePriority{
			nodeid:   nodeid,
			priority: priority,
		}

		nodes = append(nodes, node)
	}
	sort.Sort(nodes)
	nodeids := make([]int, 0, clusterNodes.Len())
	for _, val := range nodes {
		nodeids = append(nodeids, val.nodeid)
	}
	return nodeids
}

func getClusterNodeByid(d *schema.ResourceData, id int) map[string]interface{} {
	log.Printf("[DEBUG]  citrixadc-provider: In getClusterNodeByid")
	for _, item := range d.Get("clusternode").(*schema.Set).List() {
		val := item.(map[string]interface{})
		if val["nodeid"].(int) == id {
			return val
		}
	}
	return nil
}

func getClusterNodegroupByName(d *schema.ResourceData, nodegroupName string) map[string]interface{} {
	log.Printf("[DEBUG]  citrixadc-provider: In getClusterNodegroupByName")
	for _, item := range d.Get("clusternodegroup").(*schema.Set).List() {
		val := item.(map[string]interface{})
		if val["name"].(string) == nodegroupName {
			return val
		}
	}
	return nil
}

func clusternodegroupExistsInCluster(client *service.NitroClient, nodegroupName string) (bool, error) {
	log.Printf("[DEBUG]  citrixadc-provider: In clusternodegroupExistsInCluster")
	findParams := service.FindParams{
		ResourceType:             "clusternodegroup",
		ResourceName:             nodegroupName,
		ResourceMissingErrorCode: 258,
	}
	data, err := client.FindResourceArrayWithParams(findParams)
	log.Printf("[DEBUG]  citrixadc-provider: data %v", data)

	if err != nil {
		return false, err
	}

	if len(data) == 0 {
		return false, nil
	} else if len(data) == 1 {
		return true, nil
	} else {
		return false, fmt.Errorf("Got multiple node groups existing for name %s. ( %v )", nodegroupName, data)
	}

}

func getClusternodegroupFromCluster(client *service.NitroClient, nodegroupName string) (map[string]interface{}, error) {
	log.Printf("[DEBUG]  citrixadc-provider: In getClusternodegroupFromCluster")
	findParams := service.FindParams{
		ResourceType:             "clusternodegroup",
		ResourceName:             nodegroupName,
		ResourceMissingErrorCode: 258,
	}
	data, err := client.FindResourceArrayWithParams(findParams)
	log.Printf("[DEBUG]  citrixadc-provider: data %v", data)

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	} else if len(data) == 1 {
		return data[0], nil
	} else {
		return nil, fmt.Errorf("Got multiple node groups existing for name %s. ( %v )", nodegroupName, data)
	}
}

func isClusterModeL3(d *schema.ResourceData) bool {
	log.Printf("[DEBUG]  citrixadc-provider: In isClusterModeL3")
	inc := d.Get("inc").(string)
	if _, ok := d.GetOk("clusternodegroup"); !ok {
		return false
	}
	nodegroupList := d.Get("clusternodegroup").(*schema.Set).List()
	if inc == "ENABLED" && len(nodegroupList) > 0 {
		return true
	} else {
		return false
	}
}

func bootstrapCluster(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In bootstrapCluster")
	var err error

	if err = createFirstClusterNode(d, meta); err != nil {
		return err
	}

	if isClusterModeL3(d) {
		err = addClusterNodegroups(d, meta)
		if err != nil {
			return err
		}
	}

	// Join rest of nodes to the cluster
	nodeids := getSortedClusternodeIds(d)
	for _, nodeid := range nodeids[1:] {
		nodeData := getClusterNodeByid(d, nodeid)
		err = addSingleClusterNode(d, meta, nodeData)
		if err != nil {
			return err
		}
	}

	return nil
}

func createFirstClusterNode(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createFirstClusterNode")

	nodeids := getSortedClusternodeIds(d)
	log.Printf("[DEBUG]  citrixadc-provider: sorted node ids %v", nodeids)
	firstNode := getClusterNodeByid(d, nodeids[0])

	// The provider endpoint is persumed to be the CLIP
	// We need to instantiate a separate go-nitro client
	// with the credentials of the first node

	nodeClient, err := instantiateNodeClient(d, meta, firstNode)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG]  citrixadc-provider: first node client %v", nodeClient)

	clid := d.Get("clid").(int)
	clusterId := strconv.Itoa(clid)
	if err != nil {
		return err
	}

	clusterinstance := cluster.Clusterinstance{
		Backplanebasedview:         d.Get("backplanebasedview").(string),
		Clid:                       d.Get("clid").(int),
		Deadinterval:               d.Get("deadinterval").(int),
		Hellointerval:              d.Get("hellointerval").(int),
		Inc:                        d.Get("inc").(string),
		Nodegroup:                  d.Get("nodegroup").(string),
		Preemption:                 d.Get("preemption").(string),
		Processlocal:               d.Get("processlocal").(string),
		Quorumtype:                 d.Get("quorumtype").(string),
		Retainconnectionsoncluster: d.Get("retainconnectionsoncluster").(string),
	}
	_, err = nodeClient.AddResource("clusterinstance", clusterId, &clusterinstance)
	if err != nil {
		return err
	}

	// In L3 we need to create the nodegroup prior to adding the node
	if isClusterModeL3(d) {
		nodegroupName := firstNode["nodegroup"].(string)
		nodegroupData := getClusterNodegroupByName(d, nodegroupName)
		if nodegroupData == nil {
			return fmt.Errorf("Cannot find node group %s in configuration", nodegroupName)
		}

		clusternodegroup := cluster.Clusternodegroup{
			Name:     nodegroupData["name"].(string),
			Priority: nodegroupData["priority"].(int),
			State:    nodegroupData["state"].(string),
			Sticky:   nodegroupData["sticky"].(string),
			Strict:   nodegroupData["strict"].(string),
		}

		_, err := nodeClient.AddResource("clusternodegroup", clusternodegroup.Name, &clusternodegroup)
		if err != nil {
			return err
		}
	}

	// Add first cluster node

	clusternode := cluster.Clusternode{
		Backplane:            firstNode["backplane"].(string),
		Clearnodegroupconfig: firstNode["clearnodegroupconfig"].(string),
		Delay:                firstNode["delay"].(int),
		Ipaddress:            firstNode["ipaddress"].(string),
		Nodegroup:            firstNode["nodegroup"].(string),
		Nodeid:               firstNode["nodeid"].(int),
		Priority:             firstNode["priority"].(int),
		State:                firstNode["state"].(string),
		Tunnelmode:           firstNode["tunnelmode"].(string),
	}

	log.Printf("[DEBUG]  citrixadc-provider: Nodeid %v", clusternode.Nodeid)
	_, err = nodeClient.AddResource("clusternode", strconv.FormatUint(uint64(clusternode.Nodeid), 10), &clusternode)
	if err != nil {
		return err
	}

	// Add CLIP to first node
	ipaddress := d.Get("clip").(string)
	nsip := ns.Nsip{
		Ipaddress: ipaddress,
		Netmask:   "255.255.255.255",
		Type:      "CLIP",
	}

	_, err = nodeClient.AddResource(service.Nsip.Type(), ipaddress, &nsip)
	if err != nil {
		return err
	}

	// Enable cluster instance on first node
	clusterinstanceEnabler := cluster.Clusterinstance{
		Clid: d.Get("clid").(int),
	}
	err = nodeClient.ActOnResource("clusterinstance", &clusterinstanceEnabler, "enable")
	if err != nil {
		return err
	}

	// Save config
	nodeClient.SaveConfig()

	// Reboot Instance
	log.Printf("[DEBUG]  citrixadc-provider: Rebooting first node")
	reboot := ns.Reboot{
		Warm: true,
	}
	if err := nodeClient.ActOnResource("reboot", &reboot, ""); err != nil {
		return err
	}

	// Poll CLIP for bootstrap operation completion

	var bootstrap_poll_interval time.Duration
	if bootstrap_poll_interval, err = time.ParseDuration(d.Get("bootstrap_poll_interval").(string)); err != nil {
		return err
	}

	var bootstrap_poll_delay time.Duration
	if bootstrap_poll_delay, err = time.ParseDuration(d.Get("bootstrap_poll_delay").(string)); err != nil {
		return err
	}

	var bootstrap_total_timeout time.Duration
	if bootstrap_total_timeout, err = time.ParseDuration(d.Get("bootstrap_total_timeout").(string)); err != nil {
		return err
	}

	var bootstrap_poll_timeout time.Duration
	if bootstrap_poll_timeout, err = time.ParseDuration(d.Get("bootstrap_poll_timeout").(string)); err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"cluster_unreachable"},
		Target:       []string{"cluster_reachable"},
		Refresh:      resourceClusterPoll(d, meta, bootstrap_poll_timeout),
		Timeout:      bootstrap_total_timeout,
		PollInterval: bootstrap_poll_interval,
		Delay:        bootstrap_poll_delay,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}

	// Verify that the first node is actually part of the cluster
	client := meta.(*NetScalerNitroClient).client
	data, err := client.FindAllResources("clusternode")
	if err != nil {
		return err
	}

	fetchedIpaddress := data[0]["ipaddress"]
	configIpaddress := firstNode["ipaddress"]
	if fetchedIpaddress != configIpaddress {
		return fmt.Errorf("Fetched first node address differs from configuration. Fetched: %s Configuration: %s", fetchedIpaddress, configIpaddress)
	}

	return nil
}

func readClusterNodegroups(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readClusterNodegroups")

	client := meta.(*NetScalerNitroClient).client
	clusternodegroups, err := client.FindAllResources("clusternodegroup")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG]  citrixadc-provider: fetched clusternodegroups %v", clusternodegroups)

	processedClusterNodegroups := make([]interface{}, 0, len(clusternodegroups))
	for _, clusternodegroup := range clusternodegroups {
		node := make(map[string]interface{})

		if _, ok := clusternodegroup["name"]; ok {
			node["name"] = clusternodegroup["name"].(string)
		}

		if _, ok := clusternodegroup["priority"]; ok {
			node["priority"] = clusternodegroup["priority"].(int)
		}

		if _, ok := clusternodegroup["state"]; ok {
			node["state"] = clusternodegroup["state"].(string)
		}

		if _, ok := clusternodegroup["sticky"]; ok {
			node["sticky"] = clusternodegroup["sticky"].(string)
		}

		if _, ok := clusternodegroup["strict"]; ok {
			node["strict"] = clusternodegroup["strict"].(string)
		}

		if node["name"] == "DEFAULT_NG" {
			continue
		} else {
			processedClusterNodegroups = append(processedClusterNodegroups, node)
		}
	}

	updatedSet := schema.NewSet(clusternodegroupMappingHash, processedClusterNodegroups)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("clusternodegroup", updatedSet); err != nil {
		return err
	}

	return nil
}

func addClusterNodegroups(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addClusterNodegroups")
	client := meta.(*NetScalerNitroClient).client

	nodegroupList := d.Get("clusternodegroup").(*schema.Set).List()
	for _, nodegroup := range nodegroupList {
		nodegroupData := nodegroup.(map[string]interface{})
		name := nodegroupData["name"].(string)

		exists, err := clusternodegroupExistsInCluster(client, name)

		if err != nil {
			return err
		}

		if !exists {
			err = addSingleClusterNodegroup(d, meta, nodegroupData)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func updateClusterNodegroups(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateClusterNodegroups")
	client := meta.(*NetScalerNitroClient).client

	// Do updates
	nodegroupList := d.Get("clusternodegroup").(*schema.Set).List()
	for _, nodegroup := range nodegroupList {
		nodegroupData := nodegroup.(map[string]interface{})
		name := nodegroupData["name"].(string)

		exists, err := clusternodegroupExistsInCluster(client, name)

		if err != nil {
			return err
		}

		if exists {
			fetchedNodegroupData, err := getClusternodegroupFromCluster(client, nodegroupData["name"].(string))
			if err != nil {
				return err
			}
			needsUpdate := false
			// Check if any attribute needs update
			for k, configValue := range nodegroupData {
				fetchedValue, _ := fetchedNodegroupData[k]
				if k == "priority" {
					// Priority is numeric
					if configValue != 0 && fetchedValue != configValue {
						log.Printf("[DEBUG]  citrixadc-provider: needs update k:%v config:%v fetched:%v", k, configValue, fetchedValue)
						needsUpdate = true
					}
				} else {
					// Other keys are strings
					if configValue != "" && fetchedValue != configValue {
						log.Printf("[DEBUG]  citrixadc-provider: needs update k:%v config:%v fetched:%v", k, configValue, fetchedValue)
						needsUpdate = true
					}
				}
			}

			log.Printf("[DEBUG]  citrixadc-provider: nodegroup:%v needsUpdate:%v", nodegroupData["name"], needsUpdate)
			if needsUpdate {
				err = updateSingleClusterNodegroup(d, meta, nodegroupData)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func deleteClusterNodegroups(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusterNodegroups")

	o, n := d.GetChange("clusternodegroup")
	oldSet := o.(*schema.Set)
	newSet := n.(*schema.Set)

	log.Printf("[DEBUG]  citrixadc-provider: old nodegroup set: %v", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: new nodegroup set: %v", newSet)

	nameInSet := func(s *schema.Set, name string) bool {
		for _, v := range s.List() {
			if v.(map[string]interface{})["name"].(string) == name {
				return true
			}
		}
		return false
	}

	// Do deletes
	for _, nodegroup := range oldSet.List() {
		nodegroupData := nodegroup.(map[string]interface{})
		name := nodegroupData["name"].(string)

		// Ignore DEFAULT_NG node group
		if name == "DEFAULT_NG" {
			continue
		}

		if !nameInSet(newSet, name) {
			deleteSingleClusterNodegroup(d, meta, nodegroupData)
		}
	}
	return nil
}

func updateSingleClusterNodegroup(d *schema.ResourceData, meta interface{}, nodegroupData map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSingleClusterNodegroup")
	client := meta.(*NetScalerNitroClient).client

	log.Printf("[DEBUG]  citrixadc-provider: Updating nodegroup %v", nodegroupData)

	clusternodegroup := cluster.Clusternodegroup{
		Name:     nodegroupData["name"].(string),
		Priority: nodegroupData["priority"].(int),
		State:    nodegroupData["state"].(string),
		Sticky:   nodegroupData["sticky"].(string),
		Strict:   nodegroupData["strict"].(string),
	}
	err := client.UpdateUnnamedResource("clusternodegroup", &clusternodegroup)
	if err != nil {
		return err
	}
	return nil
}

func deleteSingleClusterNodegroup(d *schema.ResourceData, meta interface{}, nodegroupData map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleClusterNodegroup")
	client := meta.(*NetScalerNitroClient).client

	nodeName := nodegroupData["name"].(string)
	err := client.DeleteResource("clusternodegroup", nodeName)
	if err != nil {
		return err
	}
	return nil
}

func addSingleClusterNodegroup(d *schema.ResourceData, meta interface{}, nodegroupData map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleClusterNodegroup")

	client := meta.(*NetScalerNitroClient).client

	log.Printf("[DEBUG]  citrixadc-provider: Adding nodegroup %v", nodegroupData)

	clusternodegroup := cluster.Clusternodegroup{
		Name:     nodegroupData["name"].(string),
		Priority: nodegroupData["priority"].(int),
		State:    nodegroupData["state"].(string),
		Sticky:   nodegroupData["sticky"].(string),
		Strict:   nodegroupData["strict"].(string),
	}

	// Add cluster nodegroup on cluster configuration coordinator
	_, err := client.AddResource("clusternodegroup", clusternodegroup.Name, &clusternodegroup)
	if err != nil {
		return err
	}

	return nil
}

func readClusterNodes(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readClusterNodes")
	client := meta.(*NetScalerNitroClient).client
	clusternodes, err := client.FindAllResources("clusternode")
	if err != nil {
		return err
	}
	log.Printf("[DEBUG]  citrixadc-provider: fetched clusternodes %v", clusternodes)

	processedClusterNodes := make([]interface{}, len(clusternodes))
	for i, clusternode := range clusternodes {
		processedClusterNodes[i] = make(map[string]interface{})
		node := processedClusterNodes[i].(map[string]interface{})

		node["nodeid"], err = strconv.Atoi(clusternode["nodeid"].(string))
		if err != nil {
			return err
		}

		if val, ok := clusternode["clearnodegroupconfig"]; ok {
			node["clearnodegroupconfig"] = val.(string)
		}

		if _, ok := clusternode["backplane"]; ok {
			node["backplane"] = clusternode["backplane"].(string)
		}
		node["delay"] = int(clusternode["delay"].(float64))
		node["ipaddress"] = clusternode["ipaddress"].(string)
		node["nodegroup"] = clusternode["nodegroup"].(string)

		node["priority"], err = strconv.Atoi(clusternode["priority"].(string))
		if err != nil {
			return err
		}
		node["state"] = clusternode["state"].(string)
		node["tunnelmode"] = clusternode["tunnelmode"].(string)
	}

	updatedSet := schema.NewSet(clusternodeMappingHash, processedClusterNodes)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("clusternode", updatedSet); err != nil {
		return err
	}
	return nil
}

func updateClusterNodes(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateClusterNodes")

	oldSet, newSet := d.GetChange("clusternode")

	oldNodes := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	newNodes := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	log.Printf("[DEBUG]  citrixadc-provider: Old nodes set %v", oldNodes)
	log.Printf("[DEBUG]  citrixadc-provider: New nodes set %v", newNodes)
	toRemoveNodes := make([]interface{}, 0, oldNodes.Len())
	toCreateNodes := make([]interface{}, 0, oldNodes.Len())

	// Inline updates and recreates
	for _, old := range oldNodes.List() {
		oldVal := old.(map[string]interface{})
		for _, new := range newNodes.List() {
			newVal := new.(map[string]interface{})

			describeNodeMapDiff(oldVal, newVal)
			if compareNeedsUpdate(oldVal, newVal) {

				log.Printf("[DEBUG]  citrixadc-provider: Updating node %v", oldVal["nodeid"])
				if err := updateSingleClusterNode(d, meta, newVal); err != nil {
					return err
				}
				break
			}
			if compareNeedsRecreate(oldVal, newVal) {
				toRemoveNodes = append(toRemoveNodes, oldVal)
				toCreateNodes = append(toCreateNodes, newVal)

				log.Printf("[DEBUG]  citrixadc-provider: recreatting node %v -> %v", oldVal["nodeid"], newVal["nodeid"])
				break
			}
		}
	}
	// Do the recreates here
	// We need to do all removes first for node swaps to work
	for _, v := range toRemoveNodes {
		if err := deleteSingleClusterNode(d, meta, v.(map[string]interface{}), true); err != nil {
			return err
		}
	}

	for _, v := range toCreateNodes {
		if err := addSingleClusterNode(d, meta, v.(map[string]interface{})); err != nil {
			return err
		}
	}

	// Create new nodes
	for _, new := range newNodes.List() {
		newVal := new.(map[string]interface{})
		needsCreate := true
		for _, old := range oldNodes.List() {
			oldVal := old.(map[string]interface{})
			describeNodeMapDiff(newVal, oldVal)
			if !compareNeedsCreate(newVal, oldVal) {
				needsCreate = false
				break
			}
		}
		if needsCreate {
			log.Printf("[DEBUG]  citrixadc-provider: creating node %v", newVal["nodeid"])
			if err := addSingleClusterNode(d, meta, newVal); err != nil {
				return err
			}

		}
	}

	// Delete old nodes
	for _, old := range oldNodes.List() {
		oldVal := old.(map[string]interface{})
		needsDelete := true
		for _, new := range newNodes.List() {
			newVal := new.(map[string]interface{})
			if !compareNeedsDelete(oldVal, newVal) {
				needsDelete = false
				break
			}
		}
		if needsDelete {
			err := deleteSingleClusterNode(d, meta, oldVal, true)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func compareNeedsUpdate(oldNodeMap, newNodeMap map[string]interface{}) bool {
	log.Printf("[DEBUG]  citrixadc-provider: In compareNeedsUpdate")

	// Different nodes
	if oldNodeMap["nodeid"].(int) != newNodeMap["nodeid"].(int) {
		return false
	}

	// Fallthrough

	// Check non updateable attributes
	// Any change in them necessitates recreate
	if oldNodeMap["ipaddress"].(string) != newNodeMap["ipaddress"].(string) {
		return false
	}

	if oldNodeMap["nodegroup"].(string) != newNodeMap["nodegroup"].(string) {
		return false
	}

	// Fallthrough

	// Check rest of the attributes for changes
	// Any change in them necessiates update

	log.Printf("[DEBUG]  citrixadc-provider: comparing node ids %v %v", oldNodeMap["nodeid"], newNodeMap["nodeid"])
	needsUpdate := false

	if oldNodeMap["backplane"].(string) != newNodeMap["backplane"].(string) {
		needsUpdate = true
	}

	if oldNodeMap["delay"].(int) != newNodeMap["delay"].(int) {
		log.Printf("[DEBUG]  citrixadc-provider: delays differ %v %v", oldNodeMap["delay"], newNodeMap["delay"])
		needsUpdate = true
	}

	if oldNodeMap["priority"].(int) != newNodeMap["priority"].(int) {
		log.Printf("[DEBUG]  citrixadc-provider: priorities differ %v %v", oldNodeMap["priority"], newNodeMap["priority"])
		needsUpdate = true
	}

	if oldNodeMap["state"].(string) != newNodeMap["state"].(string) {
		log.Printf("[DEBUG]  citrixadc-provider: state differ %v %v", oldNodeMap["state"], newNodeMap["state"])
		needsUpdate = true
	}

	if oldNodeMap["tunnelmode"].(string) != newNodeMap["tunnelmode"].(string) {
		log.Printf("[DEBUG]  citrixadc-provider: tunnelmode differ %v %v", oldNodeMap["tunnelmode"], newNodeMap["tunnelmode"])
		needsUpdate = true
	}

	return needsUpdate
}

func compareNeedsRecreate(oldNodeMap, newNodeMap map[string]interface{}) bool {
	log.Printf("[DEBUG]  citrixadc-provider: In compareNeedsRecreate")

	// Different nodes
	if oldNodeMap["nodeid"].(int) != newNodeMap["nodeid"].(int) {
		return false
	}

	// Fallthrough

	// Check non updateable attributes
	// Any change in them necessitates recreate
	log.Printf("[DEBUG]  citrixadc-provider: comparing node ids %v %v", oldNodeMap["nodeid"], newNodeMap["nodeid"])
	needsRecreate := false
	if oldNodeMap["ipaddress"].(string) != newNodeMap["ipaddress"].(string) {
		log.Printf("[DEBUG]  citrixadc-provider: ipaddress differ %v %v", oldNodeMap["ipaddress"], newNodeMap["ipaddress"])
		needsRecreate = true
	}

	if oldNodeMap["nodegroup"].(string) != newNodeMap["nodegroup"].(string) {
		log.Printf("[DEBUG]  citrixadc-provider: nodegroup differ %v %v", oldNodeMap["nodegroup"], newNodeMap["nodegroup"])
		needsRecreate = true
	}
	return needsRecreate
}

func compareNeedsCreate(newNodeMap, oldNodeMap map[string]interface{}) bool {
	log.Printf("[DEBUG]  citrixadc-provider: In compareNeedsCreate")

	// Different nodes
	if oldNodeMap["nodeid"].(int) != newNodeMap["nodeid"].(int) {
		return true
	} else {
		return false
	}
}

func compareNeedsDelete(oldNodeMap, newNodeMap map[string]interface{}) bool {
	log.Printf("[DEBUG]  citrixadc-provider: In compareNeedsDelete")
	if oldNodeMap["nodeid"] != newNodeMap["nodeid"] {
		return true
	} else {
		return false
	}

}

func addSingleClusterNode(d *schema.ResourceData, meta interface{}, nodeData map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleClusterNode")
	client := meta.(*NetScalerNitroClient).client

	log.Printf("[DEBUG]  citrixadc-provider: Adding single node %v", nodeData)

	// Add cluster node at CLIP
	clusternode := cluster.Clusternode{
		Backplane:            nodeData["backplane"].(string),
		Clearnodegroupconfig: nodeData["clearnodegroupconfig"].(string),
		Delay:                nodeData["delay"].(int),
		Ipaddress:            nodeData["ipaddress"].(string),
		Nodegroup:            nodeData["nodegroup"].(string),
		Nodeid:               nodeData["nodeid"].(int),
		Priority:             nodeData["priority"].(int),
		State:                nodeData["state"].(string),
		Tunnelmode:           nodeData["tunnelmode"].(string),
	}

	// Add cluster node on cluster configuration coordinator
	_, err := client.AddResource("clusternode", strconv.FormatUint(uint64(clusternode.Nodeid), 10), &clusternode)
	if err != nil {
		return err
	}

	// Register node ip to the CLIP if flag is set
	if nodeData["addsnip"].(bool) {
		nodeNsip := nsip{
			Ownernode:  strconv.Itoa(nodeData["nodeid"].(int)),
			Ipaddress:  nodeData["snip_ipaddress"].(string),
			Mgmtaccess: "ENABLED",
			Netmask:    nodeData["snip_netmask"].(string),
			Type:       "SNIP",
		}
		log.Printf("[DEBUG]  citrixadc-provider: registering node ip %v", nodeNsip)
		_, err := client.AddResource(service.Nsip.Type(), nodeData["ipaddress"].(string), &nodeNsip)
		if err != nil {
			return err
		}
	}

	// Do the VTYSH commands if flag is set
	if nodeData["vtysh_enable"].(bool) {
		cmdList := nodeData["vtysh"].([]interface{})
		for _, val := range cmdList {
			vtyshCmdString := val.(string)

			routerdynamicrouting := router.Routerdynamicrouting{
				Commandstring: vtyshCmdString,
			}
			err := client.ActOnResource("routerdynamicrouting", &routerdynamicrouting, "apply")
			if err != nil {
				return err
			}
		}
	}

	// Instantiate node client
	nodeClient, err := instantiateNodeClient(d, meta, nodeData)
	if err != nil {
		return err
	}

	// join cluster from node
	log.Printf("[DEBUG]  citrixadc-provider: node id  %v joining cluster", clusternode.Nodeid)
	cluster := cluster.Cluster{
		Clip:     d.Get("clip").(string),
		Password: meta.(*NetScalerNitroClient).Password,
	}
	err = nodeClient.ActOnResource("cluster", &cluster, "join")
	if err != nil {
		return err
	}

	// Save config
	nodeClient.SaveConfig()

	// Reboot node
	log.Printf("[DEBUG]  citrixadc-provider: Rebooting first node")
	reboot := ns.Reboot{
		Warm: true,
	}
	if err := nodeClient.ActOnResource("reboot", &reboot, ""); err != nil {
		return err
	}

	// Poll cluster configuration coordinator for added node id status

	var node_add_poll_interval time.Duration
	if node_add_poll_interval, err = time.ParseDuration(d.Get("node_add_poll_interval").(string)); err != nil {
		return err
	}

	var node_add_poll_delay time.Duration
	if node_add_poll_delay, err = time.ParseDuration(d.Get("node_add_poll_delay").(string)); err != nil {
		return err
	}

	var node_add_total_timeout time.Duration
	if node_add_total_timeout, err = time.ParseDuration(d.Get("node_add_total_timeout").(string)); err != nil {
		return err
	}

	nodeid := nodeData["nodeid"].(int)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"adding_node"},
		Target:       []string{"node_added"},
		Refresh:      pollClusterNodeWithid(d, meta, nodeid),
		Timeout:      node_add_total_timeout,
		PollInterval: node_add_poll_interval,
		Delay:        node_add_poll_delay,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}

	return nil
}

func deleteSingleClusterNode(d *schema.ResourceData, meta interface{}, nodeData map[string]interface{}, wait_clip bool) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleClusterNode")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG]  citrixadc-provider: Deleting single node %v", nodeData)

	var err error

	nodeId := strconv.Itoa(nodeData["nodeid"].(int))
	// When we delete the CCO node the following error is expected
	if err = client.DeleteResource("clusternode", nodeId); err != nil {
		if !strings.Contains(err.Error(), "read: connection reset by peer") {
			return err
		} else {
			log.Printf("[DEBUG]  citrixadc-provider: lost CLIP when deleted node %v", nodeData["nodeid"])
		}
	}

	// Wait for clip if flag is set
	if wait_clip {
		var clip_migration_total_timeout time.Duration
		if clip_migration_total_timeout, err = time.ParseDuration(d.Get("clip_migration_total_timeout").(string)); err != nil {
			return err
		}

		var clip_migration_poll_timeout time.Duration
		if clip_migration_poll_timeout, err = time.ParseDuration(d.Get("clip_migration_poll_timeout").(string)); err != nil {
			return err
		}

		var clip_migration_poll_interval time.Duration
		if clip_migration_poll_interval, err = time.ParseDuration(d.Get("clip_migration_poll_interval").(string)); err != nil {
			return err
		}

		var clip_migration_poll_delay time.Duration
		if clip_migration_poll_delay, err = time.ParseDuration(d.Get("clip_migration_poll_delay").(string)); err != nil {
			return err
		}

		stateConf := &resource.StateChangeConf{
			Pending:      []string{"cluster_unreachable"},
			Target:       []string{"cluster_reachable"},
			Refresh:      resourceClusterPoll(d, meta, clip_migration_poll_timeout),
			Timeout:      clip_migration_total_timeout,
			PollInterval: clip_migration_poll_interval,
			Delay:        clip_migration_poll_delay,
		}

		_, err = stateConf.WaitForState()
		if err != nil {
			return err
		}
	}
	return nil
}

func updateSingleClusterNode(d *schema.ResourceData, meta interface{}, nodeData map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSingleClusterNode")
	client := meta.(*NetScalerNitroClient).client

	log.Printf("[DEBUG]  citrixadc-provider: Updating single node %v", nodeData)

	// Only include attributes that can be present in HTTP PUT operation
	clusternode := cluster.Clusternode{
		Backplane:  nodeData["backplane"].(string),
		Delay:      nodeData["delay"].(int),
		Nodeid:     nodeData["nodeid"].(int),
		Priority:   nodeData["priority"].(int),
		State:      nodeData["state"].(string),
		Tunnelmode: nodeData["tunnelmode"].(string),
	}

	if err := client.UpdateUnnamedResource("clusternode", &clusternode); err != nil {
		return err
	}
	return nil
}

func describeNodeMapDiff(left, right map[string]interface{}) {
	log.Printf("[DEBUG]  citrixadc-provider: describeNodeMapDiff")
	keys := []string{
		"backplane",
		"clearnodegroupconfig",
		"delay",
		"ipaddress",
		"nodegroup",
		"nodeid",
		"priority",
		"state",
		"tunnelmode",
		"endpoint",
		"username",
		"password",
		"insecure_skip_verify",
	}
	for _, key := range keys {
		leftVal, leftOk := left[key]
		rightVal, rightOk := right[key]
		var leftMsg string
		var rightMsg string

		if leftOk {
			leftMsg = fmt.Sprintf("\"%v\" %T", leftVal, leftVal)
		} else {
			leftMsg = "absent"
		}

		if rightOk {
			rightMsg = fmt.Sprintf("\"%v\" %T", rightVal, rightVal)
		} else {
			rightMsg = "absent"
		}
		log.Printf("key %s left:%s right:%s equal:%v", key, leftMsg, rightMsg, leftVal == rightVal)
	}
}

func instantiateNodeClient(d *schema.ResourceData, meta interface{}, nodeMap map[string]interface{}) (*service.NitroClient, error) {
	log.Printf("[DEBUG]  citrixadc-provider: In instantiateNodeClient")

	var nodeEndpoint string
	var nodeUsername string
	var nodePassword string
	var nodeSslVerrify bool

	// Required field always exists
	nodeEndpoint = nodeMap["endpoint"].(string)

	// Default to provider credential
	nodeUsername = nodeMap["username"].(string)
	if nodeUsername == "" {
		nodeUsername = meta.(*NetScalerNitroClient).Username
	}

	// Default to provider credential
	nodePassword = nodeMap["password"].(string)
	if nodePassword == "" {
		nodePassword = meta.(*NetScalerNitroClient).Password
	}

	// Always exists has default value
	nodeSslVerrify = !nodeMap["insecure_skip_verify"].(bool)

	params := service.NitroParams{
		Url:       nodeEndpoint,
		Username:  nodeUsername,
		Password:  nodePassword,
		SslVerify: nodeSslVerrify,
	}

	log.Printf("[DEBUG]  citrixadc-provider: node client params %v", params)

	nodeClient, err := service.NewNitroClientFromParams(params)
	return nodeClient, err
}

func resourceClusterPoll(d *schema.ResourceData, meta interface{}, timeout time.Duration) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] netscaler-provider: In resourceClusterPoll")
		err := pollNode(d, meta, timeout)
		if err != nil {
			if err.Error() == "Timeout" {
				log.Printf("[DEBUG] netscaler-provider: Cluster poll result \"cluster_unreachable\"")
				return nil, "cluster_unreachable", nil
			} else {
				return nil, "cluster_unreachable", err
			}
		}
		log.Printf("[DEBUG] netscaler-provider: Cluster poll result \"cluster_reachable\"")
		return "cluster_reachable", "cluster_reachable", nil
	}
}

func pollNode(d *schema.ResourceData, meta interface{}, timeout time.Duration) error {
	log.Printf("[DEBUG] netscaler-provider: In pollLicense")

	username := meta.(*NetScalerNitroClient).Username
	password := meta.(*NetScalerNitroClient).Password
	endpoint := meta.(*NetScalerNitroClient).Endpoint
	url := fmt.Sprintf("%s/nitro/v1/config/nslicense", endpoint)

	var err error
	c := http.Client{
		Timeout: timeout,
	}
	buff := &bytes.Buffer{}
	req, _ := http.NewRequest("GET", url, buff)
	req.Header.Set("X-NITRO-USER", username)
	req.Header.Set("X-NITRO-PASS", password)
	resp, err := c.Do(req)
	if err != nil {
		if !strings.Contains(err.Error(), "Client.Timeout exceeded") {
			// Unexpected error
			return err
		} else {
			// Expected timeout error
			return fmt.Errorf("Timeout")
		}
	} else {
		log.Printf("Status code is %v\n", resp.Status)
	}
	// No error
	return nil
}

func pollClusterNodeWithid(d *schema.ResourceData, meta interface{}, nodeid int) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] citrixadc-provider: In pollClusterNodeWithid")
		client := meta.(*NetScalerNitroClient).client

		data, err := client.FindAllResources("clusternode")
		if err != nil {
			return nil, "adding_node", err
		}

		nodeidFound := false
		return_state := "adding_node"
		// verify that the nodeid entry is healthy
		for _, val := range data {
			val_nodeid, err := strconv.Atoi(val["nodeid"].(string))
			if err != nil {
				return nil, "adding_node", err
			}
			if val_nodeid == nodeid {
				nodeidFound = true
				if val["masterstate"] == "ACTIVE" {
					return_state = "node_added"
				} else {
					log.Printf("[DEBUG] citrixadc-provider: node masterstate is %s", val["masterstate"])
				}
				break
			}
		}
		// There is something very wrong
		if !nodeidFound {
			return nil, "adding_node", fmt.Errorf("Node id %v not in retrieved nodes list", nodeid)
		}
		// Node is being added. Wait.
		if return_state == "adding_node" {
			return nil, "adding_node", nil
		} else {
			// Node added. Continue
			return "node_added", "node_added", nil
		}
	}
}
