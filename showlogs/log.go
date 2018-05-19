package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

const f = "./data.json"

//var r = flag.String("r", "false", "Use -r <A or B>")
//var t = flag.String("t", "false", "Use -t <nginx or project>")

var qs = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Select{
			Message: "Choose A or B",
			Options: []string{"A", "B"},
			Default: "bjtb",
		},
	},
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "Choose a type:",
			Options: []string{"nginx", "project"},
			Default: "project",
		},
	},
}

func Choose() (Room, Class string) {
	answers := struct {
		Name          string `survey:"name"`
		FavoriteColor string `survey:"type"`
	}{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	return answers.Name, answers.FavoriteColor
}

func ConnHost(host string) {
	check := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}
	con := host + ":22"
	client, err := ssh.Dial("tcp", con, &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{ssh.Password("password")},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})
	check(err, "dial")

	session, err := client.NewSession()
	check(err, "new session")
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("xterm", 25, 80, modes)
	check(err, "request pty")

	err = session.Shell()
	check(err, "start shell")

	err = session.Wait()
	check(err, "return")
}

func Select() {
	arg := os.Args[1]
	room, class := Choose()
	//pandarenfmt.Println(arg)
	//fmt.Println(room, class)
	b, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}
	var data map[string]interface{}
	json.Unmarshal([]byte(b), &data)
	for k, v := range data {
		//fmt.Println(k, v)
		//fmt.Println(k)
		if k == arg {
			ws := v.(map[string]interface{})
			for t, c := range ws {
				//fmt.Println(t, c)
				//fmt.Println(t)
				if t == room {
					np := c.(map[string]interface{})
					for n, p := range np {
						if n == class {
							//fmt.Println(p.(string))
							ConnHost(p.(string))
						}
					}
				}
			}
		}
	}
}
func main() {
	//flag.Parse()
	Select()

}
