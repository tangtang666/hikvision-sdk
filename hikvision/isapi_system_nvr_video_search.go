package hikvision

import (
	"encoding/xml"
	"github.com/google/uuid"
	"net/url"
	"strings"
)

type TrackList struct {
	TrackID string `xml:"trackID"`
}
type TimeSpan struct {
	StartTime string `xml:"startTime"`
	EndTime   string `xml:"endTime"`
}
type TimeSpanList struct {
	TimeSpan TimeSpan `xml:"timeSpan"`
}

type MetadataList struct {
	MetadataDescriptor string `xml:"metadataDescriptor"`
}
type CMSearchDescription struct {
	XMLName             xml.Name     `xml:"CMSearchDescription"`
	SearchID            string       `xml:"searchID"`
	TrackList           TrackList    `xml:"trackList"`
	MetadataList        MetadataList `xml:"metadataList"`
	TimeSpanList        TimeSpanList `xml:"timeSpanList"`
	MaxResults          string       `xml:"maxResults"`
	SearchResultPostion string       `xml:"searchResultPostion"`
}
type CMSearchResult struct {
	XMLName            xml.Name `xml:"CMSearchResult"`
	Version            string   `xml:"version,attr"`
	Xmlns              string   `xml:"xmlns,attr"`
	SearchID           string   `xml:"searchID"`
	ResponseStatus     string   `xml:"responseStatus"`
	ResponseStatusStrg string   `xml:"responseStatusStrg"`
	NumOfMatches       string   `xml:"numOfMatches"`
	MatchList          struct {
		SearchMatchItem []struct {
			SourceID string `xml:"sourceID"`
			TrackID  string `xml:"trackID"`
			TimeSpan struct {
				StartTime string `xml:"startTime"`
				EndTime   string `xml:"endTime"`
			} `xml:"timeSpan"`
			MediaSegmentDescriptor struct {
				ContentType string `xml:"contentType"`
				CodecType   string `xml:"codecType"`
				PlaybackURI string `xml:"playbackURI"`
			} `xml:"mediaSegmentDescriptor"`
			MetadataMatches struct {
				MetadataDescriptor string `xml:"metadataDescriptor"`
			} `xml:"metadataMatches"`
		} `xml:"searchMatchItem"`
	} `xml:"matchList"`
}

// SearchVideoByTimeAndName ISAPI/ContentMgmt/search  按时间和名称搜索
func (c *Client) SearchVideoByTimeAndName(name string, startTime string, endTime string) (resp *CMSearchResult, err error) {
	path := "/ISAPI/ContentMgmt/search"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	uuid := uuid.New()
	uuidString := strings.ToUpper(uuid.String())
	payload := CMSearchDescription{
		XMLName:      xml.Name{},
		SearchID:     uuidString,
		TrackList:    TrackList{name},
		MetadataList: MetadataList{"//metadata.ISAPI.org/VideoMotion"},
		TimeSpanList: TimeSpanList{TimeSpan{
			StartTime: startTime,
			EndTime:   endTime,
		}},
		MaxResults:          "40",
		SearchResultPostion: "0",
	}
	body, err := c.PostXML(u, payload)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body))
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
