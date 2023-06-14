package hikvision

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

type StreamingChannel struct {
	XMLName     xml.Name `xml:"StreamingChannel"`
	Version     string   `xml:"version,attr"`
	Xmlns       string   `xml:"xmlns,attr"`
	ID          string   `xml:"id"`
	ChannelName string   `xml:"channelName"`
	Enabled     string   `xml:"enabled"`
	Transport   struct {
		MaxPacketSize       string `xml:"maxPacketSize"`
		ControlProtocolList struct {
			ControlProtocol []struct {
				StreamingTransport string `xml:"streamingTransport"`
			} `xml:"ControlProtocol"`
		} `xml:"ControlProtocolList"`
		Unicast struct {
			Enabled          string `xml:"enabled"`
			RtpTransportType string `xml:"rtpTransportType"`
		} `xml:"Unicast"`
		Multicast struct {
			Enabled         string `xml:"enabled"`
			DestIPAddress   string `xml:"destIPAddress"`
			VideoDestPortNo string `xml:"videoDestPortNo"`
			AudioDestPortNo string `xml:"audioDestPortNo"`
			FecInfo         struct {
				FecRatio      string `xml:"fecRatio"`
				FecDestPortNo string `xml:"fecDestPortNo"`
			} `xml:"FecInfo"`
		} `xml:"Multicast"`
		Security struct {
			Enabled         string `xml:"enabled"`
			CertificateType string `xml:"certificateType"`
		} `xml:"Security"`
	} `xml:"Transport"`
	Video struct {
		Enabled                 string `xml:"enabled"`
		VideoInputChannelID     string `xml:"videoInputChannelID"`
		VideoCodecType          string `xml:"videoCodecType"`
		VideoScanType           string `xml:"videoScanType"`
		VideoResolutionWidth    string `xml:"videoResolutionWidth"`
		VideoResolutionHeight   string `xml:"videoResolutionHeight"`
		VideoQualityControlType string `xml:"videoQualityControlType"`
		ConstantBitRate         string `xml:"constantBitRate"`
		FixedQuality            string `xml:"fixedQuality"`
		VbrUpperCap             string `xml:"vbrUpperCap"`
		VbrLowerCap             string `xml:"vbrLowerCap"`
		MaxFrameRate            string `xml:"maxFrameRate"`
		KeyFrameInterval        string `xml:"keyFrameInterval"`
		SnapShotImageType       string `xml:"snapShotImageType"`
		GovLength               string `xml:"GovLength"`
		SVC                     struct {
			Enabled string `xml:"enabled"`
		} `xml:"SVC"`
		PacketType  []string `xml:"PacketType"`
		Smoothing   string   `xml:"smoothing"`
		H265Profile string   `xml:"H265Profile"`
		SmartCodec  struct {
			Enabled string `xml:"enabled"`
		} `xml:"SmartCodec"`
	} `xml:"Video"`
	Audio struct {
		Enabled              string `xml:"enabled"`
		AudioInputChannelID  string `xml:"audioInputChannelID"`
		AudioCompressionType string `xml:"audioCompressionType"`
	} `xml:"Audio"`
}

func (c *Client) GetCameraInfoS() (resp []StreamingChannel, err error) {
	for i := 1; i < 4; i++ {
		res, err := c.GetCameraInfoByChannelId(i)
		if err == nil {
			resp = append(resp, *res)
		}
	}
	return resp, nil
}

func (c *Client) GetCameraInfoByChannelId(channelId int) (resp *StreamingChannel, err error) {
	path := fmt.Sprintf("/ISAPI/Streaming/channels/10%d", channelId)
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
