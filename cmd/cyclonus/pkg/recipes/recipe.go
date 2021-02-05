package recipes

import (
	"fmt"
	"github.com/mattfenwick/cyclonus/pkg/connectivity/synthetic"
	"github.com/mattfenwick/cyclonus/pkg/explainer"
	"github.com/mattfenwick/cyclonus/pkg/matcher"
	"github.com/mattfenwick/cyclonus/pkg/utils"
	v1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/yaml"
)

type Recipe struct {
	PolicyYamls []string
	Resources   *synthetic.Resources
	Protocol    v1.Protocol
	Port        int
}

func (r *Recipe) Policies() []*networkingv1.NetworkPolicy {
	var policies []*networkingv1.NetworkPolicy
	for _, yamlString := range r.PolicyYamls {
		netpol := networkingv1.NetworkPolicy{}
		err := yaml.Unmarshal([]byte(yamlString), &netpol)
		utils.DoOrDie(err)
		policies = append(policies, &netpol)
	}
	return policies
}

func (r *Recipe) RunProbe() *synthetic.Result {
	return synthetic.RunSyntheticProbe(&synthetic.Request{
		Protocol:  r.Protocol,
		Port:      intstr.FromInt(r.Port),
		Policies:  matcher.BuildNetworkPolicies(r.Policies()),
		Resources: r.Resources,
	})
}

var AllRecipes = []*Recipe{
	{[]string{Recipe01}, Resources01, v1.ProtocolTCP, 80},
	{[]string{Recipe02}, Resources02, v1.ProtocolTCP, 80},
	{[]string{Recipe01, Recipe02A}, Resources02A, v1.ProtocolTCP, 80},
	{[]string{Recipe03}, Resources03, v1.ProtocolTCP, 80},
	{[]string{Recipe04}, Resources04, v1.ProtocolTCP, 80},
	{[]string{Recipe01, Recipe05}, Resources05, v1.ProtocolTCP, 80},
	{[]string{Recipe06}, Resources06, v1.ProtocolTCP, 80},
	{[]string{Recipe07}, Resources07, v1.ProtocolTCP, 80},
	{[]string{Recipe01, Recipe08}, Resources08, v1.ProtocolTCP, 80},
	{[]string{Recipe09}, Resources09, v1.ProtocolTCP, 5000},
	{[]string{Recipe10}, Resources10, v1.ProtocolTCP, 80},
	{[]string{Recipe11_1}, Resources11_1, v1.ProtocolTCP, 80},
	{[]string{Recipe11_2}, Resources11_2, v1.ProtocolTCP, 53},
	{[]string{Recipe12}, Resources12, v1.ProtocolTCP, 80},
	{[]string{Recipe14}, Resources14, v1.ProtocolTCP, 80},
}

func Run() {
	for _, recipe := range AllRecipes {
		result := recipe.RunProbe()

		fmt.Printf("Policies:\n%s\n", explainer.TableExplainer(matcher.BuildNetworkPolicies(recipe.Policies())))

		fmt.Printf("resources:\n%s\n", recipe.Resources.Table())

		fmt.Printf("Results:\n%s\n", result.Table.Table())

		fmt.Printf("\n\n\n")
	}
}
