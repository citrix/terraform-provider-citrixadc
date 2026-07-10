package nshmackey

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NshmackeyResourceModel describes the resource data model.
type NshmackeyResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Comment           types.String `tfsdk:"comment"`
	Digest            types.String `tfsdk:"digest"`
	Keyvalue          types.String `tfsdk:"keyvalue"`
	KeyvalueWo        types.String `tfsdk:"keyvalue_wo"`
	KeyvalueWoVersion types.Int64  `tfsdk:"keyvalue_wo_version"`
	Name              types.String `tfsdk:"name"`
}

func (r *NshmackeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nshmackey resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this encryption key.",
			},
			"digest": schema.StringAttribute{
				Required:    true,
				Description: "Digest (hash) function to be used in the HMAC computation.",
			},
			"keyvalue": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "The hex-encoded key to be used in the HMAC computation. The key can be any length (up to a Citrix ADC-imposed maximum of 255 bytes). If the length is less than the digest block size, it will be zero padded up to the block size. If it is greater than the block size, it will be hashed using the digest function to the block size. The block size for each digest is:\n   MD2    - 16 bytes\n   MD4    - 16 bytes\n   MD5    - 16 bytes\n   SHA1   - 20 bytes\n   SHA224 - 28 bytes\n   SHA256 - 32 bytes\n   SHA384 - 48 bytes\n   SHA512 - 64 bytes\nNote that the key will be encrypted when it it is saved\n\nThere is a special key value AUTO which generates a new random key for the specified digest. This kind of key is\nintended for use cases where the NetScaler both generates and verifies an HMAC on  the same data.",
			},
			"keyvalue_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "The hex-encoded key to be used in the HMAC computation. The key can be any length (up to a Citrix ADC-imposed maximum of 255 bytes). If the length is less than the digest block size, it will be zero padded up to the block size. If it is greater than the block size, it will be hashed using the digest function to the block size. The block size for each digest is:\n   MD2    - 16 bytes\n   MD4    - 16 bytes\n   MD5    - 16 bytes\n   SHA1   - 20 bytes\n   SHA224 - 28 bytes\n   SHA256 - 32 bytes\n   SHA384 - 48 bytes\n   SHA512 - 64 bytes\nNote that the key will be encrypted when it it is saved\n\nThere is a special key value AUTO which generates a new random key for the specified digest. This kind of key is\nintended for use cases where the NetScaler both generates and verifies an HMAC on  the same data.",
			},
			"keyvalue_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a keyvalue_wo update.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Key name.  This follows the same syntax rules as other expression entity names:\n   It must begin with an alpha character (A-Z or a-z) or an underscore (_).\n   The rest of the characters must be alpha, numeric (0-9) or underscores.\n   It cannot be re or xp (reserved for regular and XPath expressions).\n   It cannot be an expression reserved word (e.g. SYS or HTTP).\n   It cannot be used for an existing expression object (HTTP callout, patset, dataset, stringmap, or named expression).",
			},
		},
	}
}

func nshmackeyGetThePayloadFromthePlan(ctx context.Context, data *NshmackeyResourceModel) ns.Nshmackey {
	tflog.Debug(ctx, "In nshmackeyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nshmackey := ns.Nshmackey{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		nshmackey.Comment = data.Comment.ValueString()
	}
	if !data.Digest.IsNull() && !data.Digest.IsUnknown() {
		nshmackey.Digest = data.Digest.ValueString()
	}
	if !data.Keyvalue.IsNull() && !data.Keyvalue.IsUnknown() {
		nshmackey.Keyvalue = data.Keyvalue.ValueString()
	}
	// Skip write-only attribute: keyvalue_wo
	// Skip version tracker attribute: keyvalue_wo_version
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nshmackey.Name = data.Name.ValueString()
	}

	return nshmackey
}

func nshmackeyGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *NshmackeyResourceModel) ns.Nshmackey {
	tflog.Debug(ctx, "In nshmackeyGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model
	nshmackey := ns.Nshmackey{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		nshmackey.Comment = data.Comment.ValueString()
	}
	if !data.Digest.IsNull() && !data.Digest.IsUnknown() {
		nshmackey.Digest = data.Digest.ValueString()
	}
	if !data.Keyvalue.IsNull() && !data.Keyvalue.IsUnknown() {
		nshmackey.Keyvalue = data.Keyvalue.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nshmackey.Name = data.Name.ValueString()
	}

	return nshmackey
}

func nshmackeyGetThePayloadFromtheConfig(ctx context.Context, data *NshmackeyResourceModel, payload *ns.Nshmackey) {
	tflog.Debug(ctx, "In nshmackeyGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: keyvalue_wo -> keyvalue
	if !data.KeyvalueWo.IsNull() {
		keyvalueWo := data.KeyvalueWo.ValueString()
		if keyvalueWo != "" {
			payload.Keyvalue = keyvalueWo
		}
	}
}

func nshmackeySetAttrFromGet(ctx context.Context, data *NshmackeyResourceModel, getResponseData map[string]interface{}) *NshmackeyResourceModel {
	tflog.Debug(ctx, "In nshmackeySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["digest"]; ok && val != nil {
		data.Digest = types.StringValue(val.(string))
	} else {
		data.Digest = types.StringNull()
	}
	// keyvalue is not returned by NITRO API (secret/ephemeral) - retain from config
	// keyvalue_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// keyvalue_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
