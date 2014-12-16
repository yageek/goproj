//go:generate go run internal/cmd/generator.go -tmpl=./internal/templates/ellipsoid_index.tmpl -out=ellipsoid_index.go -type=Ellipsoid
//go:generate gofmt -w ellipsoid_index.go

//go:generate go run internal/cmd/generator.go -tmpl=./internal/templates/meridian_index.tmpl -out=meridian_index.go -type=Meridian
//go:generate gofmt -w meridian_index.go

package indexes
