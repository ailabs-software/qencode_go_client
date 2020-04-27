package qencode

import (
  "errors"
  "strconv"
  "net/url"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "log"
)

/** @fileoverview Go wrapper for Qencode's REST API. */

const API_BASE_URL = "https://api.qencode.com/v1/"

const EMPTY_STRING = ""

func sendRequest(endpoint string, form url.Values) ([]byte, error) {
  httpClient := &http.Client{}
  res, err := httpClient.PostForm(API_BASE_URL + endpoint, form)
  if (err != nil) {
    return nil, err
  }

  bodyBytes, err := ioutil.ReadAll(res.Body)
  if (err != nil) {
    return nil, err
  }

  if (res.StatusCode != 200) {
    errorMessage := "Qencode endpoint " + endpoint + " returned status code: " + strconv.Itoa(res.StatusCode)
    log.Println(errorMessage)
    log.Println( string(bodyBytes) )
    return nil, errors.New(errorMessage)
  }

  return bodyBytes, nil
}

func sendRequestAndParseResponse(data interface{}, endpoint string, form url.Values) error {
  body, err := sendRequest(endpoint, form)
  if (err != nil) {
    return err
  }
  return json.Unmarshal(body, &data)
}

type QEncodeResponse struct {
  Error int `json:"error"`
  Message string `json:"message"`
  Token string `json:"token"`
  TaskToken string `json:"task_token"`
}

func (r QEncodeResponse) IsNotSuccessful() bool {
  return r.Error != 0
}

func (r QEncodeResponse) GetError() error {
  return errors.New(r.Message)
}

func GetAccessToken(apiKey string) (string, error) {
  response := QEncodeResponse{}
  err := sendRequestAndParseResponse(&response, "access_token", url.Values{"api_key": []string{apiKey} })
  if (err != nil) {
  	return EMPTY_STRING, err
  }

  if ( response.IsNotSuccessful() ) {
  	return EMPTY_STRING, response.GetError()
  }

  return response.Token, nil
}

func CreateTask(accessToken string) (string, error) {
  response := QEncodeResponse{}
  err := sendRequestAndParseResponse(&response, "create_task", url.Values{"token": []string{accessToken} })
  if (err != nil) {
    return EMPTY_STRING, err
  }

  if ( response.IsNotSuccessful() ) {
    return EMPTY_STRING, response.GetError()
  }

  return response.TaskToken, nil
}

/** @param payload -- Can be used as an identifier */
func StartEncode(taskToken string, profile string, sourceUrl string, payload string) error {

  form := url.Values{
    "task_token": []string{taskToken},
     "profiles": []string{profile},
     "uri": []string{sourceUrl},
     "payload": []string{payload} }

  response := QEncodeResponse{}

  err := sendRequestAndParseResponse(&response, "start_encode", form)
  if (err != nil) {
    return err
  }

  if ( response.IsNotSuccessful() ) {
    return response.GetError()
  }

  return nil
}
