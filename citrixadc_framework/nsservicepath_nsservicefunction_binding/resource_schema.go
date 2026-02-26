package nsservicepath_nsservicefunction_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsservicepathNsservicefunctionBindingResourceModel describes the resource data model.
type NsservicepathNsservicefunctionBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Index           types.Int64  `tfsdk:"index"`
	Servicefunction types.String `tfsdk:"servicefunction"`
	Servicepathname types.String `tfsdk:"servicepathname"`
}

func (r *NsservicepathNsservicefunctionBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsservicepath_nsservicefunction_binding resource.",
			},
			"index": schema.Int64Attribute{
				Required:    true,
				Description: "The serviceindex of each servicefunction in path.",
			},
			"servicefunction": schema.StringAttribute{
				Required:    true,
				Description: "List of service functions constituting the chain.",
			},
			"servicepathname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Service path. Must begin with an ASCII alphanumeric or underscore (_) character, and must\n      contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-)\n      characters.",
			},
		},
	}
}

func nsservicepath_nsservicefunction_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NsservicepathNsservicefunctionBindingResourceModel) ns.Nsservicepathnsservicefunctionbinding {
	tflog.Debug(ctx, "In nsservicepath_nsservicefunction_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsservicepath_nsservicefunction_binding := ns.Nsservicepathnsservicefunctionbinding{}
	if !data.Index.IsNull() {
		nsservicepath_nsservicefunction_binding.Index = utils.IntPtr(int(data.Index.ValueInt64()))
	}
	if !data.Servicefunction.IsNull() {
		nsservicepath_nsservicefunction_binding.Servicefunction = data.Servicefunction.ValueString()
	}
	if !data.Servicepathname.IsNull() {
		nsservicepath_nsservicefunction_binding.Servicepathname = data.Servicepathname.ValueString()
	}

	return nsservicepath_nsservicefunction_binding
}

func nsservicepath_nsservicefunction_bindingSetAttrFromGet(ctx context.Context, data *NsservicepathNsservicefunctionBindingResourceModel, getResponseData map[string]interface{}) *NsservicepathNsservicefunctionBindingResourceModel {
	tflog.Debug(ctx, "In nsservicepath_nsservicefunction_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["index"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Index = types.Int64Value(intVal)
		}
	} else {
		data.Index = types.Int64Null()
	}
	if val, ok := getResponseData["servicefunction"]; ok && val != nil {
		data.Servicefunction = types.StringValue(val.(string))
	} else {
		data.Servicefunction = types.StringNull()
	}
	if val, ok := getResponseData["servicepathname"]; ok && val != nil {
		data.Servicepathname = types.StringValue(val.(string))
	} else {
		data.Servicepathname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("servicefunction:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Servicefunction.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicepathname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Servicepathname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
