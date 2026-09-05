package sslprofile_ecccurve_binding

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SslprofileEcccurveBindingResourceModel describes the resource data model.
//
// This mirrors the SDK v2 contract: ecccurvename is a LIST of curve names, each
// of which is bound to the SSL profile as a separate NITRO binding, and
// remove_existing_ecccurve_binding controls whether the profile's pre-existing
// (default) ecccurve bindings are cleared first.
type SslprofileEcccurveBindingResourceModel struct {
	Id                            types.String `tfsdk:"id"`
	Name                          types.String `tfsdk:"name"`
	Ecccurvename                  types.List   `tfsdk:"ecccurvename"`
	RemoveExistingEcccurveBinding types.Bool   `tfsdk:"remove_existing_ecccurve_binding"`
}

func (r *SslprofileEcccurveBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslprofile_ecccurve_binding resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the SSL profile.",
			},
			"ecccurvename": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "Named ECC curves bound to the SSL profile.",
			},
			"remove_existing_ecccurve_binding": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Remove existing ECC curve bindings from the SSL profile before binding the configured curves.",
			},
		},
	}
}
