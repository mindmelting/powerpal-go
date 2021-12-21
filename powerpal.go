package powerpalgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mindmelting/powerpalgo/internal/clientutils"
)

const BaseUrl string = "https://readings.powerpal.net"

type PowerpalDevice struct {
	SerialNumber           string  `json:"serial_number"`
	TotalMeterReadingCount int     `json:"total_meter_reading_count"`
	TotalWattHours         int     `json:"total_watt_hours"`
	TotalCost              float64 `json:"total_cost"`
	FirstReadingTimestamp  int     `json:"first_reading_timestamp"`
	LastReadingTimestamp   int     `json:"last_reading_timestamp"`
	LastReadingWattHours   int     `json:"last_reading_watt_hours"`
	LastReadingCost        float64 `json:"last_reading_cost"`
	AvailableDays          int     `json:"available_days"`
}

type Powerpal struct {
	authKey, deviceId string
}

func New(authKey string, deviceId string) *Powerpal {
	return &Powerpal{
		authKey,
		deviceId,
	}
}

func (p Powerpal) getData() (*PowerpalDevice, error) {
	url := BaseUrl + "/api/v1/device/" + p.deviceId

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", p.authKey)

	resp, err := clientutils.Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return nil, &PowerpalAuthenticationError{}
		}
		if resp.StatusCode == http.StatusForbidden {
			return nil, &PowerpalAuthorizationError{}
		}
		return nil, &PowerpalRequestError{resp.StatusCode, string(body)}
	}

	var result *PowerpalDevice

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, fmt.Errorf("error deserialising response: %w", err)
	}

	return result, nil
}
