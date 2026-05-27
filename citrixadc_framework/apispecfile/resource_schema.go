package apispecfile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/api"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ApispecfileResourceModel describes the resource data model.
type ApispecfileResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`
}

func (r *ApispecfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the apispecfile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported spec file. Must begin with an ASCII alphanumeric or underscore(_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrite any existing schema file of the same name.",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL specifying the protocol, host, and path, including file name, to the spec file to be imported. For example, http://www.example.com/spec_file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func apispecfileGetThePayloadFromthePlan(ctx context.Context, data *ApispecfileResourceModel) api.Apispecfile {
	tflog.Debug(ctx, "In apispecfileGetThePayloadFromthePlan Function")

	apispecfile := api.Apispecfile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		apispecfile.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		apispecfile.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		apispecfile.Src = data.Src.ValueString()
	}

	return apispecfile
}

func apispecfileSetAttrFromGet(ctx context.Context, data *ApispecfileResourceModel, getResponseData map[string]interface{}) *ApispecfileResourceModel {
	tflog.Debug(ctx, "In apispecfileSetAttrFromGet Function")

	// Convert API response to model.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// "overwrite" is a write-only Import input. The GET response never echoes it,
	// so do NOT touch data.Overwrite here — preserve the prior plan/state value to
	// avoid a perpetual diff.
	// "src" is the URL supplied at Import time; the NITRO server normalizes it
	// (e.g., strips the "local://" protocol prefix) and returns the bare filename.
	// Preserve the user-configured value to avoid an "inconsistent result after
	// apply" error from Terraform.

	// NOTE: Do not recompute data.Id here. The resource Create sets the ID once
	// from the user-supplied name, and the datasource Read sets it explicitly
	// after calling this helper.

	return data
}
