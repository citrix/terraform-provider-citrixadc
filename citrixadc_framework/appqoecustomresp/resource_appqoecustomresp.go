package appqoecustomresp

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppqoecustomrespResource{}
var _ resource.ResourceWithConfigure = (*AppqoecustomrespResource)(nil)
var _ resource.ResourceWithImportState = (*AppqoecustomrespResource)(nil)

func NewAppqoecustomrespResource() resource.Resource {
	return &AppqoecustomrespResource{}
}

// AppqoecustomrespResource defines the resource implementation.
type AppqoecustomrespResource struct {
	client *service.NitroClient
}

func (r *AppqoecustomrespResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppqoecustomrespResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appqoecustomresp"
}

func (r *AppqoecustomrespResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppqoecustomrespResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppqoecustomrespResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appqoecustomresp resource")

	appqoecustomresp := appqoecustomrespGetThePayloadFromtheConfig(ctx, &data)

	// Named resource created via the NITRO "import" action (POST ?action=Import).
	// Mirror the SDK v2 verb casing exactly: lower-case "import".
	err := r.client.ActOnResource(service.Appqoecustomresp.Type(), &appqoecustomresp, "import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appqoecustomresp, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appqoecustomresp resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	if !r.readAppqoecustomrespFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appqoecustomresp not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoecustomrespResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppqoecustomrespResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appqoecustomresp resource")

	found := r.readAppqoecustomrespFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoecustomrespResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppqoecustomrespResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// name and src are both RequiresReplace and NITRO exposes no in-place update
	// for the imported page (SDK v2 declared no Update), so Terraform never
	// invokes Update for a real attribute change; just refresh from the API.
	tflog.Debug(ctx, "Update is a no-op for appqoecustomresp; all attributes are RequiresReplace")

	if !r.readAppqoecustomrespFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appqoecustomresp not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoecustomrespResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppqoecustomrespResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appqoecustomresp resource")
	// Named resource - delete using DeleteResource (matches SDK v2)
	appqoecustomrespName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Appqoecustomresp.Type(), appqoecustomrespName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appqoecustomresp, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appqoecustomresp resource")
}

// Helper function to read appqoecustomresp data from API.
//
// appqoecustomresp has no GET-by-name endpoint (only "get (all)"), so mirror the
// SDK v2 resource: fetch all instances and match by name.
func (r *AppqoecustomrespResource) readAppqoecustomrespFromApi(ctx context.Context, data *AppqoecustomrespResourceModel, diags *diag.Diagnostics) bool {
	appqoecustomrespName := data.Id.ValueString()

	dataArr, err := r.client.FindAllResources(service.Appqoecustomresp.Type())
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appqoecustomresp, got error: %s", err))
		return false
	}
	if len(dataArr) == 0 {
		return false
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["name"].(string) == appqoecustomrespName {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return false
	}

	appqoecustomrespSetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
