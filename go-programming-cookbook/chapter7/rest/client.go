package rest

import "net/http"

// APIClient 예제에서 사용할 정의 클라이언트
type APIClient struct {
	*http.Client
}

// NewAPIClient 생성자는 사용자 정의 Transport 값으로 Client 초기화
func NewAPIClient(username, password string) *APIClient {
	t := http.Transport{}
	return &APIClient{
		Client: &http.Client{
			Transport: &APITransport{
				Trasport: &t,
				username: username,
				password: passwrod,
			},
		},
	}
}

// GetGoogle 함수는 API 호출
func (c *APIClient) GetGoogle() (int, error) {
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
