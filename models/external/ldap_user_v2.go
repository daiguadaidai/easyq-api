package external

import (
	"encoding/json"
	"fmt"
	"strings"
)

type LdapUserV2 struct {
	DN              string `json:"dn"`
	Uid             string `json:"uid"`
	Mail            string `json:"mail"`
	TelephoneNumber string `json:"telephoneNumber"`
	EmployeeNumber  string `json:"employeeNumber"`
	DisplayName     string `json:"displayName"`
	Mobile          string `json:"mobile"`
	SN              string `json:"sn"`
	CN              string `json:"cn"`
	ObjectClass     string `json:"objectClass"`
	UserPassword    string `json:"userPassword"`
}

func NewLdapUserV2(fieldValueMap map[string]string) (*LdapUserV2, error) {
	raw, err := json.Marshal(fieldValueMap)
	if err != nil {
		return nil, fmt.Errorf("LdapUserV2 map->json 出错. %v", err.Error())
	}

	var ldapUser LdapUserV2
	if err := json.Unmarshal(raw, &ldapUser); err != nil {
		return nil, fmt.Errorf("LdapUserV2 json->LdapUser 出错. %v", err.Error())
	}

	return &ldapUser, nil
}

func (this *LdapUserV2) GetNameZh() string {
	return strings.TrimRight(this.CN, fmt.Sprintf("#%s", this.Uid))
}

func (this *LdapUserV2) GetNameEn() string {
	return strings.TrimRight(this.Mail, "@yonghui.cn")
}
