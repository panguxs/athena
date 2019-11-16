package api

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	log "gopkg.in/logger.v1"
)

// FileResource is the REST layer to the File domain
type FileResource struct {
}

//File file struct
type File struct {
	Name    string    `json:"name,omitempty"`
	IsDir   bool      `json:"is_dir,omitempty"`
	ModTime time.Time `json:"mod_time,omitempty"`
	Mode    string    `json:"mode,omitempty"`
}

// WebService creates a new service that can handle REST requests for File resources.
func (f FileResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/files").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	tags := []string{"files"}
	ws.Route(ws.GET("/").To(f.list).
		// docs
		Doc("get all files").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []File{}))

	return ws
}

func (f FileResource) list(request *restful.Request, response *restful.Response) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	nFiles := make([]*File, 0)

	for _, f := range files {

		file := &File{}

		fmt.Println(f.Name())
		file.Name = f.Name()
		file.IsDir = f.IsDir()
		log.Info(f.Mode())
		log.Info(f.ModTime())
		file.ModTime = f.ModTime()
		file.Mode = f.Mode().String()

		nFiles = append(nFiles, file)
	}
	response.WriteEntity(nFiles)
}
