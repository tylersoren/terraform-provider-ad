package ad

import ldap "gopkg.in/ldap.v2"

func addGroupToAD(groupName string, dnName string, adConn *ldap.Conn, desc string) error {
	addRequest := ldap.NewAddRequest(dnName)
	addRequest.Attribute("objectClass", []string{"group"})
	addRequest.Attribute("sAMAccountName", []string{groupName})
	if desc != "" {
		addRequest.Attribute("description", []string{desc})
	}
	err := adConn.Add(addRequest)
	if err != nil {
		return err
	}
	return nil
}

func deleteGroupFromAD(dnName string, adConn *ldap.Conn) error {
	delRequest := ldap.NewDelRequest(dnName, nil)
	err := adConn.Del(delRequest)
	if err != nil {
		return err
	}
	return nil
}
