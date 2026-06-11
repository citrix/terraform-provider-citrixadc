package fis_interface_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// FisInterfaceBindingResourceModel describes the resource data model.
type FisInterfaceBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ifnum     types.String `tfsdk:"ifnum"`
	Name      types.String `tfsdk:"name"`
	Ownernode types.Int64  `tfsdk:"ownernode"`
}

func (r *FisInterfaceBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the fis_interface_binding resource.",
			},
			"ifnum": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interface to be bound to the FIS, specified in slot/port notation (for example, 1/3)",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the FIS to which you want to bind interfaces.",
			},
			"ownernode": schema.Int64Attribute{
				// Cluster/show-only attribute; not a bind arg. Optional (no Computed).
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address.",
			},
		},
	}
}

func fis_interface_bindingGetThePayloadFromthePlan(ctx context.Context, data *FisInterfaceBindingResourceModel) network.Fisinterfacebinding {
	tflog.Debug(ctx, "In fis_interface_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// ownernode is a cluster/show-only attribute, NOT a bind arg; it is excluded from the payload.
	fis_interface_binding := network.Fisinterfacebinding{}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		fis_interface_binding.Ifnum = data.Ifnum.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		fis_interface_binding.Name = data.Name.ValueString()
	}

	return fis_interface_binding
}
