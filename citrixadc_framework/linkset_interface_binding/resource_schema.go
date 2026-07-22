package linkset_interface_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LinksetInterfaceBindingResourceModel describes the resource data model.
//
// The NITRO parent key for this binding is literally named "id" (the linkset name).
// Because Terraform reserves the "id" attribute for the synthetic resource
// identifier, the parent key is exposed here as "linksetid" while the composite
// resource ID string still uses the "id" key (id:<linkset>,ifnum:<intf>).
type LinksetInterfaceBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Linksetid types.String `tfsdk:"linksetid"`
	Ifnum     types.String `tfsdk:"ifnum"`
}

func (r *LinksetInterfaceBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the linkset_interface_binding resource.",
			},
			"linksetid": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ID of the linkset to which to bind the interfaces.",
			},
			"ifnum": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The interfaces to be bound to the linkset.",
			},
		},
	}
}

func linkset_interface_bindingGetThePayloadFromthePlan(ctx context.Context, data *LinksetInterfaceBindingResourceModel) network.Linksetinterfacebinding {
	tflog.Debug(ctx, "In linkset_interface_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	linkset_interface_binding := network.Linksetinterfacebinding{}
	if !data.Linksetid.IsNull() && !data.Linksetid.IsUnknown() {
		linkset_interface_binding.Id = data.Linksetid.ValueString()
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		linkset_interface_binding.Ifnum = data.Ifnum.ValueString()
	}

	return linkset_interface_binding
}

// linkset_interface_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied values (linksetid, ifnum are RequiresReplace
// identity attrs) and does NOT recompute the ID, which is set exactly once in Create.
func linkset_interface_bindingSetAttrFromGet(ctx context.Context, data *LinksetInterfaceBindingResourceModel, getResponseData map[string]interface{}) *LinksetInterfaceBindingResourceModel {
	tflog.Debug(ctx, "In linkset_interface_bindingSetAttrFromGet Function")

	// linksetid (parent key) and ifnum are RequiresReplace identity attrs supplied
	// by the plan; preserve them. Only adopt the GET value when missing (import).
	if val, ok := getResponseData["id"]; ok && val != nil {
		if data.Linksetid.IsNull() || data.Linksetid.ValueString() == "" {
			data.Linksetid = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		if data.Ifnum.IsNull() || data.Ifnum.ValueString() == "" {
			data.Ifnum = types.StringValue(val.(string))
		}
	}

	return data
}

// linkset_interface_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and composes the ID, since the datasource has no Create to
// seed those values.
func linkset_interface_bindingSetAttrFromGetForDatasource(ctx context.Context, data *LinksetInterfaceBindingResourceModel, getResponseData map[string]interface{}) *LinksetInterfaceBindingResourceModel {
	tflog.Debug(ctx, "In linkset_interface_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		data.Linksetid = types.StringValue(val.(string))
	} else {
		data.Linksetid = types.StringNull()
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	} else {
		data.Ifnum = types.StringNull()
	}

	// Set ID for the datasource. Composite key id:<linkset>,ifnum:<intf>.
	data.Id = types.StringValue(linkset_interface_bindingComposeId(data.Linksetid.ValueString(), data.Ifnum.ValueString()))

	return data
}

// linkset_interface_bindingComposeId builds the composite resource ID string.
// Format: id:<linkset>,ifnum:<interface> (both UrlEncoded; linkset and interface
// ids contain '/').
func linkset_interface_bindingComposeId(linksetid, ifnum string) string {
	return "id:" + utils.UrlEncode(linksetid) + ",ifnum:" + utils.UrlEncode(ifnum)
}
