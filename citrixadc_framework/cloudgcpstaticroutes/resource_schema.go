package cloudgcpstaticroutes

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cloud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. Because these attributes revert to their
// default (or absence) after unset, marking the plan unknown also avoids a
// "provider produced inconsistent result" error, which a static Default would
// trigger.
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// CloudgcpstaticroutesResourceModel describes the resource data model.
type CloudgcpstaticroutesResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Status  types.String `tfsdk:"status"`
	Project types.String `tfsdk:"project"`
}

func (r *CloudgcpstaticroutesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudgcpstaticroutes resource.",
			},
			"status": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "status to push routes or not. Possible values = ENABLED, DISABLED",
			},
			"project": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "GCP project name for which static routes functionality is enabled.",
			},
		},
	}
}

func cloudgcpstaticroutesGetThePayloadFromthePlan(ctx context.Context, data *CloudgcpstaticroutesResourceModel) cloud.Cloudgcpstaticroutes {
	tflog.Debug(ctx, "In cloudgcpstaticroutesGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudgcpstaticroutes := cloud.Cloudgcpstaticroutes{}
	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		cloudgcpstaticroutes.Status = data.Status.ValueString()
	}
	if !data.Project.IsNull() && !data.Project.IsUnknown() {
		cloudgcpstaticroutes.Project = data.Project.ValueString()
	}

	return cloudgcpstaticroutes
}

func cloudgcpstaticroutesSetAttrFromGet(ctx context.Context, data *CloudgcpstaticroutesResourceModel, getResponseData map[string]interface{}) *CloudgcpstaticroutesResourceModel {
	tflog.Debug(ctx, "In cloudgcpstaticroutesSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["status"]; ok && val != nil {
		data.Status = types.StringValue(val.(string))
	} else {
		data.Status = types.StringNull()
	}
	if val, ok := getResponseData["project"]; ok && val != nil {
		data.Project = types.StringValue(val.(string))
	} else {
		data.Project = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID (singleton)
	data.Id = types.StringValue("cloudgcpstaticroutes-config")

	return data
}
