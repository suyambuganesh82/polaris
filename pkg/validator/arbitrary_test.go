// Copyright 2019 FairwindsOps Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"testing"

	conf "github.com/fairwindsops/polaris/pkg/config"
	"github.com/fairwindsops/polaris/pkg/kube"
	"github.com/stretchr/testify/assert"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestValidateOtherKind(t *testing.T) {
	c := conf.Configuration{
		Checks: map[string]conf.Severity{
			"pdbDisruptionsAllowedGreaterThanZero": conf.SeverityWarning,
		},
	}
	pdb := unstructured.Unstructured{}
	res, err := kube.NewGenericResourceFromUnstructured(&pdb)
	res.Kind = "PodDisruptionBudget"

	actualResult, err := applyNonControllerSchemaChecks(&c, res)
	if err != nil {
		panic(err)
	}
	results := actualResult.Results["pdbDisruptionsAllowedGreaterThanZero"]

	assert.False(t, results.Success)
	assert.Equal(t, conf.SeverityWarning, results.Severity)
	assert.Equal(t, "Reliability", results.Category)
	assert.EqualValues(t, "disruptionsAllowed is not greater than zero", results.Message)

	// tls := extv1beta1.IngressTLS{
	// Hosts:      []string{"test"},
	// SecretName: "secret",
	// }

	// ingress.Spec.TLS = []extv1beta1.IngressTLS{tls}
	// actualResult, err = ValidateIngress(&c, ingress)
	// if err != nil {
	// panic(err)
	// }
	// results = actualResult.Results["pdbDisruptionsAllowedGreaterThanZero"]

	// assert.True(t, results.Success)
	// assert.Equal(t, conf.Severity("warning"), results.Severity)
	// assert.Equal(t, "Security", results.Category)
	// assert.EqualValues(t, "Ingress has TLS configured", results.Message)

}