package clusternodegroup_authenticationvserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ClusternodegroupAuthenticationvserverBindingResourceModel describes the resource data model.
type ClusternodegroupAuthenticationvserverBindingResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Vserver types.String `tfsdk:"vserver"`
}

func (r *ClusternodegroupAuthenticationvserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusternodegroup_authenticationvserver_binding resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
			"vserver": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "vserver that need to be bound to this nodegroup.",
			},
		},
	}
}

func clusternodegroup_authenticationvserver_bindingGetThePayloadFromthePlan(ctx context.Context, data *ClusternodegroupAuthenticationvserverBindingResourceModel) cluster.Clusternodegroupauthenticationvserverbinding {
	tflog.Debug(ctx, "In clusternodegroup_authenticationvserver_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	clusternodegroup_authenticationvserver_binding := cluster.Clusternodegroupauthenticationvserverbinding{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		clusternodegroup_authenticationvserver_binding.Name = data.Name.ValueString()
	}
	if !data.Vserver.IsNull() && !data.Vserver.IsUnknown() {
		clusternodegroup_authenticationvserver_binding.Vserver = data.Vserver.ValueString()
	}

	return clusternodegroup_authenticationvserver_binding
}

// clusternodegroup_authenticationvserver_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied values (name, vserver are both RequiresReplace identity attrs) and
// does NOT recompute the ID, which is set exactly once in Create.
func clusternodegroup_authenticationvserver_bindingSetAttrFromGet(ctx context.Context, data *ClusternodegroupAuthenticationvserverBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupAuthenticationvserverBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_authenticationvserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	}

	return data
}

// clusternodegroup_authenticationvserver_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and composes the ID, since the datasource has no Create to seed those values.
func clusternodegroup_authenticationvserver_bindingSetAttrFromGetForDatasource(ctx context.Context, data *ClusternodegroupAuthenticationvserverBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupAuthenticationvserverBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_authenticationvserver_bindingSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	} else {
		data.Vserver = types.StringNull()
	}

	// Set ID for the datasource using the new Framework key:value composite format
	// (name:<v>,vserver:<v>) to stay consistent with the resource's Create ID.
	data.Id = types.StringValue(fmt.Sprintf("name:%s,vserver:%s", utils.UrlEncode(data.Name.ValueString()), utils.UrlEncode(data.Vserver.ValueString())))

	return data
}
