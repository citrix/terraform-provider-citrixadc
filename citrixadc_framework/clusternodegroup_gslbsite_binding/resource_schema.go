package clusternodegroup_gslbsite_binding

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

// ClusternodegroupGslbsiteBindingResourceModel describes the resource data model.
type ClusternodegroupGslbsiteBindingResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Gslbsite types.String `tfsdk:"gslbsite"`
	Name     types.String `tfsdk:"name"`
}

func (r *ClusternodegroupGslbsiteBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusternodegroup_gslbsite_binding resource.",
			},
			"gslbsite": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "vserver that need to be bound to this nodegroup.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
		},
	}
}

func clusternodegroup_gslbsite_bindingGetThePayloadFromthePlan(ctx context.Context, data *ClusternodegroupGslbsiteBindingResourceModel) cluster.Clusternodegroupgslbsitebinding {
	tflog.Debug(ctx, "In clusternodegroup_gslbsite_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	clusternodegroup_gslbsite_binding := cluster.Clusternodegroupgslbsitebinding{}
	if !data.Gslbsite.IsNull() && !data.Gslbsite.IsUnknown() {
		clusternodegroup_gslbsite_binding.Gslbsite = data.Gslbsite.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		clusternodegroup_gslbsite_binding.Name = data.Name.ValueString()
	}

	return clusternodegroup_gslbsite_binding
}

func clusternodegroup_gslbsite_bindingSetAttrFromGet(ctx context.Context, data *ClusternodegroupGslbsiteBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupGslbsiteBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_gslbsite_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gslbsite"]; ok && val != nil {
		data.Gslbsite = types.StringValue(val.(string))
	} else {
		data.Gslbsite = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	// (legacy SDK v2 order: name,gslbsite)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("gslbsite:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Gslbsite.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
