package main

import (
	"github.com/derzombiiie/emessage"

	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	username := flag.String("username", emessage.DefaultUsername, "username used to authenticate with api")
	password := flag.String("passwordhash", emessage.DefaultPasswordHash, "passwordhash used to authenticate with api")

	verbose := flag.Bool("verbose", false, "verbose logging")

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatalf("Usage: emessage [flags] [--] <identifier> <message>")
	}

	login2 := &emessage.Login2Request{
		Username:     *username,
		PasswordHash: *password,
	}

	l2res, err := login2.Send()
	if err != nil {
		log.Fatalf("Error during login2: %s", err)
	}

	if *verbose {
		fmt.Printf("Got JWT: %s", l2res.JWT)
	}

	sendRuf := &emessage.SendRufRequest{
		JWT: l2res.JWT,

		Identifier:  args[0],
		MessageText: strings.Join(args[1:], " "),
	}

	srres, err := sendRuf.Send()
	if err != nil {
		log.Fatalf("Error during sendRuf: %s", err)
	}

	fmt.Printf("%s\ntracking: %s\n", srres.Status, srres.TrackingID)

	if *verbose {
		fmt.Printf("Recipients %v", srres.Recipients)
	}
}
