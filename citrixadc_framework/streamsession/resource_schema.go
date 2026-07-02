package streamsession

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/stream"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// StreamsessionResourceModel describes the resource data model.
type StreamsessionResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r *StreamsessionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the streamsession resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the stream identifier.",
			},
		},
	}
}

func streamsessionGetThePayloadFromthePlan(ctx context.Context, data *StreamsessionResourceModel) stream.Streamsession {
	tflog.Debug(ctx, "In streamsessionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	streamsession := stream.Streamsession{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		streamsession.Name = data.Name.ValueString()
	}

	return streamsession
}
