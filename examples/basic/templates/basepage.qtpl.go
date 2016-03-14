// This file is automatically generated by qtc from "templates/basepage.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/basepage.qtpl:1
package templates

//line templates/basepage.qtpl:1
import (
	"io"

	"github.com/valyala/quicktemplate"
)

// This is a base page template. All the other template pages are inherited
// from this one.
//

//line templates/basepage.qtpl:4
var (
	_ = io.Copy
	_ = quicktemplate.AcquireByteBuffer
)

//line templates/basepage.qtpl:5
type BasePage interface {
	//line templates/basepage.qtpl:5
	Title() string
	//line templates/basepage.qtpl:5
	StreamTitle(qw *quicktemplate.Writer)
	//line templates/basepage.qtpl:5
	WriteTitle(qww io.Writer)
	//line templates/basepage.qtpl:5
	Body() string
	//line templates/basepage.qtpl:5
	StreamBody(qw *quicktemplate.Writer)
	//line templates/basepage.qtpl:5
	WriteBody(qww io.Writer)
//line templates/basepage.qtpl:5
}

// Page prints a page implementing BasePage interface.

//line templates/basepage.qtpl:13
func StreamPage(qw *quicktemplate.Writer, p BasePage) {
	//line templates/basepage.qtpl:13
	qw.N().S(`
<html>
	<head>
		<title>`)
	//line templates/basepage.qtpl:16
	p.StreamTitle(qw)
	//line templates/basepage.qtpl:16
	qw.N().S(`</title>
	</head>
	<body>
		<div>
			<a href="/">return to main page</a>
		</div>
		`)
	//line templates/basepage.qtpl:22
	p.StreamBody(qw)
	//line templates/basepage.qtpl:22
	qw.N().S(`
	</body>
</html>
`)
//line templates/basepage.qtpl:25
}

//line templates/basepage.qtpl:25
func WritePage(qww io.Writer, p BasePage) {
	//line templates/basepage.qtpl:25
	qw := quicktemplate.AcquireWriter(qww)
	//line templates/basepage.qtpl:25
	StreamPage(qw, p)
	//line templates/basepage.qtpl:25
	quicktemplate.ReleaseWriter(qw)
//line templates/basepage.qtpl:25
}

//line templates/basepage.qtpl:25
func Page(p BasePage) string {
	//line templates/basepage.qtpl:25
	qb := quicktemplate.AcquireByteBuffer()
	//line templates/basepage.qtpl:25
	WritePage(qb, p)
	//line templates/basepage.qtpl:25
	qs := string(qb.B)
	//line templates/basepage.qtpl:25
	quicktemplate.ReleaseByteBuffer(qb)
	//line templates/basepage.qtpl:25
	return qs
//line templates/basepage.qtpl:25
}
