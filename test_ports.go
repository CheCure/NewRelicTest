package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"regexp"
	"encoding/json"
)

type openports struct {
	ports []string;
}


func main() {
	var plen int;
	var input string;
	var portinfo string;
	var lastport string;
	var port string;
	var ports []string;	
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)


	cmd := exec.Command("sh","./ports")
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
/*	fmt.Printf("Test: %q\n", out.String()) */
	temp1 := strings.Split(out.String(),"\n")


	plen = len(temp1)
 	plen=plen-1	
	for i := 0; i < plen; i++ {
		input = temp1[i]
			
		final := re_leadclose_whtsp.ReplaceAllString(input, "")
		final = re_inside_whtsp.ReplaceAllString(final, " ")
		temp2 := strings.Split(final," ")
		portinfo=temp2[3];

		if strings.Contains(portinfo,"::") || strings.Contains(portinfo,"127.0.0.1")  {
		} else if portinfo == lastport {
					} else {
						lastport=portinfo;
						temp3 := strings.Split(portinfo,":");
						port=temp3[1];					
						fmt.Printf("Open Port: %s\n",port);
						ports = append(ports, port)
					}		

	}

/*	var ps openports;
	ps.ports = ports;
*/
	portsJson, err := json.Marshal(ports);
	if err!= nil {
		log.Fatal("Cannot code JSON:", err)
	}	

	fmt.Printf("Ports... %s\n\n",portsJson);



}
