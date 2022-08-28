package region

import (
	"encoding/base64"
	"flswld.com/gate-genshin-api/proto"
	"fmt"
	pb "google.golang.org/protobuf/proto"
	"testing"
)

func TestRegion(t *testing.T) {
	var err error = nil
	regionListStr := "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"
	regionListBin, err := base64.StdEncoding.DecodeString(regionListStr)
	if err != nil {
		panic(err)
	}
	regionList := new(proto.QueryRegionListHttpRsp)
	err = pb.Unmarshal(regionListBin, regionList)
	if err != nil {
		panic(err)
	}
	fmt.Println(regionList)
	regionCurrStr := "GpkdCg00Ny4yNTMuNDMuMTA2ENasARoraHR0cHM6Ly9vc3VzYW9hc2VydmVyLnl1YW5zaGVuLmNvbS9yZWNoYXJnZToDVVNBQjlodHRwczovL2F1dG9wYXRjaGhrLnl1YW5zaGVuLmNvbS9jbGllbnRfZ2FtZV9yZXMvMi42X2xpdmVKPGh0dHBzOi8vYXV0b3BhdGNoaGsueXVhbnNoZW4uY29tL2NsaWVudF9kZXNpZ25fZGF0YS8yLjZfbGl2ZVKSAWh0dHBzOi8vd2Vic3RhdGljLXNlYS5ob3lvdmVyc2UuY29tL3lzL2V2ZW50L2ltLXNlcnZpY2UvaW5kZXguaHRtbD9pbV9vdXQ9ZmFsc2Umc2lnbl90eXBlPTImYXV0aF9hcHBpZD1pbV9jY3MmYXV0aGtleV92ZXI9MSZ3aW5fZGlyZWN0aW9uPXBvcnRyYWl0YggyLjZfbGl2ZXDZ7JoDkAGJxKoDmgFceyJyZW1vdGVOYW1lIjogImRhdGFfdmVyc2lvbnMiLCAibWQ1IjogIjFhNDQ1YWU0OGZiZGJlMDM4MzEyYmU4NjhmMjgxM2QyIiwgImZpbGVTaXplIjogNDQxNX2iAVt7InJlbW90ZU5hbWUiOiAiZGF0YV92ZXJzaW9ucyIsICJtZDUiOiAiZjY0NmZlY2FjMTA3ZDNlN2ZiNDdjOTA4OTIwMmQzZmYiLCAiZmlsZVNpemUiOiA1MTR9sgGBBgi9t5kDGuAFeyJyZW1vdGVOYW1lIjogInJlc192ZXJzaW9uc19leHRlcm5hbCIsICJtZDUiOiAiOWYwOTJjYTEzMDY3Y2FlNDMxMzQ5M2VlZDNkM2U5YWUiLCAiZmlsZVNpemUiOiA1MjY5NjZ9DQp7InJlbW90ZU5hbWUiOiAicmVzX3ZlcnNpb25zX21lZGl1bSIsICJtZDUiOiAiN2M5NDJlNzA0YTgwNGMyMTIyZmM2M2VlOTI4ZTcxNDgiLCAiZmlsZVNpemUiOiAyODU1NTB9DQp7InJlbW90ZU5hbWUiOiAicmVzX3ZlcnNpb25zX3N0cmVhbWluZyIsICJtZDUiOiAiZTM1ZjZkODUzZDA0Y2UwMjJhYzdjMzZkNDNmNDBlN2MiLCAiZmlsZVNpemUiOiAxMjIwMzZ9DQp7InJlbW90ZU5hbWUiOiAicmVsZWFzZV9y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"
	regionCurrBin, err := base64.StdEncoding.DecodeString(regionCurrStr)
	if err != nil {
		panic(err)
	}
	regionCurr := new(proto.QueryCurrRegionHttpRsp)
	err = pb.Unmarshal(regionCurrBin, regionCurr)
	if err != nil {
		panic(err)
	}
	fmt.Println(regionCurr)
}
