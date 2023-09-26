package provider

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/crypto/pbkdf2"
)

var (
	_ resource.Resource                = &keyResource{}
)

func NewKeyResource() resource.Resource {
	return &keyResource{}
}

type keyResource struct{}

func (r *keyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_key"
}

func (r *keyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PBKDF2 derived key.",

		Attributes: map[string]schema.Attribute{
			"iterations": schema.Int64Attribute{
				MarkdownDescription: "Number of iterations.",
				Optional:            true,
				Computed:            true,
				Default: int64default.StaticInt64(100000),
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Base secret.",
				Required: 			 true,
				Sensitive:           true,
			},
			"key": schema.StringAttribute{
				MarkdownDescription: "Derived key.",
				Computed:            true,
				Sensitive:           true,
			},
		},
	}
}

type keyResourceData struct {
	Iterations      types.Int64  `tfsdk:"iterations"`
	Password        types.String `tfsdk:"password"`
	Key 			types.String `tfsdk:"key"`
}

func (r keyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan keyResourceData
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var salt = make([]byte, 16)
	_, err := rand.Read(salt[:])
	if err != nil {
		resp.Diagnostics.AddError("Salt Error", err.Error())
		return
	}
	dk := pbkdf2.Key([]byte(plan.Password.ValueString()), salt, int(plan.Iterations.ValueInt64()), 32, sha256.New)
	dk64 :=	base64.StdEncoding.EncodeToString([]byte(dk))
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("iterations"), plan.Iterations)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("password"), plan.Password)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("key"), dk64)...)
	
}

func (r keyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Not needed
}

func (r keyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan keyResourceData
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var salt = make([]byte, 16)
	_, err := rand.Read(salt[:])
	if err != nil {
		resp.Diagnostics.AddError("Salt Error", err.Error())
		return
	}
	dk := pbkdf2.Key([]byte(plan.Password.ValueString()), salt, int(plan.Iterations.ValueInt64()), 32, sha256.New)
	dk64 :=	base64.StdEncoding.EncodeToString([]byte(dk))
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("iterations"), plan.Iterations)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("password"), plan.Password)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("key"), dk64)...)
}

func (r keyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
