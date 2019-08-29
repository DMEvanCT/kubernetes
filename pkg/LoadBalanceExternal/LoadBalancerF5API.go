package LoadBalanceExternal

import (
	"bytes"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"net/http"
)


func CreateF5Pool(vip *F5VIP, f5BaseURL string) {
	data, _ := json.Marshal(&vip)
	f5url := f5BaseURL + "mgmt/" + "tm/" + "/ltm" + "/pool"
	req, err := http.NewRequest("POST", f5url, bytes.NewBuffer(data))

	// Set the Content type to json
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal("There was an error ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()


}