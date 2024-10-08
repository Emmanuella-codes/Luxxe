package mail

import (
	"sync"
	"time"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
)

type SendPulseFetchKeyResponse struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type GlobalSendpulseKeyCacheStruct struct {
	Data      SendPulseFetchKeyResponse
	CreatedAt time.Time
	mu        sync.Mutex
}

var GlobalSendpulseKeyCache = &GlobalSendpulseKeyCacheStruct{
	Data:      SendPulseFetchKeyResponse{},
	CreatedAt: time.Now(),
}

func SetGlobalSendpulseCache(tokenType string, expiresIn int, accessToken string) {
	GlobalSendpulseKeyCache.mu.Lock()
	defer GlobalSendpulseKeyCache.mu.Unlock()

	GlobalSendpulseKeyCache.Data = SendPulseFetchKeyResponse{
		TokenType:   tokenType,
		ExpiresIn:   expiresIn,
		AccessToken: accessToken,
	}

	GlobalSendpulseKeyCache.CreatedAt = time.Now()
}

type GlobalEmailStatsCacheStruct struct {
	Data      entities.EmailStats
	CreatedAt time.Time
	mu        sync.Mutex
}

var GlobalEmailStatsCache = &GlobalEmailStatsCacheStruct{
	Data:      entities.EmailStats{},
	CreatedAt: time.Now(),
}

func SetGlobalEmailStatsCache(emailStats entities.EmailStats) {
	GlobalEmailStatsCache.mu.Lock()
	defer GlobalEmailStatsCache.mu.Unlock()

	GlobalEmailStatsCache.Data = emailStats

	GlobalEmailStatsCache.CreatedAt = time.Now()
}
