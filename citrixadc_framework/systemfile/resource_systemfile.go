package systemfile

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	stdpath "path"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemfileResource{}
var _ resource.ResourceWithConfigure = (*SystemfileResource)(nil)
var _ resource.ResourceWithImportState = (*SystemfileResource)(nil)

func NewSystemfileResource() resource.Resource {
	return &SystemfileResource{}
}

// SystemfileResource defines the resource implementation.
type SystemfileResource struct {
	client *service.NitroClient
}

// ImportState mirrors the SDK v2 importer: the import ID is the full file path
// (filelocation/filename). It is split so Read can locate the file on the ADC.
func (r *SystemfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fullPath := req.ID
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), fullPath)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("filelocation"), stdpath.Dir(fullPath))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("filename"), stdpath.Base(fullPath))...)
}

func (r *SystemfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemfile"
}

func (r *SystemfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemfile resource")

	filecontent := data.Filecontent.ValueString()
	fileencoding := data.Fileencoding.ValueString()
	filelocation := data.Filelocation.ValueString()
	filename := data.Filename.ValueString()

	if fileencoding != "BASE64" {
		resp.Diagnostics.AddError("Configuration Error", fmt.Sprintf("file encoding %s is not supported", fileencoding))
		return
	}

	var b64filecontent string
	if data.IsBase64Encoded.ValueBool() {
		tflog.Debug(ctx, "Content is marked as already base64-encoded, passing through")
		b64filecontent = filecontent
	} else {
		tflog.Debug(ctx, "Encoding content to base64")
		b64filecontent = base64.StdEncoding.EncodeToString([]byte(filecontent))
	}

	systemfile := system.Systemfile{
		Filecontent:  b64filecontent,
		Fileencoding: fileencoding,
		Filelocation: filelocation,
		Filename:     filename,
	}

	_, err := r.client.AddResource(service.Systemfile.Type(), "", &systemfile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemfile resource")

	// ID scheme matches SDK v2: full path filelocation/filename.
	data.Id = types.StringValue(stdpath.Join(filelocation, filename))

	// Read the updated state back
	if !r.readSystemfileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "systemfile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemfile resource")

	found := r.readSystemfileFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// File no longer exists on the ADC; clear state so it is recreated.
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All systemfile attributes are ForceNew (RequiresReplace) in the SDK v2
	// contract, so Update is never invoked with a real attribute change.
	// Preserve the prior ID and re-read to keep state consistent.
	var data, state SystemfileResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemfile resource (no updateable attributes; re-reading)")

	if !r.readSystemfileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "systemfile not found during update")
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemfile resource")

	// Mirror SDK v2 delete: DeleteResourceWithArgsMap("systemfile", filename, {filelocation}).
	argsMap := make(map[string]string)
	argsMap["filelocation"] = url.QueryEscape(data.Filelocation.ValueString())
	filename := data.Filename.ValueString()
	err := r.client.DeleteResourceWithArgsMap("systemfile", filename, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systemfile resource")
}

// readSystemfileFromApi reads the systemfile from the ADC using the SDK v2 lookup
// pattern (FindResourceArrayWithParams filtered by filelocation + filename).
// Returns false if the file does not exist.
func (r *SystemfileResource) readSystemfileFromApi(ctx context.Context, data *SystemfileResourceModel, diags *diag.Diagnostics) bool {
	argsMap := make(map[string]string)
	argsMap["filelocation"] = url.QueryEscape(data.Filelocation.ValueString())
	argsMap["filename"] = url.QueryEscape(data.Filename.ValueString())
	findParams := service.FindParams{
		ResourceType: "systemfile",
		ArgsMap:      argsMap,
	}

	dataArray, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		// Matches SDK v2: a lookup error is treated as "file does not exist".
		return false
	}

	if len(dataArray) == 0 {
		tflog.Warn(ctx, "systemfile does not exist. Clearing state.")
		return false
	}

	if len(dataArray) > 1 {
		diags.AddError("Client Error", "multiple entries found for file")
		return false
	}

	systemfileSetAttrFromGet(ctx, data, dataArray[0])

	return true
}
