package systemsshkey

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemsshkeyImportResource{}
var _ resource.ResourceWithConfigure = (*SystemsshkeyImportResource)(nil)
var _ resource.ResourceWithImportState = (*SystemsshkeyImportResource)(nil)

func NewSystemsshkeyImportResource() resource.Resource {
	return &SystemsshkeyImportResource{}
}

// SystemsshkeyImportResource defines the resource implementation.
type SystemsshkeyImportResource struct {
	client *service.NitroClient
}

// SystemsshkeyImportResourceModel describes the resource data model.
type SystemsshkeyImportResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Src        types.String `tfsdk:"src"`
	Sshkeytype types.String `tfsdk:"sshkeytype"`
}

func (r *SystemsshkeyImportResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemsshkeyImportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsshkey_import"
}

func (r *SystemsshkeyImportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemsshkeyImportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *SystemsshkeyImportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemsshkeyImportResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemsshkey resource")
	systemsshkey := systemsshkeyImportGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes only an "Import" action (POST ?action=Import, capital I) for
	// create; there is no add verb. Verb casing matters.
	err := r.client.ActOnResource(service.Systemsshkey.Type(), &systemsshkey, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemsshkey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemsshkey resource")

	// Set composite ID exactly once here (Pattern 6)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("sshkeytype:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Sshkeytype.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSystemsshkeyImportFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemsshkeyImportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemsshkeyImportResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemsshkey resource")

	r.readSystemsshkeyImportFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource has been deleted out-of-band - remove from state
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemsshkeyImportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemsshkeyImportResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for systemsshkey; NITRO exposes no update endpoint and all
	// attributes are RequiresReplace (Pattern 5).
	tflog.Debug(ctx, "Update is a no-op for systemsshkey; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSystemsshkeyImportFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemsshkeyImportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemsshkeyImportResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemsshkey resource")

	// Delete key = name; delete arg = sshkeytype (mandatory).
	args := []string{
		fmt.Sprintf("sshkeytype:%s", utils.UrlEncode(data.Sshkeytype.ValueString())),
	}

	err := r.client.DeleteResourceWithArgs(service.Systemsshkey.Type(), data.Name.ValueString(), args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemsshkey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systemsshkey resource")
}

// Helper function to read systemsshkey data from API.
// NITRO exposes only get (all) (no get-by-name); fetch all records and match on
// name + sshkeytype.
func (r *SystemsshkeyImportResource) readSystemsshkeyImportFromApi(ctx context.Context, data *SystemsshkeyImportResourceModel, diags *diag.Diagnostics) {

	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "sshkeytype"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Systemsshkey.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemsshkey, got error: %s", err))
		return
	}

	// Resource is missing - signal removal from state
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check name
		if idVal, ok := idMap["name"]; ok {
			if val, ok := v["name"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check sshkeytype
		if idVal, ok := idMap["sshkeytype"]; ok {
			if val, ok := v["sshkeytype"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing - signal removal from state
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	systemsshkeyImportSetAttrFromGet(ctx, data, dataArr[foundIndex])
}

func systemsshkeyImportGetThePayloadFromthePlan(ctx context.Context, data *SystemsshkeyImportResourceModel) system.Systemsshkey {
	tflog.Debug(ctx, "In systemsshkeyImportGetThePayloadFromthePlan Function")

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

// systemsshkeyImportSetAttrFromGet populates the resource model from a GET response.
// Note: src is a write-only-ish import source that GET does not return; preserve
// the existing plan/state value for it (Pattern 7). The ID is NOT recomputed here;
// it is set exactly once in Create (Pattern 6).
func systemsshkeyImportSetAttrFromGet(ctx context.Context, data *SystemsshkeyImportResourceModel, getResponseData map[string]interface{}) *SystemsshkeyImportResourceModel {
	tflog.Debug(ctx, "In systemsshkeyImportSetAttrFromGet Function")

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
