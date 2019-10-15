package dms

import (
	"encoding/json"
	"github.com/anacrolix/dms/upnpav"
	"mime"
	"path"
	"strings"
)

func newFshareCdsObject() []upnpav.Item {
	decoder := json.NewDecoder(strings.NewReader(sampleJSON))
	var m []*Movie
	decoder.Decode(&m)
	items := make([]upnpav.Item, len(m))
	for i, movie := range m {
		caption := determinateCaption(movie.Links)
		object := upnpav.Object{
			Restricted: 0,
			ParentID: "-1",
			Searchable: 0,
			Class: "object.item.videoItem.movie",

			ID: movie.Title,
			Title: movie.Title,
			AlbumArtURI: movie.Image,
		}

		mimeType := mime.TypeByExtension(path.Ext(movie.Links[0].Title))
		item := upnpav.Item{
			Object: object,
			Res: []upnpav.Resource{
				{
					ProtocolInfo: "http-get:*:" + mimeType + ":DLNA.ORG_OP=01;DLNA.ORG_FLAGS=01700000000000000000000000000000",
					URL: fshareToLocalLink(movie.Links[0].Link),
					Subtitle: caption,
					SubtitleFileType: "SRT",
					PV: "http://www.pv.com/pvns/",
				},
				{
					ProtocolInfo: "http-get:*:smi/caption:",
					URL: caption,
				},
			},
		}



		item.SetCaption(caption)
		items[i] = item
	}
	return items
}

func determinateCaption(links []*Link) string {
	for _, l := range links {
		if strings.HasSuffix(l.Title, ".srt") {
			return fshareToLocalLink(l.Link)
		}
	}
	return ""
}

func fshareToLocalLink(fshareLink string) string {
	fshareLink = strings.Replace(fshareLink, "http://", "https://", 1)
	return strings.Replace(fshareLink, "https://www.fshare.vn/file/", "http://192.168.1.111/fshare/", 1)
}

type Movie struct {
	Title string `json:"title"`
	Image string `json:"image"`
	ID    int64  `json:"id"`
	Links []*Link `json:"links"`
}

type Link struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}


var sampleJSON = `
[
  {
    "title": "Tam Đại Sư Phụ && Three Kims",
    "image": "https://www.thuvienaz.net/wp-content/uploads/2019/10/MV5BNDllMjRjOWYtZDE0Mi00NjY1LTgyZjAtYzJmOTRkMGIzNzc1XkEyXkFqcGdeQXVyMzE4MDkyNTA@._V1_SX1024_AL_.jpg",
    "id": 131878,
    "links": [
      {
        "title": "Three.Kims.2007.KOREAN.1080p.BluRay.H264.AAC-VXT.mp4",
        "link": "https://www.fshare.vn/file/ZMXJL5K6XU9F1NX"
      },
      {
        "title": "Three.Kims.2007.KOREAN.1080p.BluRay.H264.AAC-VXT.srt",
        "link": "http://www.fshare.vn/file/XXPGUL82DCB8"
      }
    ]
  },
  {
    "title": "Câu Chuyện Tokyo && Tokyo Story / Tokyo Monogatari",
    "image": "https://www.thuvienaz.net/wp-content/uploads/2019/10/MV5BM2E1ZmZlMDQtMDgzNi00MGVmLTgwMWMtYTc1MjkzOWJmOGIxXkEyXkFqcGdeQXVyMjgyNjk3MzE@._V1_SX1024_AL_.jpg",
    "id": 131872,
    "links": [
      {
        "title": "Tokyo.Story.1953.720p.BluRay.x264-CiNEFiLE.Vietsub.mkv",
        "link": "https://www.fshare.vn/file/26E8QNUZ7PP6GFK"
      },
      {
        "title": "Tokyo.Story.1953.REMASTERED.1080p.BluRay.H264.AAC-VXT.mp4",
        "link": "https://www.fshare.vn/file/SIVL5NH9KRW12X9"
      },
      {
        "title": "Tokyo.Story.aka.Tokyo.monogatari.1953.Criterion.Collection.Bluray.Remux.AVC.LPCM.1.0-RuTracker.mkv",
        "link": "https://www.fshare.vn/file/3SJ18RY7UFAVBMX"
      },
      {
        "title": "Tokyo.Story.aka.Tokyo.monogatari.1953.Criterion.Collection.Bluray.Remux.AVC.LPCM.1.0-RuTracker.srt",
        "link": "https://www.fshare.vn/file/IIR3RD7OKHHPVXQ"
      }
    ]
  },
  {
    "title": "Trứng Thiên Thần && Angel's Egg",
    "image": "https://www.thuvienaz.net/wp-content/uploads/2019/10/MV5BMWM2N2RjOTAtYzRkNi00NDY5LWFhODAtZWEwYjdiM2FlMjljXkEyXkFqcGdeQXVyNjUwNzk3NDc@._V1_SX1024_AL_.jpg",
    "id": 131865,
    "links": [
      {
        "title": "Angels.Egg.1985.JAPANESE.1080p.BluRay.H264.AAC-VXT.mp4",
        "link": "https://www.fshare.vn/file/1KE2UR3RB22VXRT"
      },
      {
        "title": "Angels.Egg.1985.JAPANESE.1080p.BluRay.H264.AAC-VXT.srt",
        "link": "http://www.fshare.vn/file/64SFEFNQHY5J"
      }
    ]
  },
  {
    "title": "Hai Phóng Viên Đặc Biệt && Special Correspondents",
    "image": "https://www.thuvienaz.net/wp-content/uploads/2019/10/MV5BYzBhMDNhZDktMDQ1Yi00ZGEyLThjYzYtZWZiNDRhMjE5ZGEwXkEyXkFqcGdeQXVyNjUwNzk3NDc@._V1_SX1024_AL_.jpg",
    "id": 131858,
    "links": [
      {
        "title": "Special.Correspondents.2016.1080p.WEBRip.x264-RARBG.mp4",
        "link": "https://www.fshare.vn/file/ZWDTR2P48F8KXN9"
      },
      {
        "title": "Special.Correspondents.2016.1080p.WEBRip.x264-RARBG.srt",
        "link": "http://www.fshare.vn/file/E3X83DFOU18T"
      }
    ]
  },
  {
    "title": "Băng Đảng Đường Phố && Street Flow / Banlieusards",
    "image": "https://www.thuvienaz.net/wp-content/uploads/2019/10/MV5BYTljODZkZTItNTU1OS00NmZkLThlNzMtMDhiMjI0M2Y1ZTk1XkEyXkFqcGdeQXVyODIyOTEyMzY@._V1_SX1024_AL_.jpg",
    "id": 131855,
    "links": [
      {
        "title": "Banlieusards.2019.1080p.NF.WEB-DL.DDP5.1.H264-CMRG.mkv",
        "link": "https://www.fshare.vn/file/JJS4WWLYQJAHARM"
      },
      {
        "title": "Banlieusards.2019.1080p.NF.WEB-DL.DDP5.1.H264-CMRG.srt",
        "link": "http://www.fshare.vn/file/AYSSBWCPA1YJ"
      }
    ]
  }
]
`