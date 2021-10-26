package fetch

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func GetLokasi(url string)([]byte,error) {
    fmt.Println("Starting the application...")
    response, err := http.Get("https://dev.farizdotid.com/"+url)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } 
     data, _ := ioutil.ReadAll(response.Body)
    // fmt.Println(string(data))

	return data,nil	
}

func TestLokasi(url string)([]byte,error){
    client := &http.Client{}
     req, err := http.NewRequest("GET", "https://dev.farizdotid.com/"+url, nil)
     if err != nil {
      fmt.Print(err.Error())
     }
     req.Header.Add("Accept", "application/json")
     req.Header.Add("Content-Type", "application/json")
     resp, err := client.Do(req)
     if err != nil {
      fmt.Print(err.Error())
     }
    defer resp.Body.Close()
     bodyBytes, err := ioutil.ReadAll(resp.Body)
     if err != nil {
      fmt.Print(err.Error())
     }
    // fmt.Println(string(bodyBytes))
    // var responseObject Response
    //  json.Unmarshal(bodyBytes, &responseObject)
    //  fmt.Printf("API Response as struct %+v\n", responseObject)
    return bodyBytes, nil
}


func LocationGet(Url string) ([]byte, error) {

	url := "https://dev.farizdotid.com/"+Url
	method := "GET"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body))
	return body, nil
}
