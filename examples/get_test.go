package examples

import "testing"

func TestGetEndpoint(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want Endpoints
	}{
		{
			name: "Test Get Endpoint",
			args: args{
				endpoint: "https://api.github.com",
			},
			want: Endpoints{
				CurrentUserURL:                   "https://api.github.com/user",
				AuthorizationsURL:                "https://api.github.com/authorizations",
				CurrentUserAuthorizationsHTMLURL: "https://github.com/settings/connections/applications{/client_id}",
				RepositoryURL:                    "https://api.github.com/repos/{owner}/{repo}",
				CodeSearchURL:                    "https://api.github.com/search/code?q={query}{&page,per_page,sort,order}",
				CommitSearchURL:                  "https://api.github.com/search/commits?q={query}{&page,per_page,sort,order}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGithubEndpoints(tt.args.endpoint)
			if err != nil {
				t.Errorf("GetGithubEndpoints() error = %v", err)
				return
			}
			if got.CurrentUserURL != tt.want.CurrentUserURL {
				t.Errorf("GetGithubEndpoints() got = %v, want %v", got, tt.want)
			}
			if got.AuthorizationsURL != tt.want.AuthorizationsURL {
				t.Errorf("GetGithubEndpoints() got = %v, want %v", got, tt.want)
			}
			if got.CurrentUserAuthorizationsHTMLURL != tt.want.CurrentUserAuthorizationsHTMLURL {
				t.Errorf("GetGithubEndpoints() got = %v, want %v", got, tt.want)
			}
			if got.RepositoryURL != tt.want.RepositoryURL {
				t.Errorf("GetGithubEndpoints() got = %v, want %v", got, tt.want)
			}
			if got.CodeSearchURL != tt.want.CodeSearchURL {
				t.Errorf("GetGithubEndpoints() got = %v, want %v", got, tt.want)
			}
			if got.CommitSearchURL != tt.want.CommitSearchURL {
				t.Errorf("GetGithubEndpoints() got = %v, want %v", got, tt.want)
			}
		})
	}
}
