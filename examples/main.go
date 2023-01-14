package main

import (
	"fmt"
	"log"
	"math/rand"
	xr "metagrpc"
	"metagrpc/pkg/oper/interfaces"
	"time"
)

var p = `{
    "Cisco-IOS-XR-ifmgr-oper:interface-properties": [null]
}`

var intbrief = `{
    "Cisco-IOS-XR-pfi-im-cmd-oper:interfaces":
    {
        "interface-briefs": {"interface-brief": {"interface-name":[null], "type":[null]}}
     
    }
}`

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := r.Int63n(10000)
	router, err := xr.BuildRouter(
		xr.WithUsername("grpc"),
		xr.WithPassword("53cret"),
		xr.WithHost("192.168.88.3:57344"),
		xr.WithTimeout(600),
	)
	conn, ctx, err := xr.Connect(*router)
	if err != nil {
		log.Fatalf("could not setup a client connection to %s, %v", router.Host, err)
	}
	defer conn.Close()
	output, err := xr.Get(ctx, conn, intbrief, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\ninterface properties from %s\n %s\n", router.Host, output)

	// New method of getting shit
	io := interfaces.Interfaces{}
	err = interfaces.Unmarshal([]byte(output), &io)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range io.Interfaces.InterfaceBriefs.InterfaceBrief {
		fmt.Println(k, *v.InterfaceName)
	}

}
