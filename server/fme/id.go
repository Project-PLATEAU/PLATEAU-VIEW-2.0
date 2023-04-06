package fme

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

type ID struct {
	ItemID    string
	AssetID   string
	ProjectID string
}

var ErrInvalidID = errors.New("invalid id")

func ParseID(id, secret string) (ID, error) {
	sig, payload, found := strings.Cut(id, ":")
	if !found {
		return ID{}, ErrInvalidID
	}

	s := strings.SplitN(payload, ";", 4)
	if len(s) != 3 {
		return ID{}, ErrInvalidID
	}

	if sig != sign(payload, secret) {
		return ID{}, ErrInvalidID
	}

	return ID{
		ItemID:    s[0],
		AssetID:   s[1],
		ProjectID: s[2],
	}, nil
}

func (i ID) String(secret string) string {
	payload := fmt.Sprintf("%s;%s;%s", i.ItemID, i.AssetID, i.ProjectID)
	sig := sign(payload, secret)
	return fmt.Sprintf("%s:%s", sig, payload)
}

func sign(payload, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}
