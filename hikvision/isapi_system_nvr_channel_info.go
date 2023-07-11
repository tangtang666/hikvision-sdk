package hikvision

import (
	"encoding/xml"
	"net/url"
)

type InputProxyChannelList struct {
	XMLName           xml.Name `xml:"InputProxyChannelList"`
	Text              string   `xml:",chardata"`
	Version           string   `xml:"version,attr"`
	Xmlns             string   `xml:"xmlns,attr"`
	Size              string   `xml:"size,attr"`
	InputProxyChannel []struct {
		Text                      string `xml:",chardata"`
		Version                   string `xml:"version,attr"`
		Xmlns                     string `xml:"xmlns,attr"`
		ID                        string `xml:"id"`
		Name                      string `xml:"name"`
		SourceInputPortDescriptor struct {
			Text                 string `xml:",chardata"`
			ProxyProtocol        string `xml:"proxyProtocol"`
			AddressingFormatType string `xml:"addressingFormatType"`
			IpAddress            string `xml:"ipAddress"`
			ManagePortNo         string `xml:"managePortNo"`
			SrcInputPort         string `xml:"srcInputPort"`
			UserName             string `xml:"userName"`
			StreamType           string `xml:"streamType"`
			DeviceID             string `xml:"deviceID"`
		} `xml:"sourceInputPortDescriptor"`
		EnableAnr string `xml:"enableAnr"`
	} `xml:"InputProxyChannel"`
}

// GetNvrChannel /ISAPI/ContentMgmt/InputProxy/channels  获取通道
func (c *Client) GetNvrChannel() (resp *InputProxyChannelList, err error) {
	path := "/ISAPI/ContentMgmt/InputProxy/channels"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}

	body, err := c.Get(u)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
