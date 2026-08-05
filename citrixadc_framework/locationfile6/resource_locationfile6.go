package locationfile6

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
var _ resource.Resource = &Locationfile6Resource{}
var _ resource.ResourceWithConfigure = (*Locationfile6Resource)(nil)
var _ resource.ResourceWithImportState = (*Locationfile6Resource)(nil)

func NewLocationfile6Resource() resource.Resource {
	return &Locationfile6Resource{}
}

// Locationfile6Resource defines the resource implementation.
type Locationfile6Resource struct {
	client *service.NitroClient
}

func (r *Locationfile6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Locationfile6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_locationfile6"
}

func (r *Locationfile6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Locationfile6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Locationfile6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating locationfile6 resource")

	locationfile6 := locationfile6GetThePayloadFromtheConfig(ctx, &data)

	// SDK v2 createLocationfile6Func used AddResource with an empty name.
	_, err := r.client.AddResource(service.Locationfile6.Type(), "", &locationfile6)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create locationfile6, got error: %s", err))
		return
	}

	// Backward-compatible ID: SDK v2 used d.SetId(locationfile).
	data.Id = types.StringValue(data.Locationfile.ValueString())

	tflog.Trace(ctx, "Created locationfile6 resource")

	// Read the updated state back
	if !r.readLocationfile6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "locationfile6 not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Locationfile6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Locationfile6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading locationfile6 resource")

	found := r.readLocationfile6FromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Locationfile6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// locationfile6 exposes no NITRO update endpoint and every schema attribute is
	// RequiresReplace (SDK v2 ForceNew), so Terraform never invokes Update for a
	// real attribute change. Preserve the prior ID and re-read state.
	var data, state Locationfile6ResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for locationfile6; NITRO has no update endpoint and all attributes are RequiresReplace")

	// Read the updated state back
	if !r.readLocationfile6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "locationfile6 not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Locationfile6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Locationfile6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting locationfile6 resource")

	// SDK v2 deleteLocationfile6Func used DeleteResource with an empty name.
	err := r.client.DeleteResource(service.Locationfile6.Type(), "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete locationfile6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted locationfile6 resource")
}

// readLocationfile6FromApi reads the locationfile6 configuration from the ADC.
// Returns false (without an error diagnostic) when the resource is not found so
// callers can drop it from state, mirroring the SDK v2 read which cleared the ID
// on a find failure.
func (r *Locationfile6Resource) readLocationfile6FromApi(ctx context.Context, data *Locationfile6ResourceModel, diags *diag.Diagnostics) bool {
	getResponseData, err := r.client.FindResource(service.Locationfile6.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read locationfile6, got error: %s", err))
		return false
	}

	locationfile6SetAttrFromGet(ctx, data, getResponseData)

	return true
}
