package appfwcustomsettings

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwcustomsettingsResourceModel describes the resource data model.
type AppfwcustomsettingsResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Target types.String `tfsdk:"target"`
}

func (r *AppfwcustomsettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwcustomsettings resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"target": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
		},
	}
}

func appfwcustomsettingsGetThePayloadFromthePlan(ctx context.Context, data *AppfwcustomsettingsResourceModel) appfw.Appfwcustomsettings {
	tflog.Debug(ctx, "In appfwcustomsettingsGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwcustomsettings := appfw.Appfwcustomsettings{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwcustomsettings.Name = data.Name.ValueString()
	}
	if !data.Target.IsNull() && !data.Target.IsUnknown() {
		appfwcustomsettings.Target = data.Target.ValueString()
	}

	return appfwcustomsettings
}
