package emq

import "testing"

func TestNewClient(t *testing.T) {
	NewClient()
}

func TestClient(t *testing.T) {
	NewClient()
	Publish(MqttClient)
	//Sub(MqttClient)
	//MqttClient.Disconnect(250)
}
