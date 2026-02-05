package systemfile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemfileResourceModel describes the resource data model.
type SystemfileResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Filecontent  types.String `tfsdk:"filecontent"`
	Fileencoding types.String `tfsdk:"fileencoding"`
	Filelocation types.String `tfsdk:"filelocation"`
	Filename     types.String `tfsdk:"filename"`
}

func (r *SystemfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemfile resource.",
			},
			"filecontent": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "file content in Base64 format.",
			},
			"fileencoding": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("BASE64"),
				Description: "encoding type of the file content.",
			},
			"filelocation": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "location of the file on Citrix ADC.",
			},
			"filename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the file. It should not include filepath.",
			},
		},
	}
}

func systemfileGetThePayloadFromtheConfig(ctx context.Context, data *SystemfileResourceModel) system.Systemfile {
	tflog.Debug(ctx, "In systemfileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemfile := system.Systemfile{}
	if !data.Filecontent.IsNull() {
		systemfile.Filecontent = data.Filecontent.ValueString()
	}
	if !data.Fileencoding.IsNull() {
		systemfile.Fileencoding = data.Fileencoding.ValueString()
	}
	if !data.Filelocation.IsNull() {
		systemfile.Filelocation = data.Filelocation.ValueString()
	}
	if !data.Filename.IsNull() {
		systemfile.Filename = data.Filename.ValueString()
	}

	return systemfile
}

func systemfileSetAttrFromGet(ctx context.Context, data *SystemfileResourceModel, getResponseData map[string]interface{}) *SystemfileResourceModel {
	tflog.Debug(ctx, "In systemfileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["filecontent"]; ok && val != nil {
		data.Filecontent = types.StringValue(val.(string))
	} else {
		data.Filecontent = types.StringNull()
	}
	if val, ok := getResponseData["fileencoding"]; ok && val != nil {
		data.Fileencoding = types.StringValue(val.(string))
	} else {
		data.Fileencoding = types.StringNull()
	}
	if val, ok := getResponseData["filelocation"]; ok && val != nil {
		data.Filelocation = types.StringValue(val.(string))
	} else {
		data.Filelocation = types.StringNull()
	}
	if val, ok := getResponseData["filename"]; ok && val != nil {
		data.Filename = types.StringValue(val.(string))
	} else {
		data.Filename = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Filename.ValueString(), data.Filelocation.ValueString()))

	return data
}
