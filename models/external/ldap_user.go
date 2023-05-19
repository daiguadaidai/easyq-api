package external

import (
	"encoding/json"
	"fmt"
)

type LdapUser struct {
	DN            string `json:"dn"`
	UserPassword  string `json:"userPassword"`
	GivenName     string `json:"givenName"`
	HomeDirectory string `json:"homeDirectory"`
	UidNumber     string `json:"uidNumber"`
	GidNumber     string `json:"gidNumber"`
	SN            string `json:"sn"`
	Mail          string `json:"mail"`
	Uid           string `json:"uid"`
	CN            string `json:"cn"`
}

func NewLdapUser(fieldValueMap map[string]string) (*LdapUser, error) {
	raw, err := json.Marshal(fieldValueMap)
	if err != nil {
		return nil, fmt.Errorf("LdapUser map->json 出错. %v", err.Error())
	}

	var ldapUser LdapUser
	if err := json.Unmarshal(raw, &ldapUser); err != nil {
		return nil, fmt.Errorf("LdapUser json->LdapUser 出错. %v", err.Error())
	}

	return &ldapUser, nil
}
