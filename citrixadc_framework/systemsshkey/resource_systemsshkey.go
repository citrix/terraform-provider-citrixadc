package systemsshkey

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemsshkeyResource{}
var _ resource.ResourceWithConfigure = (*SystemsshkeyResource)(nil)
var _ resource.ResourceWithImportState = (*SystemsshkeyResource)(nil)

func NewSystemsshkeyResource() resource.Resource {
	return &SystemsshkeyResource{}
}

// SystemsshkeyResource defines the resource implementation.
type SystemsshkeyResource struct {
	client *service.NitroClient
}

func (r *SystemsshkeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemsshkeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsshkey"
}

func (r *SystemsshkeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemsshkeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemsshkeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemsshkey resource")
	systemsshkey := systemsshkeyGetThePayloadFromthePlan(ctx, &data)

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
	r.readSystemsshkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemsshkeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemsshkeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemsshkey resource")

	r.readSystemsshkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemsshkeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemsshkeyResourceModel

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
	r.readSystemsshkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemsshkeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemsshkeyResourceModel

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
func (r *SystemsshkeyResource) readSystemsshkeyFromApi(ctx context.Context, data *SystemsshkeyResourceModel, diags *diag.Diagnostics) {

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

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "systemsshkey returned empty array")
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

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "systemsshkey not found with the provided ID attributes")
		return
	}

	systemsshkeySetAttrFromGet(ctx, data, dataArr[foundIndex])
}
