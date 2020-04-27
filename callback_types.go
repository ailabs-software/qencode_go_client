package qencode

import (
  "strconv"
)

/** @fileoverview Structs for parsing callback. */

type QEncodeStatusVideoMeta struct {
  Width int `json: "width"`
  Height int `json: "height"`
}

type QEncodeStatusVideo struct {
  Url string `json: "url"`
  Duration string `json: "duration"`
  Meta QEncodeStatusVideoMeta `json: "meta"`
}

func (sv QEncodeStatusVideo) GetDurationFloat() (float64, error) {
  return strconv.ParseFloat(sv.Duration, 64)
}

func (sv QEncodeStatusVideo) GetDurationInt() (int, error) {
  f, err := sv.GetDurationFloat()
  if (err != nil) {
  	return 0, err
  }
  return int(f), nil
}

type QEncodeStatus struct {
  Status string `json: "status"`
  Videos []QEncodeStatusVideo `json: "videos"`
}
