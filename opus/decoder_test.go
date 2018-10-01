package opus

import (
	"testing"
)

func TestDecoderNew(t *testing.T) {
	dec, err := NewDecoder(48000, 1)
	if err != nil || dec == nil {
		t.Errorf("Error creating new decoder: %v", err)
	}
	dec, err = NewDecoder(12345, 1)
	if err == nil || dec != nil {
		t.Errorf("Expected error for illegal samplerate 12345")
	}
}

func TestDecoderUnitialized(t *testing.T) {
	var dec Decoder
	_, err := dec.Decode(nil, nil)
	if err != errDecUninitialized {
		t.Errorf("Expected \"unitialized decoder\" error: %v", err)
	}
	_, err = dec.DecodeFloat32(nil, nil)
	if err != errDecUninitialized {
		t.Errorf("Expected \"unitialized decoder\" error: %v", err)
	}
	err = dec.DecodeFEC(nil, nil)
	if err != errDecUninitialized {
		t.Errorf("Expected \"unitialized decoder\" error: %v", err)
	}
	err = dec.DecodeFECFloat32(nil, nil)
	if err != errDecUninitialized {
		t.Errorf("Expected \"unitialized decoder\" error: %v", err)
	}
}
func TestDecoder_GetLastPacketDuration(t *testing.T) {
	const G4 = 391.995
	const sampleRate = 48000
	const frameSizeMs = 60
	const frameSize = sampleRate * frameSizeMs / 1000
	pcm := make([]int16, frameSize)
	enc, err := NewEncoder(sampleRate, 1, AppVoIP)
	if err != nil || enc == nil {
		t.Fatalf("Error creating new encoder: %v", err)
	}
	addSine(pcm, sampleRate, G4)

	data := make([]byte, 1000)
	n, err := enc.Encode(pcm, data)
	if err != nil {
		t.Fatalf("Couldn't encode data: %v", err)
	}
	data = data[:n]

	dec, err := NewDecoder(sampleRate, 1)
	if err != nil || dec == nil {
		t.Fatalf("Error creating new decoder: %v", err)
	}
	n, err = dec.Decode(data, pcm)
	if err != nil {
		t.Fatalf("Couldn't decode data: %v", err)
	}
	samples, err := dec.LastPacketDuration()
	if err != nil {
		t.Fatalf("Couldn't get last packet duration: %v", err)
	}
	if samples != n {
		t.Fatalf("Wrong duration length. Expected %d. Got %d", n, samples)
	}
}
