package systemcpuparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemcpuparamResourceModel describes the resource data model.
type SystemcpuparamResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Pemode types.String `tfsdk:"pemode"`
}

func (r *SystemcpuparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemcpuparam resource.",
			},
			"pemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set PEmode to DEFAULT/CPUBOUND. Distribute the PE weights equally if PEmode is set to CPUBOUND.",
			},
		},
	}
}

func systemcpuparamGetThePayloadFromthePlan(ctx context.Context, data *SystemcpuparamResourceModel) system.Systemcpuparam {
	tflog.Debug(ctx, "In systemcpuparamGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemcpuparam := system.Systemcpuparam{}
	if !data.Pemode.IsNull() && !data.Pemode.IsUnknown() {
		systemcpuparam.Pemode = data.Pemode.ValueString()
	}

	return systemcpuparam
}

// systemcpuparamSetAttrFromGet populates the resource model from the GET response.
// This is a settable singleton (PUT set / GET): pemode is Optional+Computed and the
// GET (get-all) response always echoes the server-applied value (or default), so we
// faithfully copy it. The synthetic ID is set exactly once in Create (Pattern 6), so
// it is NOT recomputed here.
func systemcpuparamSetAttrFromGet(ctx context.Context, data *SystemcpuparamResourceModel, getResponseData map[string]interface{}) *SystemcpuparamResourceModel {
	tflog.Debug(ctx, "In systemcpuparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["pemode"]; ok && val != nil {
		data.Pemode = types.StringValue(val.(string))
	}

	return data
}

// systemcpuparamSetAttrFromGetForDatasource faithfully copies every field from the GET
// response and sets the synthetic ID, because the datasource never calls Create
// (Pattern 7).
func systemcpuparamSetAttrFromGetForDatasource(ctx context.Context, data *SystemcpuparamResourceModel, getResponseData map[string]interface{}) *SystemcpuparamResourceModel {
	tflog.Debug(ctx, "In systemcpuparamSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["pemode"]; ok && val != nil {
		data.Pemode = types.StringValue(val.(string))
	} else {
		data.Pemode = types.StringNull()
	}

	// Datasource has no Create, so set the synthetic ID here.
	data.Id = types.StringValue("systemcpuparam-config")

	return data
}
