package crvserver_analyticsprofile_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cr"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CrvserverAnalyticsprofileBindingResourceModel describes the resource data model.
type CrvserverAnalyticsprofileBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Analyticsprofile types.String `tfsdk:"analyticsprofile"`
	Name             types.String `tfsdk:"name"`
}

func (r *CrvserverAnalyticsprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the crvserver_analyticsprofile_binding resource.",
			},
			"analyticsprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the analytics profile bound to the CR vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the cache redirection virtual server to which to bind the cache redirection policy.",
			},
		},
	}
}

func crvserver_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx context.Context, data *CrvserverAnalyticsprofileBindingResourceModel) cr.Crvserveranalyticsprofilebinding {
	tflog.Debug(ctx, "In crvserver_analyticsprofile_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	crvserver_analyticsprofile_binding := cr.Crvserveranalyticsprofilebinding{}
	if !data.Analyticsprofile.IsNull() {
		crvserver_analyticsprofile_binding.Analyticsprofile = data.Analyticsprofile.ValueString()
	}
	if !data.Name.IsNull() {
		crvserver_analyticsprofile_binding.Name = data.Name.ValueString()
	}

	return crvserver_analyticsprofile_binding
}

func crvserver_analyticsprofile_bindingSetAttrFromGet(ctx context.Context, data *CrvserverAnalyticsprofileBindingResourceModel, getResponseData map[string]interface{}) *CrvserverAnalyticsprofileBindingResourceModel {
	tflog.Debug(ctx, "In crvserver_analyticsprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["analyticsprofile"]; ok && val != nil {
		data.Analyticsprofile = types.StringValue(val.(string))
	} else {
		data.Analyticsprofile = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("analyticsprofile:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Analyticsprofile.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
