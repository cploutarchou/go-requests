package http

import "net/http"

type Headers interface {
	SetHeaders(http.Header) Headers
	GetHeaders() http.Header
	ApplicationJSON() Headers
	ApplicationXML() Headers
	ApplicationForm() Headers
	ApplicationOctetStream() Headers
	ApplicationProtobuf() Headers
	ApplicationMsgpack() Headers
	ApplicationYaml() Headers
	ApplicationText() Headers
	ApplicationHtml() Headers
	ApplicationJavascript() Headers
	ApplicationPdf() Headers
	ApplicationZip() Headers
	ApplicationGzip() Headers
	ApplicationTar() Headers
	ApplicationRar() Headers
	Application7z() Headers
	ApplicationXlsx() Headers
	ApplicationDocx() Headers
	ApplicationPptx() Headers
	ApplicationEpub() Headers
	ApplicationMobi() Headers
	ApplicationJsonld() Headers
	ApplicationRss() Headers
	ApplicationAtom() Headers
	ApplicationWebmanifest() Headers
	ApplicationWebp() Headers
	ApplicationAvif() Headers
	ApplicationJpeg() Headers
	ApplicationPng() Headers
	ApplicationGif() Headers
	ApplicationSvg() Headers
	ApplicationTiff() Headers
	ApplicationBmp() Headers
	ApplicationIco() Headers
	ApplicationFlac() Headers
	ApplicationMp3() Headers
	ApplicationM4a() Headers
	ApplicationOgg() Headers
	ApplicationWav() Headers
	ApplicationWma() Headers
	ApplicationAac() Headers
	ApplicationXls() Headers
	ApplicationDoc() Headers
	ApplicationPpt() Headers
	ApplicationEpub2() Headers
	ApplicationMobi2() Headers
	ApplicationJsonld2() Headers
	ApplicationXlsx2() Headers
}

type headersImpl struct {
	header http.Header
}

func newHeadersImpl() *headersImpl {
	return &headersImpl{header: make(http.Header)}
}

func (c headersImpl) SetHeaders(h http.Header) {
	c.header = h

}
func (c headersImpl) GetHeaders() http.Header {
	return c.header
}
