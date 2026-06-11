package policypatsetfile

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PolicypatsetfileResource{}
var _ resource.ResourceWithConfigure = (*PolicypatsetfileResource)(nil)
var _ resource.ResourceWithImportState = (*PolicypatsetfileResource)(nil)

func NewPolicypatsetfileResource() resource.Resource {
	return &PolicypatsetfileResource{}
}

// PolicypatsetfileResource defines the resource implementation.
type PolicypatsetfileResource struct {
	client *service.NitroClient
}

func (r *PolicypatsetfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicypatsetfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policypatsetfile"
}

func (r *PolicypatsetfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicypatsetfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicypatsetfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policypatsetfile resource")
	policypatsetfile := policypatsetfileGetThePayloadFromthePlan(ctx, &data)

	// Create via the NITRO Import action (POST ?action=Import, capital "I").
	err := r.client.ActOnResource(service.Policypatsetfile.Type(), &policypatsetfile, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policypatsetfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created policypatsetfile resource")

	// Set ID for the resource before reading state (single key: name)
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	r.readPolicypatsetfileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicypatsetfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policypatsetfile resource")

	r.readPolicypatsetfileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state PolicypatsetfileResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for policypatsetfile; all write attributes are RequiresReplace")

	// Read the current state back
	r.readPolicypatsetfileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicypatsetfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policypatsetfile resource")

	// Plain DELETE /policypatsetfile/<name>; no ?action, no args.
	// NITRO quirk: this DELETE actually removes the imported patset file but
	// reports a spurious errorcode 258 "No such resource [name, ...]". Treat
	// that specific response as a successful delete.
	err := r.client.DeleteResource(service.Policypatsetfile.Type(), data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "No such resource") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete policypatsetfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted policypatsetfile resource")
}

// Helper function to read policypatsetfile data from API
func (r *PolicypatsetfileResource) readPolicypatsetfileFromApi(ctx context.Context, data *PolicypatsetfileResourceModel, diags *diag.Diagnostics) {
	// An imported patset file is NOT retrievable via a plain GET
	// /policypatsetfile/<name> (that returns nothing). It is only listed by the
	// filtered GET /policypatsetfile?args=imported:true (Pattern 15: "imported"
	// is a GET-only filter param). Fetch that list and match by name.
	name := data.Id.ValueString()

	getResponseData, err := findImportedPolicypatsetfileByName(r.client, name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policypatsetfile, got error: %s", err))
		return
	}

	if getResponseData == nil {
		diags.AddError("Client Error", fmt.Sprintf("policypatsetfile %s not found.", name))
		return
	}

	policypatsetfileSetAttrFromGet(ctx, data, getResponseData)
}

// findImportedPolicypatsetfileByName lists imported patset files
// (GET /policypatsetfile?args=imported:true) and returns the entry whose "name"
// matches the supplied name, or (nil, nil) if no such entry exists.
func findImportedPolicypatsetfileByName(client *service.NitroClient, name string) (map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType: service.Policypatsetfile.Type(),
		ArgsMap:      map[string]string{"imported": "true"},
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}
	for _, item := range dataArr {
		if val, ok := item["name"]; ok && val != nil {
			if val.(string) == name {
				return item, nil
			}
		}
	}
	return nil, nil
}
