package kafkacluster

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// KafkaclusterDataSourceModel is the data-source-specific model, decoupled from
// KafkaclusterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (activesvc, totalsvc, topicname, numtopics). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares,
// which is why it cannot reuse the resource model.
type KafkaclusterDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/kafkacluster.json). Never settable; populated from GET.
	Activesvc types.Int64  `tfsdk:"activesvc"`
	Totalsvc  types.Int64  `tfsdk:"totalsvc"`
	Topicname types.String `tfsdk:"topicname"`
	Numtopics types.Int64  `tfsdk:"numtopics"`
}

func KafkaclusterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Kafka cluster",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"activesvc": schema.Int64Attribute{
				Computed:    true,
				Description: "Total active services bound to servicegroup.",
			},
			"totalsvc": schema.Int64Attribute{
				Computed:    true,
				Description: "Total services bound to servicegroup.",
			},
			"topicname": schema.StringAttribute{
				Computed:    true,
				Description: "Topic of the servicegroup.",
			},
			"numtopics": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of topic servicegroups bound.",
			},
		},
	}
}

// kafkaclusterDataSourceSetAttrFromGet projects a NITRO kafkacluster GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func kafkaclusterDataSourceSetAttrFromGet(ctx context.Context, data *KafkaclusterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In kafkaclusterDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read-only metadata.
	data.Activesvc = utils.MapGetInt64(g, "activesvc")
	data.Totalsvc = utils.MapGetInt64(g, "totalsvc")
	data.Topicname = utils.MapGetString(g, "topicname")
	data.Numtopics = utils.MapGetInt64(g, "numtopics")
}
