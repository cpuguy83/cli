package system

import (
	"runtime"
	"time"

	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/templates"
	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var versionTemplate = `Client:
{{- with .Client}}
{{- if ne .Platform.Name ""}}
 Platform:	{{.Platform.Name}}
{{- end}}
 Version:	{{.Version}}
 API version:	{{.APIVersion}}{{if ne .APIVersion .DefaultAPIVersion}} (downgraded from {{.DefaultAPIVersion}}){{end}}
 Go version:	{{.GoVersion}}
 Git commit:	{{.GitCommit}}
 Built:	{{.BuildTime}}
 OS/Arch:	{{.Os}}/{{.Arch}}
{{- end}}

{{- if .ServerOK}}{{with .Server}}

Server:
 {{- if ne .Platform.Name ""}}
 Platform:	{{.Platform.Name}}
 {{- end}}
 {{- range $component := .Components}}
 {{$component.Name}}:
  {{- if eq $component.Name "Engine" }}
  Version:	{{.Version}}
  API version:	{{index .Details "APIVersion"}} (minimum version {{index .Details "MinAPIVersion"}})
  Go version:	{{index .Details "GoVersion"}}
  Git commit:	{{index .Details "GitCommit"}}
  Built:	{{index .Details "BuildTime"}}
  OS/Arch:	{{index .Details "Os"}}/{{index .Details "Arch"}}
  Experimental:	{{index .Details "Experimental"}}
  {{- else -}}
  Version:	{{$component.Version}}
   {{- range $k, $v := $component.Details}}
  {{$k}}:	{{$v}}
   {{- end}}
  {{- end}}
 {{- end}}
{{- end}}{{end}}`

type versionOptions struct {
	format string
}

// versionInfo contains version information of both the Client, and Server
type versionInfo struct {
	Client clientVersion
	Server *types.Version
}

type clientVersion struct {
	Platform struct{ Name string } `json:",omitempty"`

	Version           string
	APIVersion        string `json:"ApiVersion"`
	DefaultAPIVersion string `json:"DefaultAPIVersion,omitempty"`
	GitCommit         string
	GoVersion         string
	Os                string
	Arch              string
	BuildTime         string `json:",omitempty"`
}

/*
type serverVersion struct {
	types.Version
}

func (sv *serverVersion) UnmarshalJSON(data []byte) error {
	v := &sv.Version
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	foundEngine := false
	for _, component := range sv.Components {
		if component.Name == "Engine" {
			foundEngine = true
			break
		}
	}
	if !foundEngine {
		sv.Components = append(sv.Components, types.ComponentVersion{
			Name:    "Engine",
			Version: v.Version,
			Details: map[string]string{
				"APIVersion":    v.APIVersion,
				"MinAPIVersion": v.MinAPIVersion,
				"GitCommit":     v.GitCommit,
				"GoVersion":     v.GoVersion,
				"Os":            v.Os,
				"Arch":          v.Arch,
				"BuildTime":     v.BuildTime,
			},
		})
	}
	return nil
}
*/

// ServerOK returns true when the client could connect to the docker server
// and parse the information received. It returns false otherwise.
func (v versionInfo) ServerOK() bool {
	return v.Server != nil
}

// NewVersionCommand creates a new cobra.Command for `docker version`
func NewVersionCommand(dockerCli *command.DockerCli) *cobra.Command {
	var opts versionOptions

	cmd := &cobra.Command{
		Use:   "version [OPTIONS]",
		Short: "Show the Docker version information",
		Args:  cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion(dockerCli, &opts)
		},
	}

	flags := cmd.Flags()

	flags.StringVarP(&opts.format, "format", "f", "", "Format the output using the given Go template")

	return cmd
}

func runVersion(dockerCli *command.DockerCli, opts *versionOptions) error {
	ctx := context.Background()

	templateFormat := versionTemplate
	if opts.format != "" {
		templateFormat = opts.format
	}

	tmpl, err := templates.Parse(templateFormat)
	if err != nil {
		return cli.StatusError{StatusCode: 64,
			Status: "Template parsing error: " + err.Error()}
	}

	vd := versionInfo{
		Client: clientVersion{
			Version:           cli.Version,
			APIVersion:        dockerCli.Client().ClientVersion(),
			DefaultAPIVersion: dockerCli.DefaultVersion(),
			GoVersion:         runtime.Version(),
			GitCommit:         cli.GitCommit,
			BuildTime:         cli.BuildTime,
			Os:                runtime.GOOS,
			Arch:              runtime.GOARCH,
		},
	}
	vd.Client.Platform.Name = cli.PlatformName

	// TODO: simply export Get to retrieve raw response
	//_, body, err := dockerCli.Client().ServerVersion(ctx)
	sv, err := dockerCli.Client().ServerVersion(ctx)
	if err == nil {
		vd.Server = &sv
		foundEngine := false
		for _, component := range sv.Components {
			if component.Name == "Engine" {
				foundEngine = true
				break
			}
		}
		if !foundEngine {
			vd.Server.Components = append(vd.Server.Components, types.ComponentVersion{
				Name:    "Engine",
				Version: sv.Version,
				Details: map[string]string{
					"ApiVersion":    sv.APIVersion,
					"MinAPIVersion": sv.MinAPIVersion,
					"GitCommit":     sv.GitCommit,
					"GoVersion":     sv.GoVersion,
					"Os":            sv.Os,
					"Arch":          sv.Arch,
					"BuildTime":     sv.BuildTime,
				},
			})
		}
		//err = json.NewDecoder(bytes.NewReader(body)).Decode(&vd.Server)
	}

	// first we need to make BuildTime more human friendly
	t, errTime := time.Parse(time.RFC3339Nano, vd.Client.BuildTime)
	if errTime == nil {
		vd.Client.BuildTime = t.Format(time.ANSIC)
	}

	if vd.ServerOK() {
		t, errTime = time.Parse(time.RFC3339Nano, vd.Server.BuildTime)
		if errTime == nil {
			vd.Server.BuildTime = t.Format(time.ANSIC)
		}
	}

	//tw := tabwriter.NewWriter(dockerCli.Out(), 0, 0, 8, ' ', tabwriter.Debug)
	tw := dockerCli.Out()
	if err2 := tmpl.Execute(tw, vd); err2 != nil && err == nil {
		err = err2
	}
	dockerCli.Out().Write([]byte{'\n'})
	return err
}
