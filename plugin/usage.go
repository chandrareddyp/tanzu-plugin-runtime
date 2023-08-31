// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/component"
)

// UsageFunc is the usage func for a plugin.
var UsageFunc = func(c *cobra.Command) error {
	t, err := template.New("usage").Funcs(TemplateFuncs).Parse(CmdTemplate)
	if err != nil {
		return err
	}
	return t.Execute(os.Stdout, c)
}

// CmdTemplate is the template for plugin commands.
const CmdTemplate = `{{ bold "Usage:" }}
  {{if .Runnable}}{{ $target := index .Annotations "target" }}{{ if or (eq $target "kubernetes") (eq $target "k8s") }}tanzu {{.UseLine}}{{ end }}{{ if and (ne $target "global") (ne $target "") }}tanzu {{ $target }} {{ else }} {{ end }}{{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}{{ $target := index .Annotations "target" }}{{ if or (eq $target "kubernetes") (eq $target "k8s") }}tanzu {{.CommandPath}} [command]{{end}}{{ if and (ne $target "global") (ne $target "") }}tanzu {{ $target }} {{ else }} {{ end }}{{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

{{ bold "Aliases:" }}
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

{{ bold "Examples:" }}
  {{.Example}}{{end}}{{if .HasAvailableSubCommands}}

{{ bold "Available Commands:" }}{{range .Commands}}{{if .IsAvailableCommand }}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

{{ bold "Flags:" }}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

{{ bold "Global Flags:" }}
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

{{ bold "Additional help topics:" }}{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

{{ $target := index .Annotations "target" }}{{ if or (eq $target "kubernetes") (eq $target "k8s") }}Use "{{if beginsWith .CommandPath "tanzu "}}{{.CommandPath}}{{else}}tanzu {{.CommandPath}}{{end}} [command] --help" for more information about a command.{{end}}Use "{{if beginsWith .CommandPath "tanzu "}}{{.CommandPath}}{{else}}tanzu{{ $target := index .Annotations "target" }}{{ if and (ne $target "global") (ne $target "") }} {{ $target }} {{ else }} {{ end }}{{.CommandPath}}{{end}} [command] --help" for more information about a command.{{end}}
`

// TemplateFuncs are the template usage funcs.
var TemplateFuncs = template.FuncMap{
	"rpad":                    component.Rpad,
	"bold":                    component.Bold,
	"underline":               component.Underline,
	"trimTrailingWhitespaces": component.TrimRightSpace,
	"beginsWith":              component.BeginsWith,
}
