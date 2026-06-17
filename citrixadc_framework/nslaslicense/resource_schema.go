package nslaslicense

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NOTE: Applying a LAS license via this resource is DISRUPTIVE and
// non-idempotent on the appliance. NITRO exposes only the `apply` action
// (POST ?action=apply); there is no get/add/delete/update endpoint.

// NslaslicenseResourceModel describes the resource data model.
type NslaslicenseResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Filelocation   types.String `tfsdk:"filelocation"`
	Filename       types.String `tfsdk:"filename"`
	Fixedbandwidth types.Bool   `tfsdk:"fixedbandwidth"`
}

func (r *NslaslicenseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslaslicense resource.",
			},
			"filelocation": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "location of the file on Citrix ADC.",
			},
			"filename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the file. It should not include filepath.",
			},
			"fixedbandwidth": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "apply fixed bandwidth license on ADC",
			},
		},
	}
}

func nslaslicenseGetThePayloadFromthePlan(ctx context.Context, data *NslaslicenseResourceModel) ns.Nslaslicense {
	tflog.Debug(ctx, "In nslaslicenseGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nslaslicense := ns.Nslaslicense{}
	if !data.Filelocation.IsNull() && !data.Filelocation.IsUnknown() {
		nslaslicense.Filelocation = data.Filelocation.ValueString()
	}
	if !data.Filename.IsNull() && !data.Filename.IsUnknown() {
		nslaslicense.Filename = data.Filename.ValueString()
	}
	if !data.Fixedbandwidth.IsNull() && !data.Fixedbandwidth.IsUnknown() {
		nslaslicense.Fixedbandwidth = data.Fixedbandwidth.ValueBool()
	}

	return nslaslicense
}
