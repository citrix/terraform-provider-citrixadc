package citrixadc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestPollNode_Returns200(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"nslicense": {}}`))
	}))
	defer server.Close()

	meta := &NetScalerNitroClient{
		Username: "nsroot",
		Password: "nsroot",
		Endpoint: server.URL,
	}

	err := pollNode(nil, meta, 5*time.Second)
	if err != nil {
		t.Fatalf("Expected no error for 200 OK, got: %v", err)
	}
}

func TestPollNode_Returns503(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	meta := &NetScalerNitroClient{
		Username: "nsroot",
		Password: "nsroot",
		Endpoint: server.URL,
	}

	err := pollNode(nil, meta, 5*time.Second)
	if err == nil {
		t.Fatal("Expected error for 503, got nil")
	}
	if err.Error() != "Timeout" {
		t.Fatalf("Expected 'Timeout' error, got: %v", err)
	}
}

func TestPollNode_Returns401(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	meta := &NetScalerNitroClient{
		Username: "nsroot",
		Password: "nsroot",
		Endpoint: server.URL,
	}

	err := pollNode(nil, meta, 5*time.Second)
	if err == nil {
		t.Fatal("Expected error for 401, got nil")
	}
	if err.Error() != "Timeout" {
		t.Fatalf("Expected 'Timeout' error, got: %v", err)
	}
}

func TestPollNode_ConnectionRefused(t *testing.T) {
	// Use a port that nothing is listening on
	meta := &NetScalerNitroClient{
		Username: "nsroot",
		Password: "nsroot",
		Endpoint: "http://127.0.0.1:19999",
	}

	err := pollNode(nil, meta, 2*time.Second)
	if err == nil {
		t.Fatal("Expected error for connection refused, got nil")
	}
}

func TestPollNode_Timeout(t *testing.T) {
	// Server that delays longer than the timeout
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	meta := &NetScalerNitroClient{
		Username: "nsroot",
		Password: "nsroot",
		Endpoint: server.URL,
	}

	err := pollNode(nil, meta, 1*time.Second)
	if err == nil {
		t.Fatal("Expected timeout error, got nil")
	}
	if err.Error() != "Timeout" {
		t.Fatalf("Expected 'Timeout' error, got: %v", err)
	}
}

func TestResourceClusterPoll_Returns200(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"nslicense": {}}`))
	}))
	defer server.Close()

	meta := &NetScalerNitroClient{
		Username: "nsroot",
		Password: "nsroot",
		Endpoint: server.URL,
	}

	pollFunc := resourceClusterPoll(nil, meta, 5*time.Second)
	result, state, err := pollFunc()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if state != "cluster_reachable" {
		t.Fatalf("Expected state 'cluster_reachable', got: %s", state)
	}
	if result != "cluster_reachable" {
		t.Fatalf("Expected result 'cluster_reachable', got: %v", result)
	}
}

func TestResourceClusterPoll_Returns503(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	meta := &NetScalerNitroClient{
		Username: "nsroot",
		Password: "nsroot",
		Endpoint: server.URL,
	}

	pollFunc := resourceClusterPoll(nil, meta, 5*time.Second)
	_, state, err := pollFunc()
	if err != nil {
		t.Fatalf("Expected no error (poll should retry), got: %v", err)
	}
	if state != "cluster_unreachable" {
		t.Fatalf("Expected state 'cluster_unreachable', got: %s", state)
	}
}

// Acceptance test environment variables:
//   NS_URL          - CLIP endpoint (e.g., http://10.102.201.48)
//   NS_LOGIN        - 
//   NS_PASSWORD     - 
//   CLUSTER_CLIP    - Cluster IP address (e.g., 10.102.201.48)
//   CLUSTER_NODE0_IP       - First node NSIP (e.g., 10.102.201.42)
//   CLUSTER_NODE0_ENDPOINT - First node endpoint (e.g., http://10.102.201.42)
//   CLUSTER_NODE1_IP       - Second node NSIP (e.g., 10.102.201.213)
//   CLUSTER_NODE1_ENDPOINT - Second node endpoint (e.g., http://10.102.201.213)

func clusterAccPreCheck(t *testing.T) {
	testAccPreCheck(t)
	for _, env := range []string{"CLUSTER_CLIP", "CLUSTER_NODE0_IP", "CLUSTER_NODE0_ENDPOINT"} {
		if os.Getenv(env) == "" {
			t.Fatalf("%s must be set for cluster acceptance tests", env)
		}
	}
}

func clusterTwoNodeAccPreCheck(t *testing.T) {
	clusterAccPreCheck(t)
	for _, env := range []string{"CLUSTER_NODE1_IP", "CLUSTER_NODE1_ENDPOINT"} {
		if os.Getenv(env) == "" {
			t.Fatalf("%s must be set for two-node cluster acceptance tests", env)
		}
	}
}

func testAccCheckClusterDestroy(s *terraform.State) error {
	client, err := testAccGetClient()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cluster" {
			continue
		}
		data, _ := client.FindAllResources("clusterinstance")
		if len(data) > 0 {
			clid := data[0]["clid"]
			if clid != nil && clid != "" && clid != "0" {
				return fmt.Errorf("cluster instance %s still exists", rs.Primary.ID)
			}
		}
	}
	return nil
}

func TestAccCitrixAdcCluster_oneNode(t *testing.T) {
	clip := os.Getenv("CLUSTER_CLIP")
	node0IP := os.Getenv("CLUSTER_NODE0_IP")
	node0Endpoint := os.Getenv("CLUSTER_NODE0_ENDPOINT")

	config := fmt.Sprintf(`
resource "citrixadc_cluster" "tf_cluster" {
  clid          = 1
  clip          = "%s"
  hellointerval = 200
  preemption    = "ENABLED"
  quorumtype    = "NONE"

  clusternode {
    nodeid     = 0
    delay      = 0
    priority   = 30
    endpoint   = "%s"
    backplane  = "0/1/1"
    ipaddress  = "%s"
    tunnelmode = "NONE"
    nodegroup  = "DEFAULT_NG"
    state      = "ACTIVE"
  }

  bootstrap_poll_delay    = "60s"
  bootstrap_poll_interval = "60s"
  bootstrap_total_timeout = "10m"
}
`, clip, node0Endpoint, node0IP)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { clusterAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "clid", "1"),
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "clip", clip),
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "hellointerval", "200"),
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "preemption", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "quorumtype", "NONE"),
				),
			},
		},
	})
}

func TestAccCitrixAdcCluster_twoNode(t *testing.T) {
	clip := os.Getenv("CLUSTER_CLIP")
	node0IP := os.Getenv("CLUSTER_NODE0_IP")
	node0Endpoint := os.Getenv("CLUSTER_NODE0_ENDPOINT")
	node1IP := os.Getenv("CLUSTER_NODE1_IP")
	node1Endpoint := os.Getenv("CLUSTER_NODE1_ENDPOINT")

	config := fmt.Sprintf(`
resource "citrixadc_cluster" "tf_cluster" {
  clid          = 1
  clip          = "%s"
  hellointerval = 200
  preemption    = "ENABLED"
  quorumtype    = "MAJORITY"

  clusternode {
    nodeid     = 0
    delay      = 0
    priority   = 30
    endpoint   = "%s"
    backplane  = "0/1/1"
    ipaddress  = "%s"
    tunnelmode = "NONE"
    nodegroup  = "DEFAULT_NG"
    state      = "ACTIVE"
  }

  clusternode {
    nodeid     = 1
    delay      = 0
    priority   = 31
    endpoint   = "%s"
    backplane  = "1/1/1"
    ipaddress  = "%s"
    tunnelmode = "NONE"
    state      = "ACTIVE"
  }

  bootstrap_poll_delay    = "60s"
  bootstrap_poll_interval = "60s"
  bootstrap_total_timeout = "10m"
}
`, clip, node0Endpoint, node0IP, node1Endpoint, node1IP)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { clusterTwoNodeAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "clid", "1"),
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "clip", clip),
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "hellointerval", "200"),
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "preemption", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "quorumtype", "MAJORITY"),
					resource.TestCheckResourceAttr("citrixadc_cluster.tf_cluster", "clusternode.#", "2"),
				),
			},
		},
	})
}
