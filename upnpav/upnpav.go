package upnpav

import (
	"encoding/xml"
)

const (
	NoSuchObjectErrorCode = 701
)

type Resource struct {
	XMLName          xml.Name `xml:"res"`
	ProtocolInfo     string   `xml:"protocolInfo,attr"`
	URL              string   `xml:",chardata"`
	Size             uint64   `xml:"size,attr,omitempty"`
	Bitrate          uint     `xml:"bitrate,attr,omitempty"`
	Duration         string   `xml:"duration,attr,omitempty"`
	Resolution       string   `xml:"resolution,attr,omitempty"`
	Subtitle         string   `xml:"pv:subtitleFileUri,attr,omitempty"`
	SubtitleFileType string   `xml:"pv:subtitleFileType,attr,omitempty"`
	PV               string   `xml:"xmlns:pv,attr,omitempty"`
}

type captionInfo struct {
	Type string `xml:"sec:type,attr"`
	URL  string `xml:",chardata"`
}
type CaptionInfo struct {
	captionInfo
	XMLName xml.Name `xml:"sec:CaptionInfo,omitempty"`
}

type CaptionInfoEx struct {
	captionInfo
	XMLName xml.Name `xml:"sec:CaptionInfoEx,omitempty"`
}

type Container struct {
	Object
	XMLName    xml.Name `xml:"container"`
	ChildCount int      `xml:"childCount,attr"`
}

type Item struct {
	Object
	XMLName   xml.Name `xml:"item"`
	Res       []Resource
	Caption   *CaptionInfo   `xml:",omitempty"`
	CaptionEx *CaptionInfoEx `xml:"sec:CaptionInfoEx,omitempty"`
}

type Object struct {
	ID          string `xml:"id,attr"`
	ParentID    string `xml:"parentID,attr"`
	Restricted  int    `xml:"restricted,attr"` // indicates whether the object is modifiable
	Class       string `xml:"upnp:class"`
	Icon        string `xml:"upnp:icon,omitempty"`
	Title       string `xml:"dc:title"`
	Artist      string `xml:"upnp:artist,omitempty"`
	Album       string `xml:"upnp:album,omitempty"`
	Genre       string `xml:"upnp:genre,omitempty"`
	AlbumArtURI string `xml:"upnp:albumArtURI,omitempty"`
	Searchable  int    `xml:"searchable,attr"`
}

func (c *Item) SetCaption(cap string) {
	if cap == "" {
		return
	}

	c.Caption = &CaptionInfo{
		captionInfo: captionInfo{URL: cap, Type: "srt"},
	}
	c.CaptionEx = &CaptionInfoEx{
		captionInfo: captionInfo{URL: cap, Type: "srt"},
	}
	c.Res = append(c.Res, Resource{
		ProtocolInfo: "http-get:*:text/srt:*",
		URL:          cap,
	}, Resource{
		ProtocolInfo: "http-get:*:text/srt:",
		URL:          cap,
	})
}
