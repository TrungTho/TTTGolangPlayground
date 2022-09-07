package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

const (
	sheetId   = "1NbMdJe95OTR8Wbt_KrpUuyP1LrJjZ1MzzPX3z7mqHT8"
	sheetName = "Sheet1"
	rowLimit  = 100
)

type accountAuthData struct {
	Type          string `json:"type"`
	ProjectId     string `json:"project_id"`
	PrivateKeyId  string `json:"private_key_id"`
	PrivateKey    string `json:"private_key"`
	ClientEmail   string `json:"client_email"`
	CliendId      string `json:"client_id"`
	AuthUri       string `json:"auth_uri"`
	TokenUri      string `json:"token_uri"`
	AuthProvider  string `json:"auth_provider_x509_cert_url"`
	ClientCertUrl string `json:"client_x509_cert_url"`
}

type LogData struct {
	ID      int
	Name    string
	Address string
}

func main() {
	authData := accountAuthData{
		Type:          "service_account",
		ProjectId:     "vernal-tracer-361607",
		PrivateKeyId:  "2c287edc3aa464843201b98684ead8d55cd382bd",
		PrivateKey:    "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCWJDTbRo66XuoP\ndI93oKBknzcONH22guNtou4uFkr8NdiJqkEUW/hW8jTjsKzMKHz95En7exzK2Ejq\nWlVUlpMtyEQC7MBQCdosNJO7U0kyANE0ghG6GEKTNHhsDYzY7NIyywAyebeooIR2\nqkBLQcNcCGG8NNvdxvx1I/3TEan8C3uJQ43QF9Eu0++ybXoXTFwEghxGtt2ji2Ww\nAfjsalS2Fe5ldBFg+Ub1gktEVgorHtsQA3CdEToNHkCGumHaQeO8NzE6o+sO7i4B\nHG6Ko/MvVcm53hEj2EFvrR4AWGXmAaZ4rQNTlNCgMCxqub9xUvUrhv6Jz/Q5Khok\nhrDZQFXJAgMBAAECggEAAo6WDFx2oo+qgI4hgd3tUV1hSoeEExNGVIgLcqM8TnCL\n0idhQZuh3ng0RKLuBHGVlFzakjL57e3ySvR0ItNIKSXRm7OkwFA0pB9Wm0B7PQPk\nzDZBX0gZvShN+zgdW5Xn7AfHRnTKdw+ZxewXGXkCntsBMBt92ZhxlPyijmFNAe74\n73RqAxsa107p2IhjaBz5ejQ5DtWkxKR7tm9IVDJUyGUkQCy5MmrsAxZq2PfNkeLH\n2iggvOmY0Ku+x1PzlqwpGCduSyg3jsFyoMKJF8DIRIbLHW37wsgLtXAhVNka5o9n\nvZwUqEKmUr+Gqf9Er/clC4n8DGAOCG5ucTxW2P2MYQKBgQDR/RSA1AL1Oz3WLK3T\n1izVX9xa0AXkqf3knBEvi4nO3no07kDpb8J6V9k40HWqcSX/TMQiaJIGkdZu69l4\ni71rmNwRbIe0voRuWpTWIcFxc7VLmbYQ0yir6FKi6Jw+TQ+pcCyGlA8/CAK9126Z\nsyjtEtNUwAlXkvS0sEWCVitD4QKBgQC3ChxsI04DDKwOAlrUJ9znQDnHh1TUjgmf\nt4lIVjQaN9gg84zd/zxDSAmXoBSpFHZS50p0LzO3a3DFoe+jdx2CW9TQe2SdE23u\nvMTtS2lPsG5uE4TDrGPlWfnikepB4RFXqFpZ+nvx5NHoIiy+MDJ1WMYbadH8jj+E\nb/m3I1FO6QKBgGoNomxKJ1BJYjqoCAaj9bKyHm0zALby78qk070qgSgcjqXq6pe7\nHQKDGa8rATJawPEGiUxDefSddSpCLWxHTxxncEXQhV1QlzvQvbjEBZnR8W9EK4Kl\n0rW3uPyT9E02yEEv6Rzy7BxOZGwSwMYZiQLq7hawgkdbbgFPwVtJP8KhAoGBALC6\n9IJ8/B/5pk7Ie0aJTsOBwcgjpQauNiCep9DOWvRNo0L9pa/bdyZHceuSxyAR/8VA\nSSUxRi/9bx+DocwlgLqTTEIYQidf0S9H2KR9wasN4TIram88DiAu5hWbaaI+W+5V\nQRfLwMzocLw/8w+XncCr/GwPmo7OEgofy+7GDQWxAoGAR58Yo6+N8ryqe03yMu0v\ncBX3ZZArS+otCctXO2vK8MODxxj3WFUec7gxlLm0FCZxKUquVkuk4YLb8Zeto12u\nCwgl8x7WT3nR24CDZ0GHAo1EYrKJMkru9v39NOlsFvqZYC8JjTmSe541GCemZXtA\nNHoVKuDpPmDNriVdHhmfyAE=\n-----END PRIVATE KEY-----\n",
		ClientEmail:   "test-ggsheet-io@vernal-tracer-361607.iam.gserviceaccount.com",
		CliendId:      "112623922872063864458",
		AuthUri:       "https://accounts.google.com/o/oauth2/auth",
		TokenUri:      "https://oauth2.googleapis.com/token",
		AuthProvider:  "https://www.googleapis.com/oauth2/v1/certs",
		ClientCertUrl: "https://www.googleapis.com/robot/v1/metadata/x509/test-ggsheet-io%40vernal-tracer-361607.iam.gserviceaccount.com",
	}

	// data, err := ioutil.ReadFile("credentials.json")
	data, err := json.Marshal(authData)
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	checkError(err)

	client := conf.Client(context.TODO())
	srv, err := sheets.New(client)
	checkError(err)

	spreadsheetID := sheetId
	readRange := sheetName
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	checkError(err)

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			fmt.Printf("%s, %s\n", row[0], row[1])
		}
	}

	if len(resp.Values) >= rowLimit {
		fmt.Println("row limit exceeded, can not write more")
		return
	}

	var vr sheets.ValueRange

	myval := []interface{}{99, "Ultimate User After Change auth type"}
	vr.Values = append(vr.Values, myval)
	writeRange := fmt.Sprintf("%s!A%d", sheetName, len(resp.Values)+1)

	_, err = srv.Spreadsheets.Values.Update(sheetId, writeRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
