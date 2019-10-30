
package api
import(
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)
//Register register router and start api server
func Register() {
	u := UserResource{map[string]User{}}
	restful.DefaultContainer.Add(u.WebService())

	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("swagger-ui/dist"))))

	log.Printf("start listening on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "UserService",
			Description: "Resource for managing Users",
			Contact: &spec.ContactInfo{
				Name:  "panguxs",
				Email: "support@panguxs.io",
				URL:   "https://panguxs.io",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "users",
		Description: "Managing users"}}}
}