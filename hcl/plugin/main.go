package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/hcl/v2/hclparse"
	"plugin"
)

// Config of plugin.hcl
type Plugin struct {
	Type         string   `hcl:"type,label"`
	Name         string   `hcl:"name,label"`
	Cmd          string   `hcl:"plugin_cmd"`
	PluginConfig hcl.Body `hcl:",remain"`
}
type Config struct {
	IOMode  string   `hcl:"io_mode"`
	Plugins []Plugin `hcl:"plugin,block"`
}

type Mgr struct {
	PluginType1 *plugin.Plugin
}

func main() {
	parser := hclparse.NewParser()
	f, diags := parser.ParseHCLFile("./plugin.hcl")
	if diags.HasErrors() {
		panic(diags.Error())
	}
	var c Config
	moreDiags := gohcl.DecodeBody(f.Body, nil, &c)
	diags = append(diags, moreDiags...)
	if moreDiags.HasErrors() {
		panic(diags.Error())
	}
	fmt.Printf("Config = %+v\n", c)
	var pluginMgr Mgr

	for _, sc := range c.Plugins {
		// Totally-hypothetical plugin manager (not part of HCL)
		myPlugin := pluginMgr.GetHcldecPlugin(sc.Type, sc.Cmd)

		configSpec, err := myPlugin.Lookup("ConfigSpec")
		if err != nil {
			panic(err)
		}
		spec := configSpec.(func() hcldec.Spec)()

		// Decode the block body using the plugin's given specification
		configVal, moreDiags := hcldec.Decode(sc.PluginConfig, spec, nil)
		diags = append(diags, moreDiags...)
		if moreDiags.HasErrors() {
			panic(moreDiags.Error())
		}
		fmt.Printf("configVal = %+v\n", configVal)

		fmt.Printf("PluginConfig = %+v\n", sc.PluginConfig)
		toString, err := myPlugin.Lookup("ToString")
		if err != nil {
			panic(err)
		}
		fmt.Println(toString.(func(hcl.Body) string)(sc.PluginConfig))

	}
}

func (m *Mgr) GetHcldecPlugin(myType string, cmd string) *plugin.Plugin {
	myPlugin, err := plugin.Open(cmd)
	if err != nil {
		panic(err)
	}
	switch myType {
	case "plugin_type1":
		m.PluginType1 = myPlugin
	default:
		panic("Should be one of [plugin_type1]")
	}
	return myPlugin
}
