package systemfile

import (
	"context"
	"encoding/base64"
	stdpath "path"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemfileResourceModel describes the resource data model.
// The model is shared between the resource and the datasource. It mirrors the
// legacy SDK v2 schema (filecontent/fileencoding/filelocation/filename plus the
// client-side-only is_base64_encoded flag) so existing state and configs keep
// working after the migration.
type SystemfileResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Filecontent     types.String `tfsdk:"filecontent"`
	Fileencoding    types.String `tfsdk:"fileencoding"`
	Filelocation    types.String `tfsdk:"filelocation"`
	Filename        types.String `tfsdk:"filename"`
	IsBase64Encoded types.Bool   `tfsdk:"is_base64_encoded"`
}

func (r *SystemfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemfile resource (the full path: filelocation/filename).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// SDK v2: Required, ForceNew, Sensitive
			"filecontent": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "file content in Base64 format.",
			},
			// SDK v2: Optional, ForceNew, Default "BASE64"
			"fileencoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("BASE64"),
				Description: "encoding type of the file content.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			// SDK v2: Required, ForceNew
			"filelocation": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "location of the file on Citrix ADC.",
			},
			// SDK v2: Required, ForceNew
			"filename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the file. It should not include filepath.",
			},
			// SDK v2: Optional, ForceNew, Default false. Client-side-only flag that
			// controls whether filecontent is treated as already base64-encoded.
			"is_base64_encoded": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Set to true when filecontent is already base64 encoded; otherwise the provider base64-encodes it before upload.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
		},
	}
}

// systemfileSetAttrFromGet populates the resource model from a NITRO GET
// response. It preserves the SDK v2 behavior: filecontent is returned base64
// encoded by NITRO, so it is decoded to plain text unless is_base64_encoded is
// set, in which case the raw base64 is retained. The is_base64_encoded flag is
// never returned by GET and must not be clobbered.
func systemfileSetAttrFromGet(ctx context.Context, data *SystemfileResourceModel, getResponseData map[string]interface{}) *SystemfileResourceModel {
	tflog.Debug(ctx, "In systemfileSetAttrFromGet Function")

	if val, ok := getResponseData["filecontent"]; ok && val != nil {
		raw := val.(string)
		if data.IsBase64Encoded.ValueBool() {
			// Original content was supplied already base64-encoded; keep it base64.
			data.Filecontent = types.StringValue(raw)
		} else {
			// Original content was plain text; decode the base64 back to plain text.
			if decoded, err := base64.StdEncoding.DecodeString(raw); err == nil {
				data.Filecontent = types.StringValue(string(decoded))
			} else {
				// Fall back to the raw value if it is not valid base64.
				data.Filecontent = types.StringValue(raw)
			}
		}
	}

	if val, ok := getResponseData["fileencoding"]; ok && val != nil {
		data.Fileencoding = types.StringValue(val.(string))
	} else if data.Fileencoding.IsNull() || data.Fileencoding.IsUnknown() {
		// NITRO may omit fileencoding; preserve the SDK v2 default rather than nulling.
		data.Fileencoding = types.StringValue("BASE64")
	}

	if val, ok := getResponseData["filelocation"]; ok && val != nil {
		data.Filelocation = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["filename"]; ok && val != nil {
		data.Filename = types.StringValue(val.(string))
	}

	// is_base64_encoded is a client-side-only flag never returned by GET.
	// Only default it when unknown/null; never clobber a configured value.
	if data.IsBase64Encoded.IsNull() || data.IsBase64Encoded.IsUnknown() {
		data.IsBase64Encoded = types.BoolValue(false)
	}

	// ID scheme matches SDK v2: path.Join(filelocation, filename).
	data.Id = types.StringValue(stdpath.Join(data.Filelocation.ValueString(), data.Filename.ValueString()))

	return data
}

// systemfileSetAttrFromGetForDatasource populates the shared model from a NITRO
// GET response for datasource reads. The datasource has no is_base64_encoded
// input, so filecontent is stored as returned by NITRO (base64). All fields are
// copied and the ID is set.
func systemfileSetAttrFromGetForDatasource(ctx context.Context, data *SystemfileResourceModel, getResponseData map[string]interface{}) *SystemfileResourceModel {
	tflog.Debug(ctx, "In systemfileSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["filecontent"]; ok && val != nil {
		data.Filecontent = types.StringValue(val.(string))
	}

	if val, ok := getResponseData["fileencoding"]; ok && val != nil {
		data.Fileencoding = types.StringValue(val.(string))
	} else {
		data.Fileencoding = types.StringValue("BASE64")
	}

	if val, ok := getResponseData["filelocation"]; ok && val != nil {
		data.Filelocation = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["filename"]; ok && val != nil {
		data.Filename = types.StringValue(val.(string))
	}

	if data.IsBase64Encoded.IsNull() || data.IsBase64Encoded.IsUnknown() {
		data.IsBase64Encoded = types.BoolValue(false)
	}

	data.Id = types.StringValue(stdpath.Join(data.Filelocation.ValueString(), data.Filename.ValueString()))

	return data
}
