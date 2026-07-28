package metricsprofile_csvserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// MetricsprofileCsvserverBindingResourceModel describes the resource data model.
type MetricsprofileCsvserverBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Entityname types.String `tfsdk:"entityname"`
	Entitytype types.String `tfsdk:"entitytype"`
	Name       types.String `tfsdk:"name"`
}

func (r *MetricsprofileCsvserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the metricsprofile_csvserver_binding resource.",
			},
			"entityname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the entity bound to the metrics profile.",
			},
			"entitytype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of the entity bound to the metrics profile.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the metrics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.!",
			},
		},
	}
}

func metricsprofile_csvserver_bindingGetThePayloadFromthePlan(ctx context.Context, data *MetricsprofileCsvserverBindingResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In metricsprofile_csvserver_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	metricsprofile_csvserver_binding := make(map[string]interface{})
	if !data.Entityname.IsNull() && !data.Entityname.IsUnknown() {
		metricsprofile_csvserver_binding["entityname"] = data.Entityname.ValueString()
	}
	if !data.Entitytype.IsNull() && !data.Entitytype.IsUnknown() {
		metricsprofile_csvserver_binding["entitytype"] = data.Entitytype.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		metricsprofile_csvserver_binding["name"] = data.Name.ValueString()
	}

	return metricsprofile_csvserver_binding
}

func metricsprofile_csvserver_bindingSetAttrFromGet(ctx context.Context, data *MetricsprofileCsvserverBindingResourceModel, getResponseData map[string]interface{}) *MetricsprofileCsvserverBindingResourceModel {
	tflog.Debug(ctx, "In metricsprofile_csvserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["entityname"]; ok && val != nil {
		data.Entityname = types.StringValue(val.(string))
	} else {
		data.Entityname = types.StringNull()
	}
	if val, ok := getResponseData["entitytype"]; ok && val != nil {
		data.Entitytype = types.StringValue(val.(string))
	} else {
		data.Entitytype = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("entityname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Entityname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("entitytype:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Entitytype.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
