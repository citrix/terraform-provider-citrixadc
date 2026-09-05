package nsvpxparam

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsvpxparamDataSourceModel is the data-source-specific model, decoupled from
// NsvpxparamResourceModel. A data source is a pure read surface, so it exposes the
// full GET projection: the read/write attributes (as Computed outputs) plus the
// read-only attributes the resource deliberately omits.
type NsvpxparamDataSourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Cpuyield            types.String `tfsdk:"cpuyield"`
	Kvmvirtiomultiqueue types.String `tfsdk:"kvmvirtiomultiqueue"`
	Masterclockcpu1     types.String `tfsdk:"masterclockcpu1"`
	Ownernode           types.Int64  `tfsdk:"ownernode"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nsvpxparam.json). Never settable; populated from GET.
	Vpxenvironment      types.String `tfsdk:"vpxenvironment"`
	Memorystatus        types.String `tfsdk:"memorystatus"`
	Cloudproductcode    types.String `tfsdk:"cloudproductcode"`
	Vpxoemcode          types.Int64  `tfsdk:"vpxoemcode"`
	Technicalsupportpin types.String `tfsdk:"technicalsupportpin"`
}

func NsvpxparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cpuyield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting applicable in virtual appliances, is to affect the cpu yield(relinquishing the cpu resources) in any hypervised environment.\n\n* There are 3 options for the behavior:\n1. YES - Allow the Virtual Appliance to yield its vCPUs periodically, if there is no data traffic.\n2. NO - Virtual Appliance will not yield the vCPU.\n3. DEFAULT - Restores the default behaviour, according to the license.\n\n* Its behavior in different scenarios:\n1. As this setting is node specific only, it will not be propagated to other nodes, when executed on Cluster(CLIP) and HA(Primary).\n2. In cluster setup, use '-ownerNode' to specify ID of the cluster node.\n3. This setting is a system wide implementation and not granular to vCPUs.\n4. No effect on the management PE.",
			},
			"kvmvirtiomultiqueue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting applicable on KVM VPX with virtio NICs, is to configure multiple queues for all virtio interfaces.\n\n* There are 2 options for this behavior:\n1. YES - Allows VPX to use multiple queues for each virtio interface as configured through the KVM Hypervisor.\n2. NO - Each virtio interface within VPX will use a single queue for transmit and receive.\n\n* Its behavior in different scenarios:\n1. As this setting is node specific only, it will not be propagated to other nodes, when executed on Cluster(CLIP) and HA(Primary).\n2. In cluster setup, use '-ownerNode' to specify ID of the cluster node.",
			},
			"masterclockcpu1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is deprecated.",
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the cluster node for which you are setting the cpuyield and/or KVMVirtioMultiqueue. It can be configured only through the cluster IP address.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"vpxenvironment": schema.StringAttribute{
				Computed:    true,
				Description: "Shows VPX Running Environment (VPX, AWSCLOUD, OPENSTACKCLOUD, AZURECLOUD, GOOGLECLOUD, ALICLOUD, IBMCLOUD).",
			},
			"memorystatus": schema.StringAttribute{
				Computed:    true,
				Description: "Provides the information about memory status (Sufficent, Insufficent).",
			},
			"cloudproductcode": schema.StringAttribute{
				Computed:    true,
				Description: "Cloud Product Code Description.",
			},
			"vpxoemcode": schema.Int64Attribute{
				Computed:    true,
				Description: "OEM Distribution Code.",
			},
			"technicalsupportpin": schema.StringAttribute{
				Computed:    true,
				Description: "Technical Support PIN for cloud subscription VMs.",
			},
		},
	}
}

// nsvpxparamDataSourceSetAttrFromGet projects a NITRO nsvpxparam GET response onto
// the data-source model. Attributes are simply filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func nsvpxparamDataSourceSetAttrFromGet(ctx context.Context, data *NsvpxparamDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsvpxparamDataSourceSetAttrFromGet Function")

	data.Cpuyield = utils.MapGetString(g, "cpuyield")
	data.Kvmvirtiomultiqueue = utils.MapGetString(g, "kvmvirtiomultiqueue")
	data.Masterclockcpu1 = utils.MapGetString(g, "masterclockcpu1")
	data.Ownernode = utils.MapGetInt64(g, "ownernode")

	// Datasource identity mirrors the ownernode being queried.
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))

	// Read-only attributes.
	data.Vpxenvironment = utils.MapGetString(g, "vpxenvironment")
	data.Memorystatus = utils.MapGetString(g, "memorystatus")
	data.Cloudproductcode = utils.MapGetString(g, "cloudproductcode")
	data.Vpxoemcode = utils.MapGetInt64(g, "vpxoemcode")
	data.Technicalsupportpin = utils.MapGetString(g, "technicalsupportpin")
}
