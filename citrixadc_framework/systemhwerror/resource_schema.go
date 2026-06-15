package systemhwerror

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemhwerrorResourceModel describes the resource data model.
type SystemhwerrorResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Diskcheck types.Bool   `tfsdk:"diskcheck"`
}

func (r *SystemhwerrorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemhwerror resource.",
			},
			"diskcheck": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Perform only disk error checking.",
			},
		},
	}
}

func systemhwerrorGetThePayloadFromthePlan(ctx context.Context, data *SystemhwerrorResourceModel) system.Systemhwerror {
	tflog.Debug(ctx, "In systemhwerrorGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemhwerror := system.Systemhwerror{}
	if !data.Diskcheck.IsNull() && !data.Diskcheck.IsUnknown() {
		systemhwerror.Diskcheck = data.Diskcheck.ValueBool()
	}

	return systemhwerror
}
