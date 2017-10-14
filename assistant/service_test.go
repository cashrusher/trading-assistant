package assistant

import "testing"

func TestServiceImpl_GetHistory(t *testing.T) {
	service := InitService()
	service.GetHistory()
}
