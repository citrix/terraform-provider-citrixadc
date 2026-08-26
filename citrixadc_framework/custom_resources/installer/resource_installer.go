package installer

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/utility"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	sdkid "github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &InstallerResource{}
var _ resource.ResourceWithConfigure = (*InstallerResource)(nil)

func NewInstallerResource() resource.Resource {
	return &InstallerResource{}
}

// InstallerResource models the NetScaler build-install / system-upgrade action.
//
// It is a one-shot side-effect action: Create POSTs the `install` NITRO action
// (equivalent to the CLI `install ns` / build upgrade). NITRO has no GET endpoint
// that reports install state and there is no inverse API, so Read/Update/Delete
// are no-ops and every input attribute is RequiresReplace (mirrors the legacy
// SDKv2 citrixadc_installer, whose fields were all ForceNew with Read/Delete
// Noop). The install typically reboots the appliance, so the NITRO call may fail
// with a TCP reset / EOF; that error is expected and is tolerated.
//
// When wait_until_reachable is true, Create blocks after issuing the install and
// polls the appliance's nslicense endpoint until it responds (i.e. the box has
// rebooted and come back), bounded by reachable_timeout. The reachable_*
// attributes tune that poll loop, preserving the SDKv2 StateChangeConf semantics.
type InstallerResource struct {
	client *service.NitroClient
}

// InstallerResourceModel describes the resource data model. Every schema
// attribute has a matching tfsdk field, with the same names/types as the legacy
// SDKv2 schema.
type InstallerResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Advancedoptions          types.String `tfsdk:"advancedoptions"`
	Answeryestoall           types.Bool   `tfsdk:"answeryestoall"`
	Deletesigfiles           types.Bool   `tfsdk:"deletesigfiles"`
	Dontchecknsconf          types.Bool   `tfsdk:"dontchecknsconf"`
	Dontreboot               types.Bool   `tfsdk:"dontreboot"`
	Enhancedupgrade          types.Bool   `tfsdk:"enhancedupgrade"`
	Exitonlicserverconnerror types.Bool   `tfsdk:"exitonlicserverconnerror"`
	Fipsinstall              types.Bool   `tfsdk:"fipsinstall"`
	Ignorecertcheckerrors    types.Bool   `tfsdk:"ignorecertcheckerrors"`
	Ignorensapimgrerrors     types.Bool   `tfsdk:"ignorensapimgrerrors"`
	Ignoreunsavedconfig      types.Bool   `tfsdk:"ignoreunsavedconfig"`
	Ignoreunsyncedconfig     types.Bool   `tfsdk:"ignoreunsyncedconfig"`
	L                        types.Bool   `tfsdk:"l"`
	Precheck                 types.Bool   `tfsdk:"precheck"`
	Resizeswapvar            types.Bool   `tfsdk:"resizeswapvar"`
	Url                      types.String `tfsdk:"url"`
	Y                        types.Bool   `tfsdk:"y"`
	WaitUntilReachable       types.Bool   `tfsdk:"wait_until_reachable"`
	ReachableTimeout         types.String `tfsdk:"reachable_timeout"`
	ReachablePollDelay       types.String `tfsdk:"reachable_poll_delay"`
	ReachablePollInterval    types.String `tfsdk:"reachable_poll_interval"`
	ReachablePollTimeout     types.String `tfsdk:"reachable_poll_timeout"`
}

func (r *InstallerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_installer"
}

func (r *InstallerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *InstallerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the installer resource.",
			},
			"advancedoptions": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Use this string to pass extra flags which are not yet supported. Example: -flag1 -flag2 -flag3.",
			},
			"answeryestoall": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to answer yes to all prompts.",
			},
			"deletesigfiles": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to delete all signature files and associated kernel images during installation.",
			},
			"dontchecknsconf": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to skip ns.conf version equivalence check during downgrade.",
			},
			"dontreboot": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to prevent reboot after installation when answerYesToAll is true.",
			},
			"exitonlicserverconnerror": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to exit on license server connectivity errors.",
			},
			"fipsinstall": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to perform FIPS installation.",
			},
			"ignorecertcheckerrors": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to ignore certificate digest verification errors during build update.",
			},
			"ignorensapimgrerrors": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to ignore nsapimgr symbols not found error(s) during installation.",
			},
			"ignoreunsavedconfig": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to ignore unsaved config check during build update.",
			},
			"ignoreunsyncedconfig": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to ignore unsynced HA config check during build update.",
			},
			"precheck": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to run all installation pre-checks in a single step.",
			},
			"enhancedupgrade": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag for upgrading from/to enhancement mode.",
			},
			"l": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to enable callhome.",
			},
			"resizeswapvar": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this flag to change swap size on ONLY 64bit nCore/MCNS/VMPE systems NON-VPX systems.",
			},
			"url": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Url for the build file.",
			},
			"y": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Do not prompt for yes/no before rebooting.",
			},
			"wait_until_reachable": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Set to `true` to block until the appliance is reachable again after the install (which typically reboots the box).",
			},
			"reachable_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("10m"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Overall timeout for the wait_until_reachable poll loop.",
			},
			"reachable_poll_delay": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("60s"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Initial delay before the first reachability poll.",
			},
			"reachable_poll_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("60s"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interval between reachability polls.",
			},
			"reachable_poll_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("20s"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Per-poll HTTP timeout for a single reachability check.",
			},
		},
	}
}

func (r *InstallerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data InstallerResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating installer resource")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform. Uses the same "tf-installer-" prefixed unique ID
	// scheme as the legacy SDKv2 resource.
	installerId := sdkid.PrefixedUniqueId("tf-installer-")

	// Issue the install action.
	if err := r.installerInstallBuild(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to install build, got error: %s", err))
		return
	}

	// Optionally block until the appliance is reachable again.
	if data.WaitUntilReachable.ValueBool() {
		if err := r.installerWaitReachable(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Appliance did not become reachable, got error: %s", err))
			return
		}
	}

	data.Id = types.StringValue(installerId)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InstallerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// The install is a one-shot action. NITRO has no GET endpoint that reports
	// install state, so Read is a pure preserve-state no-op (mirrors the SDKv2
	// Read: schema.Noop).
	var data InstallerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for installer; NITRO has no GET endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InstallerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for the install action; every schema attribute
	// is RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state InstallerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for installer; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InstallerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// The install is a one-shot side-effect action. There is no inverse NITRO API
	// (mirrors the SDKv2 Delete: schema.Noop). Delete only removes the resource
	// from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for installer; NITRO has no inverse of the install action")
}

// installerInstallBuild builds the utility.Install payload from the plan and
// issues the `install` NITRO action. The install typically reboots the box, so a
// TCP reset / EOF from the NITRO call is expected and tolerated (mirrors the
// SDKv2 installerInstallBuild).
func (r *InstallerResource) installerInstallBuild(ctx context.Context, data *InstallerResourceModel) error {
	tflog.Debug(ctx, "In installerInstallBuild")

	install := installerGetThePayloadFromthePlan(ctx, data)

	if err := r.client.ActOnResource("install", &install, ""); err != nil {
		errorStr := err.Error()
		if strings.HasSuffix(errorStr, "EOF") || strings.HasSuffix(errorStr, "connection reset by peer") {
			// This is expected since the operation results in a TCP connection reset
			// some times, especially when y = true.
			tflog.Debug(ctx, fmt.Sprintf("Ignoring go-nitro error \"%s\"", errorStr))
			return nil
		}
		return err
	}
	return nil
}

// installerWaitReachable blocks until the appliance responds on the nslicense
// endpoint or reachable_timeout elapses, honoring the reachable_poll_delay /
// reachable_poll_interval tunables (mirrors the SDKv2 StateChangeConf loop).
func (r *InstallerResource) installerWaitReachable(ctx context.Context, data *InstallerResourceModel) error {
	tflog.Debug(ctx, "In installerWaitReachable")

	timeout, err := time.ParseDuration(data.ReachableTimeout.ValueString())
	if err != nil {
		return err
	}
	pollInterval, err := time.ParseDuration(data.ReachablePollInterval.ValueString())
	if err != nil {
		return err
	}
	pollDelay, err := time.ParseDuration(data.ReachablePollDelay.ValueString())
	if err != nil {
		return err
	}

	// Initial delay before the first poll (StateChangeConf.Delay).
	time.Sleep(pollDelay)

	deadline := time.Now().Add(timeout)
	for {
		if err := r.installerPollLicense(ctx, data); err == nil {
			tflog.Debug(ctx, "Appliance is reachable")
			return nil
		} else {
			tflog.Debug(ctx, fmt.Sprintf("Unreachable: %v", err.Error()))
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout after %s waiting for the appliance to become reachable", timeout)
		}
		time.Sleep(pollInterval)
	}
}

// installerPollLicense performs a single reachability check by GETting the
// nslicense endpoint with a per-poll HTTP timeout (mirrors the SDKv2
// installerPollLicense). A nil return means the appliance is reachable.
func (r *InstallerResource) installerPollLicense(ctx context.Context, data *InstallerResourceModel) error {
	tflog.Debug(ctx, "In installerPollLicense")

	username := r.client.GetUsername()
	password := r.client.GetPassword()
	endpoint := r.client.GetURL()
	url := fmt.Sprintf("%s/nitro/v1/config/nslicense", endpoint)

	timeout, err := time.ParseDuration(data.ReachablePollTimeout.ValueString())
	if err != nil {
		return err
	}
	c := http.Client{
		Timeout: timeout,
	}
	buff := &bytes.Buffer{}
	httpReq, _ := http.NewRequest("GET", url, buff)
	httpReq.Header.Set("X-NITRO-USER", username)
	httpReq.Header.Set("X-NITRO-PASS", password)
	resp, err := c.Do(httpReq)
	if err != nil {
		return err
	}
	tflog.Debug(ctx, fmt.Sprintf("Status code is %v", resp.Status))
	resp.Body.Close()
	return nil
}

// installerGetThePayloadFromthePlan builds the utility.Install body from the
// plan attributes (mirrors the SDKv2 installerInstallBuild payload assembly).
func installerGetThePayloadFromthePlan(ctx context.Context, data *InstallerResourceModel) utility.Install {
	tflog.Debug(ctx, "In installerGetThePayloadFromthePlan Function")

	install := utility.Install{
		Advancedoptions:          data.Advancedoptions.ValueString(),
		Answeryestoall:           data.Answeryestoall.ValueBool(),
		Deletesigfiles:           data.Deletesigfiles.ValueBool(),
		Dontchecknsconf:          data.Dontchecknsconf.ValueBool(),
		Dontreboot:               data.Dontreboot.ValueBool(),
		Enhancedupgrade:          data.Enhancedupgrade.ValueBool(),
		Exitonlicserverconnerror: data.Exitonlicserverconnerror.ValueBool(),
		Fipsinstall:              data.Fipsinstall.ValueBool(),
		Ignorecertcheckerrors:    data.Ignorecertcheckerrors.ValueBool(),
		Ignorensapimgrerrors:     data.Ignorensapimgrerrors.ValueBool(),
		Ignoreunsavedconfig:      data.Ignoreunsavedconfig.ValueBool(),
		Ignoreunsyncedconfig:     data.Ignoreunsyncedconfig.ValueBool(),
		L:                        data.L.ValueBool(),
		Precheck:                 data.Precheck.ValueBool(),
		Resizeswapvar:            data.Resizeswapvar.ValueBool(),
		Url:                      data.Url.ValueString(),
		Y:                        data.Y.ValueBool(),
	}
	return install
}
