# Load Balancing with HA Primary Node Detection

This example sets up an HA pair between two NetScaler ADC nodes, detects which node is primary, and configures load balancing only on that node. HA sync automatically propagates the configuration to the secondary.

## Architecture

```
                ┌──────────────────┐
                │   Terraform      │
                └────────┬─────────┘
                         │
            ┌────────────┴────────────┐
            ▼                         ▼
   ┌─────────────────┐      ┌─────────────────┐
   │  Node1 (NSIP)   │◄────►│  Node2 (NSIP)   │
   │  HA Primary     │  HA  │  HA Secondary   │
   │                 │ Sync │                 │
   │  LB VServer     │      │  (auto-synced)  │
   │  ├─ Service 1   │      │                 │
   │  └─ Service 2   │      │                 │
   └─────────────────┘      └─────────────────┘
```

## How It Works

### Phase 1 — HA Pair Formation

- Configures two `citrixadc` providers, one per node.
- Creates `citrixadc_hanode` resources from both sides to form the HA pair.
- Node1 adds node2 as peer first, then node2 adds node1 (sequenced via `depends_on`).
- The initiating node (node1) typically becomes Primary.

### Phase 2 — Primary Detection + LB Setup

- Queries each node's `hanode/0` NITRO API endpoint via `data "http"` to read its HA state.
- Parses the JSON response to determine which node is Primary vs Secondary.
- Uses `count = local.nodeX_is_primary ? 1 : 0` on every LB resource so only the primary gets configured:
  - **nsfeature** — Enables the LB feature.
  - **server** — Creates backend server objects.
  - **lbvserver** — Creates the LB virtual server (VIP).
  - **service** — Creates services bound to the LB VIP via `servername`.
- HA sync automatically propagates the configuration from primary to secondary.

## Why Two Phases?

Terraform evaluates data sources at **plan time**, before any resources are created. On a single apply:

1. Data sources query `hanode/0` → both nodes are standalone → both report "Primary".
2. `count` resolves for both nodes → LB resources are created on **both**.
3. The `citrixadc_hanode` resources get created too late — data sources already ran.

With `-target` in Phase 1, Terraform skips the data sources and only creates the HA pair. In Phase 2, the HA pair exists, so data sources correctly return one Primary and one Secondary.

## Usage

### 1. Configure Variables

Copy the example tfvars and fill in your values:

```bash
cp terraform.tfvars.example terraform.tfvars
# Edit terraform.tfvars with your node IPs, credentials, VIP, and backend servers
```

### 2. Initialize

```bash
terraform init
```

### 3. Phase 1 — Create the HA Pair

```bash
terraform apply \
  -target=citrixadc_hanode.ha_peer_on_node1 \
  -target=citrixadc_hanode.ha_peer_on_node2
```

### 4. Phase 2 — Detect Primary and Configure LB

```bash
terraform apply
```

## Outputs

| Output             | Description                                        |
|--------------------|----------------------------------------------------|
| `node1_is_primary` | Whether node1 is the HA primary (`true`/`false`)   |
| `node2_is_primary` | Whether node2 is the HA primary (`true`/`false`)   |
| `primary_node_ip`  | NSIP of whichever node is primary                  |

## Variables

| Variable             | Description                          | Default        |
|----------------------|--------------------------------------|----------------|
| `node1_nsip`         | NSIP of the first NetScaler node     | —              |
| `node2_nsip`         | NSIP of the second NetScaler node    | —              |
| `adc_admin_username` | Admin username for both nodes        | `nsroot`       |
| `adc_admin_password` | Admin password for both nodes        | —              |
| `lb_vserver_name`    | Name of the LB virtual server       | `app_lb_vserver` |
| `lb_vip`             | Virtual IP for the load balancer     | —              |
| `lb_port`            | Port for the LB virtual server      | `80`           |
| `lb_service_type`    | Service type (HTTP, HTTPS, TCP, etc.)| `HTTP`         |
| `lb_method`          | Load balancing method                | `ROUNDROBIN`   |
| `backend_servers`    | List of backend servers (name, ip, port) | 2 sample servers |

## File Structure

```
lb_ha_primary_detection/
├── main.tf                 # HA pair + primary detection + LB resources
├── variables.tf            # Variable definitions
├── terraform.tfvars        # Your actual values (git-ignored)
├── terraform.tfvars.example# Template with placeholder values
└── README.md               # This file
```

## Cleanup

```bash
terraform destroy
```

## Notes

- `rpcnodepassword` is write-only — Terraform cannot read it back from the ADC. This causes `terraform plan` to always show the hanode resources as needing replacement. This is a known provider limitation and is harmless if ignored.
- The `citrixadc_service` resources use `servername` (not `ip`) to reference the already-created `citrixadc_server` objects, avoiding 409 Conflict errors from duplicate server auto-creation.
