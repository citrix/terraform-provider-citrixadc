package systemsshkey

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SystemsshkeyResourceModel describes the resource data model.
type SystemsshkeyResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Src        types.String `tfsdk:"src"`
	Sshkeytype types.String `tfsdk:"sshkeytype"`
}

func (r *SystemsshkeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemsshkey resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL \\(protocol, host, path, and file name\\) from where the location file will be imported.\n            NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL \\(protocol, host, path, and file name\\) from where the location file will be imported.\n            NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
			"sshkeytype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The type of the ssh key whether public or private key",
			},
		},
	}
}

func systemsshkeyGetThePayloadFromthePlan(ctx context.Context, data *SystemsshkeyResourceModel) system.Systemsshkey {
	tflog.Debug(ctx, "In systemsshkeyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	// Note: _nextgenapiresource is read-only and excluded from the Import payload (Pattern 15)
	systemsshkey := system.Systemsshkey{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		systemsshkey.Name = data.Name.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		systemsshkey.Src = data.Src.ValueString()
	}
	if !data.Sshkeytype.IsNull() && !data.Sshkeytype.IsUnknown() {
		systemsshkey.Sshkeytype = data.Sshkeytype.ValueString()
	}

	return systemsshkey
}

// systemsshkeySetAttrFromGet populates the resource model from a GET response.
// Note: src is a write-only-ish import source that GET does not return; preserve
// the existing plan/state value for it (Pattern 7). The ID is NOT recomputed here;
// it is set exactly once in Create (Pattern 6).
func systemsshkeySetAttrFromGet(ctx context.Context, data *SystemsshkeyResourceModel, getResponseData map[string]interface{}) *SystemsshkeyResourceModel {
	tflog.Debug(ctx, "In systemsshkeySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["sshkeytype"]; ok && val != nil {
		data.Sshkeytype = types.StringValue(val.(string))
	}
	// src is not returned by GET; preserve the configured value (do not touch).

	return data
}

// systemsshkeySetAttrFromGetForDatasource faithfully copies the GET response into
// the model for the datasource flow (which has no prior plan/state to preserve)
// and sets the composite ID. src is not returned by GET, so it remains null.
func systemsshkeySetAttrFromGetForDatasource(ctx context.Context, data *SystemsshkeyResourceModel, getResponseData map[string]interface{}) *SystemsshkeyResourceModel {
	tflog.Debug(ctx, "In systemsshkeySetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["sshkeytype"]; ok && val != nil {
		data.Sshkeytype = types.StringValue(val.(string))
	} else {
		data.Sshkeytype = types.StringNull()
	}
	// src is not returned by GET.
	data.Src = types.StringNull()

	// Set composite ID (datasource has no Create)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("sshkeytype:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Sshkeytype.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
