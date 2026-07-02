package dnscaarec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnscaarecResource{}
var _ resource.ResourceWithConfigure = (*DnscaarecResource)(nil)
var _ resource.ResourceWithImportState = (*DnscaarecResource)(nil)

func NewDnscaarecResource() resource.Resource {
	return &DnscaarecResource{}
}

// DnscaarecResource defines the resource implementation.
type DnscaarecResource struct {
	client *service.NitroClient
}

func (r *DnscaarecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnscaarecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnscaarec"
}

func (r *DnscaarecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnscaarecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnscaarecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnscaarec resource")
	dnscaarec := dnscaarecGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes an 'add' verb (POST) for dnscaarec.
	_, err := r.client.AddResource(service.Dnscaarec.Type(), "", &dnscaarec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnscaarec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnscaarec resource")

	// The recordid is server-assigned and unknown until after create.
	// Locate the newly created record by matching the supplied attributes.
	r.findDnscaarecAfterCreate(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnscaarecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnscaarecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnscaarec resource")

	r.readDnscaarecFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnscaarecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnscaarecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID and server-assigned recordid from prior state.
	data.Id = state.Id
	data.Recordid = state.Recordid

	// dnscaarec has no NITRO update endpoint ("You cannot modify a CAA resource
	// record"). All settable attributes use RequiresReplace, so Update is a no-op.
	tflog.Debug(ctx, "Update is a no-op for dnscaarec; all attributes are RequiresReplace")

	// Read the current state back
	r.readDnscaarecFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnscaarecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnscaarecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnscaarec resource")

	// Delete targets the specific record via the parent domain (URL) plus the
	// server-assigned recordid as a query arg (multiple records share a domain).
	if data.Domain.IsNull() || data.Recordid.IsNull() {
		resp.Diagnostics.AddError("Client Error", "Unable to delete dnscaarec: domain or recordid missing from state")
		return
	}

	args := []string{fmt.Sprintf("recordid:%d", data.Recordid.ValueInt64())}
	err := r.client.DeleteResourceWithArgs(service.Dnscaarec.Type(), data.Domain.ValueString(), args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnscaarec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnscaarec record")
}

// findDnscaarecAfterCreate locates the newly created record (whose recordid is
// unknown) by matching the supplied attributes, and populates data (including
// the server-assigned recordid and composite ID).
func (r *DnscaarecResource) findDnscaarecAfterCreate(ctx context.Context, data *DnscaarecResourceModel, diags *diag.Diagnostics) {
	domain := data.Domain.ValueString()

	findParams := service.FindParams{
		ResourceType:             service.Dnscaarec.Type(),
		ResourceName:             domain,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnscaarec, got error: %s", err))
		return
	}
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "dnscaarec returned empty array.")
		return
	}

	foundIndex := -1
	for i, v := range dataArr {
		match := true
		if !data.Valuestring.IsNull() {
			if s, ok := v["valuestring"].(string); !ok || s != data.Valuestring.ValueString() {
				match = false
			}
		}
		if match && !data.Tag.IsNull() {
			if s, ok := v["tag"].(string); ok && s != data.Tag.ValueString() {
				match = false
			}
		}
		if match {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		diags.AddError("Client Error", "dnscaarec not found after create with the supplied attributes")
		return
	}

	dnscaarecSetAttrFromGet(ctx, data, dataArr[foundIndex])
}

// Helper function to read dnscaarec data from API by domain + recordid.
func (r *DnscaarecResource) readDnscaarecFromApi(ctx context.Context, data *DnscaarecResourceModel, diags *diag.Diagnostics) {
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	domainName, ok := idMap["domain"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'domain' not found in ID string")
		return
	}
	recordidStr, ok := idMap["recordid"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'recordid' not found in ID string")
		return
	}

	findParams := service.FindParams{
		ResourceType:             service.Dnscaarec.Type(),
		ResourceName:             domainName,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnscaarec, got error: %s", err))
		return
	}
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "dnscaarec returned empty array.")
		return
	}

	// Locate the specific record by recordid.
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["recordid"]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				if fmt.Sprintf("%d", intVal) == recordidStr {
					foundIndex = i
					break
				}
			}
		}
	}

	if foundIndex == -1 {
		diags.AddError("Client Error", "dnscaarec not found with the provided recordid")
		return
	}

	dnscaarecSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
