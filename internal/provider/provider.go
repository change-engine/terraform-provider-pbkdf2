package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &pbkdf2Provider{}

type pbkdf2Provider struct{
	version string
}

func (p *pbkdf2Provider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "pbkdf2"
	resp.Version = p.version
}

func (p *pbkdf2Provider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This is for using PBKDF2 for deriving keys.",
	}
}

func (p *pbkdf2Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *pbkdf2Provider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *pbkdf2Provider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewKeyResource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &pbkdf2Provider{
			version: version,
		}
	}
}