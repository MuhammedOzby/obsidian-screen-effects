package youtube

// Client YouTube API ile iletişim kurar
type Client struct {
	APIKey    string
	ChannelID string
	// Gerçek uygulamada buraya bir http.Client ve OAuth token bilgileri de eklenir.
}

// NewClient yeni bir YouTube istemcisi oluşturur
func NewClient(apiKey, channelID string) *Client {
	return &Client{
		APIKey:    apiKey,
		ChannelID: channelID,
	}
}

// SubscriberInfo son abone bilgilerini temsil eder
type SubscriberInfo struct {
	Name          string `json:"name"`
	ProfilePicURL string `json:"profile_pic_url"`
}

// GetLatestSubscriber son abonenin bilgilerini döndürür.
// DİKKAT: Bu fonksiyon şu an için sahte veri döndürmektedir.
// Gerçek veri için OAuth 2.0 entegrasyonu gereklidir.
func (c *Client) GetLatestSubscriber() (*SubscriberInfo, error) {
	// Burası gelecekte gerçek API çağrısıyla doldurulacak.
	// Şimdilik sabit bir JSON dönüyoruz.
	mockSubscriber := &SubscriberInfo{
		Name:          "Kral Abone",
		ProfilePicURL: "www.youtube.com",
	}

	return mockSubscriber, nil
}
