package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	xr "metagrpc"
	"time"
)

var rpl = `
  {"Cisco-IOS-XR-policy-repository-cfg:routing-policy": {
   "route-policies": {
    "route-policy": [ null ]
   }
  }}`

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
	output, err := xr.GetConfig(ctx, conn, rpl, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\nrpl store via GRPC cfg from %s\n %s\n", router.Host, output)

	rpls, _ := UnmarshalRPLExample([]byte(output))
	for _, v := range rpls.CiscoIOSXRPolicyRepositoryCFGRoutingPolicy.RoutePolicies.RoutePolicy {
		fmt.Println(v.RoutePolicyName)
		fmt.Println("-------------")
		fmt.Println(v.RplRoutePolicy)
	}

}

func UnmarshalRPLExample(data []byte) (RPLExample, error) {
	var r RPLExample
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RPLExample) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RPLExample struct {
	CiscoIOSXRPolicyRepositoryCFGRoutingPolicy CiscoIOSXRPolicyRepositoryCFGRoutingPolicy `json:"Cisco-IOS-XR-policy-repository-cfg:routing-policy"`
}

type CiscoIOSXRPolicyRepositoryCFGRoutingPolicy struct {
	RoutePolicies RoutePolicies `json:"route-policies"`
}

type RoutePolicies struct {
	RoutePolicy []RoutePolicy `json:"route-policy"`
}

type RoutePolicy struct {
	RoutePolicyName string `json:"route-policy-name"`
	RplRoutePolicy  string `json:"rpl-route-policy"`
}
