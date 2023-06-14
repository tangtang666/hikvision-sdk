package hikvision

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

type ONVIF struct {
	Enable bool `xml:"enable"`
}

// autocode https://tool.hiofd.com/xml-to-go/
type OnvifInfo struct {
	XMLName xml.Name `xml:"Integrate"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	ONVIF   ONVIF    `xml:"ONVIF"`
	ISAPI   ONVIF    `xml:"ISAPI"`
}
type UserList struct {
	XMLName xml.Name `xml:"UserList"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	User    []struct {
		Version  string `xml:"version,attr"`
		Xmlns    string `xml:"xmlns,attr"`
		ID       string `xml:"id"`
		UserName string `xml:"userName"`
		UserType string `xml:"userType"`
	} `xml:"User"`
}
type User struct {
	XMLName  xml.Name `xml:"User"`
	ID       int      `xml:"id"`
	UserName string   `xml:"userName"`
	Password string   `xml:"password"`
	UserType string   `xml:"userType"`
}

func (c *Client) GetOnvifStatus() (resp *OnvifInfo, err error) {
	path := "/ISAPI/System/Network/Integrate"
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

func (c *Client) EnableOnvif() (resp *ResponseStatus, err error) {
	path := "/ISAPI/System/Network/Integrate"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	payload := OnvifInfo{
		XMLName: xml.Name{},
		Version: "1.0",
		Xmlns:   "",
		ONVIF:   ONVIF{true},
		ISAPI:   ONVIF{true},
	}
	body, err := c.PutXML(u, payload)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) DisableOnvif() (resp *ResponseStatus, err error) {
	path := "/ISAPI/System/Network/Integrate"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	payload := OnvifInfo{
		XMLName: xml.Name{},
		Version: "1.0",
		Xmlns:   "",
		ONVIF:   ONVIF{false},
		ISAPI:   ONVIF{true},
	}
	body, err := c.PutXML(u, payload)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListOnvifUser /ISAPI/Security/ONVIF/users/ get获取出onvif列表
func (c *Client) ListOnvifUser() (resp *UserList, err error) {
	path := "/ISAPI/Security/ONVIF/users"
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

// AddOnvifUser /ISAPI/Security/ONVIF/users/ put 添加用户
func (c *Client) AddOnvifUser(id int, userName, password, userType string) (resp *ResponseStatus, err error) {
	path := "/ISAPI/Security/ONVIF/users"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	user := User{
		XMLName:  xml.Name{},
		ID:       id,
		UserName: userName,
		Password: password,
		UserType: userType,
	}
	body, err := c.PutXML(u, user)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteOnvifUser /ISAPI/Security/ONVIF/users/ delete 删除用户
func (c *Client) DeleteOnvifUser(id int) (resp *ResponseStatus, err error) {
	path := fmt.Sprintf("/ISAPI/Security/ONVIF/users/%d", id)
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}

	body, err := c.Delete(u)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
