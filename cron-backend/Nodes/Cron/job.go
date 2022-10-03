package cronjob

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type JobID int

type Job struct {
	ID             JobID  `json:"id"`
	WebhookURL     string `json:"webhook_url"`
	WebhookMethod  string `json:"webhook_method"`
	Body           string `json:"body"`
	CronExpression string `json:"expression"`
	Timezone       string `json:"timezone"`
}

func (j *Job) Trigger() {
	// fmt.Printf("%s - %d - %v\n", time.Now().Format(time.RubyDate), j.ID, j)
	client := http.DefaultClient
	var (
		url        string
		bodyReader io.Reader
	)

	if j.WebhookMethod == "GET" {
		// Insert body param in URL in case of GET method
		url = fmt.Sprintf(`%s?body="%s"`, j.WebhookURL, j.Body)
		bodyReader = nil
	} else {
		// Form body in every other method
		url = j.WebhookURL
		jsonBody := fmt.Sprintf(`{"body":"%s"}`, j.Body) // JSON encoding "à pedreiro"
		bodyReader = bytes.NewReader([]byte(jsonBody))
	}

	req, err := http.NewRequest(j.WebhookMethod, url, bodyReader)
	if err != nil {
		fmt.Println("error 'Trigger': Creating request: ", err.Error(), *j)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		fmt.Println("error 'Trigger': Sending request: ", err.Error(), *j)
		return
	}

	return
}
