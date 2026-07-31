package systembackup

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	sdkresource "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystembackupCreateResource{}
var _ resource.ResourceWithConfigure = (*SystembackupCreateResource)(nil)

func NewSystembackupCreateResource() resource.Resource {
	return &SystembackupCreateResource{}
}

// SystembackupCreateResource models the NITRO systembackup `?action=create`
// action, which produces a persistent backup object (*.tgz). This is an
// action-with-read-delete resource: Create fires the create action, Read looks
// up the produced backup file via GET (clearing state when it is gone), and
// Delete removes the backup file. This mirrors the SDK v2
// citrixadc_systembackup_create resource exactly.
type SystembackupCreateResource struct {
	client *service.NitroClient
}

// SystembackupCreateResourceModel describes the resource data model. The
// attribute set mirrors the SDK v2 schema exactly (comment, filename,
// includekernel, level, uselocaltimezone). No provider-side-only attributes
// exist for this resource.
type SystembackupCreateResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Comment          types.String `tfsdk:"comment"`
	Filename         types.String `tfsdk:"filename"`
	Includekernel    types.String `tfsdk:"includekernel"`
	Level            types.String `tfsdk:"level"`
	Uselocaltimezone types.Bool   `tfsdk:"uselocaltimezone"`
}

func (r *SystembackupCreateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systembackup_create"
}

func (r *SystembackupCreateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystembackupCreateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systembackup_create resource.",
			},
			// All configurable attributes are ForceNew in SDK v2, mapped here to
			// RequiresReplace() plan modifiers. None are Computed, matching SDK v2.
			"comment": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Comment specified at the time of creation of the backup file(*.tgz).",
			},
			"filename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the backup file(*.tgz) to be restored.",
			},
			"includekernel": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Use this option to add kernel in the backup file",
			},
			"level": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Level of data to be backed up.",
			},
			"uselocaltimezone": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "This option will create backup file with local timezone timestamp",
			},
		},
	}
}

func (r *SystembackupCreateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystembackupCreateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systembackup_create resource")

	// Mirror SDK v2 id: resource.PrefixedUniqueId(filename + "-").
	systembackupName := sdkresource.PrefixedUniqueId(data.Filename.ValueString() + "-")

	systembackup := systembackup_createGetThePayloadFromthePlan(ctx, &data)

	// systembackup create is a POST ?action=create action (ActOnResource with
	// the "create" verb in SDK v2).
	err := r.client.ActOnResource(service.Systembackup.Type(), &systembackup, "create")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systembackup, got error: %s", err))
		return
	}

	data.Id = types.StringValue(systembackupName)

	tflog.Trace(ctx, "Created systembackup_create resource")

	// Mirror SDK v2 which calls Read at the end of Create.
	r.readSystembackupCreateFromApi(ctx, &data, resp.Diagnostics.AddError)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystembackupCreateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystembackupCreateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systembackup_create resource")

	// Look up the produced backup file. When it is gone, mirror SDK v2's
	// d.SetId("") by removing the resource from state.
	found := r.readSystembackupCreateFromApi(ctx, &data, resp.Diagnostics.AddError)
	if !found {
		tflog.Warn(ctx, "systembackup not found on remote; removing from state")
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystembackupCreateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for systembackup create; every configurable
	// attribute is RequiresReplace, so Terraform never invokes Update for a real
	// change. SDK v2 declared no Update function.
	var data, state SystembackupCreateResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systembackup_create; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystembackupCreateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystembackupCreateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systembackup_create resource")

	// Mirror SDK v2: DeleteResource on filename + ".tgz".
	systembackupName := data.Filename.ValueString() + ".tgz"
	err := r.client.DeleteResource(service.Systembackup.Type(), systembackupName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systembackup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systembackup_create resource")
}

// readSystembackupCreateFromApi looks up the produced backup file
// (filename + ".tgz") and populates the attributes SDK v2's Read populated
// (comment and uselocaltimezone). It returns false when the backup file is no
// longer present so callers can clear state, mirroring SDK v2's d.SetId("").
func (r *SystembackupCreateResource) readSystembackupCreateFromApi(ctx context.Context, data *SystembackupCreateResourceModel, addError func(string, string)) bool {
	tflog.Debug(ctx, "In readSystembackupCreateFromApi Function")

	getResponseData, err := r.client.FindResource(service.Systembackup.Type(), data.Filename.ValueString()+".tgz")
	if err != nil {
		// Mirror SDK v2: on lookup failure clear the resource from state.
		return false
	}

	systembackup_createSetAttrFromGet(ctx, data, getResponseData)
	return true
}

// systembackup_createGetThePayloadFromthePlan builds the create action body.
// Only fields the NITRO create action accepts are included; there are no
// provider-side-only attributes for this resource.
func systembackup_createGetThePayloadFromthePlan(ctx context.Context, data *SystembackupCreateResourceModel) system.Systembackup {
	tflog.Debug(ctx, "In systembackup_createGetThePayloadFromthePlan Function")

	// Mirror SDK v2 payload construction: all attributes are read directly with
	// their zero values when unset (d.Get returns the zero value).
	systembackup := system.Systembackup{
		Filename:         data.Filename.ValueString(),
		Uselocaltimezone: data.Uselocaltimezone.ValueBool(),
		Level:            data.Level.ValueString(),
		Includekernel:    data.Includekernel.ValueString(),
		Comment:          data.Comment.ValueString(),
	}

	return systembackup
}

// systembackup_createSetAttrFromGet populates the model from a GET response.
// SDK v2's Read only set comment and uselocaltimezone, so this mirrors that
// exactly and leaves the ForceNew config-driven attributes untouched.
func systembackup_createSetAttrFromGet(ctx context.Context, data *SystembackupCreateResourceModel, getResponseData map[string]interface{}) {
	tflog.Debug(ctx, "In systembackup_createSetAttrFromGet Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		if s, isStr := val.(string); isStr {
			data.Comment = types.StringValue(s)
		}
	}
	if val, ok := getResponseData["uselocaltimezone"]; ok && val != nil {
		switch v := val.(type) {
		case bool:
			data.Uselocaltimezone = types.BoolValue(v)
		case string:
			data.Uselocaltimezone = types.BoolValue(v == "true" || v == "YES" || v == "ENABLED")
		}
	}
}
