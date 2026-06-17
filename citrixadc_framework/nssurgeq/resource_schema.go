package nssurgeq

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NssurgeqResourceModel describes the resource data model.
type NssurgeqResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Port       types.Int64  `tfsdk:"port"`
	Servername types.String `tfsdk:"servername"`
}

func (r *NssurgeqResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nssurgeq resource.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of a virtual server, service or service group for which the SurgeQ must be flushed.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "port on which server is bound to the entity(Servicegroup).",
			},
			"servername": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of a service group member. This argument is needed when you want to flush the SurgeQ of a service group.",
			},
		},
	}
}

// nssurgeqGetThePayloadFromthePlan builds the action payload, including only the set args.
func nssurgeqGetThePayloadFromthePlan(ctx context.Context, data *NssurgeqResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nssurgeqGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		payload["name"] = data.Name.ValueString()
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		payload["servername"] = data.Servername.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		payload["port"] = int(data.Port.ValueInt64())
	}

	return payload
}
