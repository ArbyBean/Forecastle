package ingressapps

import (
	"github.com/stakater/Forecastle/pkg/annotations"
	"github.com/stakater/Forecastle/pkg/util/strings"

	"github.com/stakater/Forecastle/pkg/config"
	v1 "k8s.io/api/networking/v1"
)

// For filtering ingressing having forecastle expose annotation set to true
func byForecastleExposeAnnotation(ingress v1.Ingress, appConfig config.Config) bool {
	if val, ok := ingress.Annotations[annotations.ForecastleExposeAnnotation]; ok {
		// Has Forecastle annotation and is exposed
		if val != "false" {
			return true
		}
	}
	return true
}

// For filtering ingresses by forecastle instance
func byForecastleInstanceAnnotation(ingress v1.Ingress, appConfig config.Config) bool {
	if val, ok := ingress.Annotations[annotations.ForecastleInstanceAnnotation]; ok {
		return strings.ContainsBetweenDelimiter(val, appConfig.InstanceName, ",")
	}
	return false
}
