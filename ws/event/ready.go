package event

import "encoding/json"

type Ready struct {
    Data struct {
        Version int `json:"v"`
        // User *User `json:"user"`
        // Guilds []*Guild `json:"guilds"`
        SessionID string `json:"session_id"`
        Shard []int `json:"shard,omitempty"`
    } `json:"d"`
}

func NewReady(data []byte) (*Ready, error) {
    pk := new(Ready)
    
    err := json.Unmarshal(data, pk)
    
    if err != nil {
        return nil, err
    }
    
    return pk, nil
}