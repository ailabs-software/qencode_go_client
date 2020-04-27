package qencode

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

type QEncodeStatus struct {
  Status string `json: "status"`
  Videos []QEncodeStatusVideo `json: "videos"`
}
