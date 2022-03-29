package gensnippets

import (
	"fmt"
	"strings"

	"cloud.google.com/go/internal/gapicgen/gensnippets/metadata"
)

type apiInfo struct {
	// protoPkg is the proto namespace for the API package.
	protoPkg string
	// libPkg is the gapic import path.
	libPkg string
	// protoServices are the list of services associated to the proto package.
	protoServices []*service
	// version is the Go module version for the gapic client.
	version string
	// shortName for the service.
	shortName string
}

// RegionTags gets the region tags keyed by client name and method name.
func (ai *apiInfo) RegionTags() map[string]map[string]string {
	regionTags := map[string]map[string]string{}
	for _, svc := range ai.protoServices {
		regionTags[svc.name] = map[string]string{}
		for _, r := range svc.methods {
			regionTags[svc.name][r.name] = r.regionTag
		}
	}
	return regionTags
}

// RegionTags gets the region tags keyed by client name and method name.
func (ai *apiInfo) ToSnippetMetadata() *metadata.Index {
	index := &metadata.Index{
		ClientLibrary: &metadata.ClientLibrary{
			Name:     ai.libPkg,
			Version:  ai.version,
			Language: metadata.Language_GO,
			Apis: []*metadata.Api{
				{
					Id:      ai.protoPkg,
					Version: ai.protoVersion(),
				},
			},
		},
	}

	for _, service := range ai.protoServices {
		for _, method := range service.methods {
			snip := &metadata.Snippet{
				RegionTag:   method.regionTag,
				Title:       fmt.Sprintf("%s %s Sample", ai.shortName, method.name),
				Description: method.doc,
				File:        fmt.Sprintf("%s/%s/main.go", service.name, method.name),
				Language:    metadata.Language_GO,
				Canonical:   false,
				Origin:      *metadata.Snippet_API_DEFINITION.Enum(),
				// TODO: segments
				ClientMethod: &metadata.ClientMethod{
					ShortName:  method.name,
					FullName:   fmt.Sprintf("%s.%s.%s", ai.protoPkg, strings.Title(ai.shortName), method.name),
					Async:      false,
					ResultType: method.result,
					Client: &metadata.ServiceClient{
						ShortName: service.name,
						FullName:  fmt.Sprintf("%s.%s", ai.protoPkg, service.name),
					},
					Method: &metadata.Method{
						ShortName: method.name,
						FullName:  fmt.Sprintf("%s.%s.%s", ai.protoPkg, service.name, method.name),
						Service: &metadata.Service{
							ShortName: service.protoName,
							FullName:  fmt.Sprintf("%s.%s", ai.protoPkg, service.protoName),
						},
					},
				},
			}
			for _, param := range method.params {
				methParam := &metadata.ClientMethod_Parameter{
					Type: param.pType,
					Name: param.pType,
				}
				snip.ClientMethod.Parameters = append(snip.ClientMethod.Parameters, methParam)
			}
			index.Snippets = append(index.Snippets, snip)
		}
	}
	return index
}

func (ai *apiInfo) protoVersion() string {
	ss := strings.Split(ai.protoPkg, ".")
	return ss[len(ss)-1]
}

type service struct {
	// protoName is the name of the proto service.
	protoName string
	// name is the name of the corresponding gapic client for the proto service.
	name string
	// methods are the list of methods associated to the gapic client.
	methods []*method
}

type method struct {
	// name is the name of the method.
	name string
	// doc is the documention for the methods
	doc string
	// regionTag is the region tag that will be used for the generated snippet.
	regionTag string
	//params are the input parameters for the gapic method.
	params []*param
	// result is the return value for the method.
	result string
}

type param struct {
	// name of the parameter.
	name string
	// pType is the Go type for the parameter.
	pType string
}
