// Code generated by pluginator on IAMPolicyGenerator; DO NOT EDIT.
// pluginator {(devel)  unknown   }

package builtins

import (
	"sigs.k8s.io/kustomize/api/filters/iampolicygenerator"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/yaml"
)

type IAMPolicyGeneratorPlugin struct {
	types.IAMPolicyGeneratorArgs
}

func (p *IAMPolicyGeneratorPlugin) Config(h *resmap.PluginHelpers, config []byte) (err error) {
	p.IAMPolicyGeneratorArgs = types.IAMPolicyGeneratorArgs{}
	err = yaml.Unmarshal(config, p)
	return
}

func (p *IAMPolicyGeneratorPlugin) Generate() (resmap.ResMap, error) {
	r := resmap.New()
	err := r.ApplyFilter(iampolicygenerator.Filter{
		IAMPolicyGenerator: p.IAMPolicyGeneratorArgs,
	})
	return r, err
}

func NewIAMPolicyGeneratorPlugin() resmap.GeneratorPlugin {
	return &IAMPolicyGeneratorPlugin{}
}
