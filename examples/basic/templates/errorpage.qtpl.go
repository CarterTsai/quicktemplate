// This file is automatically generated by qtc from "templates/errorpage.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/errorpage.qtpl:1
package templates

//line templates/errorpage.qtpl:1
import (
	"io"

	"github.com/valyala/quicktemplate"
)

// Error page template. Implements BasePage methods.
//

//line templates/errorpage.qtpl:3
var (
	_ = io.Copy
	_ = quicktemplate.AcquireByteBuffer
)

//line templates/errorpage.qtpl:4
type ErrorPage struct {
	// error path
	Path []byte
}

//line templates/errorpage.qtpl:11
func (p *ErrorPage) StreamTitle(qw *quicktemplate.Writer) {
	//line templates/errorpage.qtpl:11
	qw.N().S(`
	Error
`)
//line templates/errorpage.qtpl:13
}

//line templates/errorpage.qtpl:13
func (p *ErrorPage) WriteTitle(qww io.Writer) {
	//line templates/errorpage.qtpl:13
	qw := quicktemplate.AcquireWriter(qww)
	//line templates/errorpage.qtpl:13
	p.StreamTitle(qw)
	//line templates/errorpage.qtpl:13
	quicktemplate.ReleaseWriter(qw)
//line templates/errorpage.qtpl:13
}

//line templates/errorpage.qtpl:13
func (p *ErrorPage) Title() string {
	//line templates/errorpage.qtpl:13
	qb := quicktemplate.AcquireByteBuffer()
	//line templates/errorpage.qtpl:13
	p.WriteTitle(qb)
	//line templates/errorpage.qtpl:13
	qs := string(qb.B)
	//line templates/errorpage.qtpl:13
	quicktemplate.ReleaseByteBuffer(qb)
	//line templates/errorpage.qtpl:13
	return qs
//line templates/errorpage.qtpl:13
}

//line templates/errorpage.qtpl:16
func (p *ErrorPage) StreamBody(qw *quicktemplate.Writer) {
	//line templates/errorpage.qtpl:16
	qw.N().S(`
	<h1>Error page</h1>
	</div>
		Unsupported path <b>`)
	//line templates/errorpage.qtpl:19
	qw.E().Z(p.Path)
	//line templates/errorpage.qtpl:19
	qw.N().S(`</b>.
	</div>
`)
//line templates/errorpage.qtpl:21
}

//line templates/errorpage.qtpl:21
func (p *ErrorPage) WriteBody(qww io.Writer) {
	//line templates/errorpage.qtpl:21
	qw := quicktemplate.AcquireWriter(qww)
	//line templates/errorpage.qtpl:21
	p.StreamBody(qw)
	//line templates/errorpage.qtpl:21
	quicktemplate.ReleaseWriter(qw)
//line templates/errorpage.qtpl:21
}

//line templates/errorpage.qtpl:21
func (p *ErrorPage) Body() string {
	//line templates/errorpage.qtpl:21
	qb := quicktemplate.AcquireByteBuffer()
	//line templates/errorpage.qtpl:21
	p.WriteBody(qb)
	//line templates/errorpage.qtpl:21
	qs := string(qb.B)
	//line templates/errorpage.qtpl:21
	quicktemplate.ReleaseByteBuffer(qb)
	//line templates/errorpage.qtpl:21
	return qs
//line templates/errorpage.qtpl:21
}
