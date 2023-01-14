package main

import (
	"flag"
	"monish/ADReaper/ldapquery"
	"os"
	"strings"
)

// Main function

func main() {

	commandStr := "\nCommand(s) to run\n"
	commandStr += "\n\tdc              - to list domain controllers"
	commandStr += "\n\tdomain-trust    - to list domain trust"
	commandStr += "\n\tusers           - to list all users"
	commandStr += "\n\tcomputers       - to list all computers"
	commandStr += "\n\tgroups          - to list all groups with members"
	commandStr += "\n\tspn             - to list service principal objects"
	commandStr += "\n\tnever-loggedon  - to list users never logged on"
	commandStr += "\n\tgpo             - to list group policy objects"
	commandStr += "\n\tou              - to list organizational units"
	commandStr += "\n\tms-sql          - to list MS-SQL servers"
	commandStr += "\n\tasreproast      - to list AS-REP roastable accounts"
	commandStr += "\n\tunconstrained   - to list Unconstrained Delegated accounts"

	commandStr += "\n\tadmin-priv      - to list AD objects with admin privilege"
	commandStr += "\n"

	filterMsg := "\nFilters to use for users/groups/computers\n"
	filterMsg += "\nlist - lists all objects only"
	filterMsg += "\nfull-data - list all objects with properties"
	filterMsg += "\nmembership - lists all members from an object"
	filterMsg += "\n\n"

	nameMsg := "\nPass object name of user/group/computer"
	nameMsg += "\n"

	// helpMsg := commandStr + filterMsg + commandStr

	var cmd = map[string]bool{
		"users":          true,
		"never-loggedon": true,
		"groups":         true,
		"computers":      true,
		"dc":             true,
		"gpo":            true,
		"spn":            true,
		"admin-priv":     true,
		"domain-trust":   true,
		"ou":             true,
		"ms-sql":         true,
		"asreproast":     true,
		"unconstrained":  true,
	}

	//getting args data
	ldapServer := flag.String("dc", "", "\nEnter the DC\n")
	ldapBind := flag.String("user", "", "\nEnter the Username\n")
	ldapPassword := flag.String("password", "", "\nEnter the Password\n")
	commands := flag.String("command", "", commandStr)
	filter := flag.String("filter", "list", filterMsg)
	name := flag.String("name", "", nameMsg)
	flag.Parse()

	// get list of all commands by splitting on "," or " "
	var commandList []string

	if strings.Contains(*commands, ",") {
		commandList = strings.Split(*commands, ",")
	} else {
		commandList = strings.Split(*commands, " ")
	}

	for _,command := range commandList{
		if !(cmd[command]) || len(command) == 0 {
			flag.PrintDefaults()
			os.Exit(-1)
		}
	}


	if len(*ldapServer) == 0 || len(*ldapBind) == 0 || len(*ldapPassword) == 0 {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	//formatting data
	s := strings.Split(*ldapServer, ".")
	baseDN := ""
	for x := 1; x < len(s); x++ {
		if x == len(s)-1 {
			baseDN += "DC=" + s[x]
		} else {
			baseDN += "DC=" + s[x] + ","
		}
	}
	*ldapServer += ":389"

	//Query LDAP
	// out, err := exec.Command("nltest", "/DSGETDC:", "/LDAPONLY").Output()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(out))

	//querying
	for _,command := range commandList {
		ldapquery.LDAPlistquery(*ldapServer, *ldapBind, *ldapPassword, baseDN, command, *filter, *name)
	}
}
