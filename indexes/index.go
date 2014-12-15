//go:generate go run internal/cmd/generator.go -tmpl=./internal/templates/ellipsoid_index.tmpl -out=ellipsoid_index.go -type=Ellipsoid
//go:generate gofmt -w ellipsoid_index.go

package indexes
