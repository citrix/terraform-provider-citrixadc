package clusternodegroup_streamidentifier_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ClusternodegroupStreamidentifierBindingResourceModel describes the resource data model.
type ClusternodegroupStreamidentifierBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Identifiername types.String `tfsdk:"identifiername"`
	Name           types.String `tfsdk:"name"`
}

func (r *ClusternodegroupStreamidentifierBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusternodegroup_streamidentifier_binding resource.",
			},
			"identifiername": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "stream identifier  and rate limit identifier that need to be bound to this nodegroup.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the nodegroup to which you want to bind a cluster node or an entity.",
			},
		},
	}
}

func clusternodegroup_streamidentifier_bindingGetThePayloadFromthePlan(ctx context.Context, data *ClusternodegroupStreamidentifierBindingResourceModel) cluster.Clusternodegroupstreamidentifierbinding {
	tflog.Debug(ctx, "In clusternodegroup_streamidentifier_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	clusternodegroup_streamidentifier_binding := cluster.Clusternodegroupstreamidentifierbinding{}
	if !data.Identifiername.IsNull() && !data.Identifiername.IsUnknown() {
		clusternodegroup_streamidentifier_binding.Identifiername = data.Identifiername.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		clusternodegroup_streamidentifier_binding.Name = data.Name.ValueString()
	}

	return clusternodegroup_streamidentifier_binding
}

// clusternodegroup_streamidentifier_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied values (name, identifiername are both RequiresReplace identity attrs) and
// does NOT recompute the ID, which is set exactly once in Create.
func clusternodegroup_streamidentifier_bindingSetAttrFromGet(ctx context.Context, data *ClusternodegroupStreamidentifierBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupStreamidentifierBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_streamidentifier_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["identifiername"]; ok && val != nil {
		data.Identifiername = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}

	return data
}

// clusternodegroup_streamidentifier_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and composes the ID, since the datasource has no Create to seed those values.
func clusternodegroup_streamidentifier_bindingSetAttrFromGetForDatasource(ctx context.Context, data *ClusternodegroupStreamidentifierBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupStreamidentifierBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_streamidentifier_bindingSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["identifiername"]; ok && val != nil {
		data.Identifiername = types.StringValue(val.(string))
	} else {
		data.Identifiername = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("identifiername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Identifiername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
