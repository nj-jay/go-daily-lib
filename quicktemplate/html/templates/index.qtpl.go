// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/index.qtpl:1
package templates

//line templates/index.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/index.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/index.qtpl:1
func StreamIndex(qw422016 *qt422016.Writer, name string) {
//line templates/index.qtpl:1
	qw422016.N().S(`
<html>
  <head>
    <title>Awesome Web</title>
  </head>
  <body>
    <h1>Hi, `)
//line templates/index.qtpl:7
	qw422016.E().S(name)
//line templates/index.qtpl:7
	qw422016.N().S(`
    <p>Welcome to the awesome web!!!</p>
  </body>
</html>
`)
//line templates/index.qtpl:11
}

//line templates/index.qtpl:11
func WriteIndex(qq422016 qtio422016.Writer, name string) {
//line templates/index.qtpl:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/index.qtpl:11
	StreamIndex(qw422016, name)
//line templates/index.qtpl:11
	qt422016.ReleaseWriter(qw422016)
//line templates/index.qtpl:11
}

//line templates/index.qtpl:11
func Index(name string) string {
//line templates/index.qtpl:11
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/index.qtpl:11
	WriteIndex(qb422016, name)
//line templates/index.qtpl:11
	qs422016 := string(qb422016.B)
//line templates/index.qtpl:11
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/index.qtpl:11
	return qs422016
//line templates/index.qtpl:11
}
