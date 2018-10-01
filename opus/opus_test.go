package opus

import (
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	if ver := Version(); !strings.HasPrefix(ver, "libopus") {
		t.Errorf("Unexpected linked libopus version: " + ver)
	}
}

func TestOpusErrstr(t *testing.T) {
	if ErrOK.Error() != "opus: success" {
		t.Errorf("Expected \"success\" error message for error code %d: %v", ErrOK, ErrOK)
	}
	if ErrBadArg.Error() != "opus: invalid argument" {
		t.Errorf("Expected \"invalid argument\" error message for error code %d: %v", ErrBadArg, ErrBadArg)
	}
	if ErrBufferTooSmall.Error() != "opus: buffer too small" {
		t.Errorf("Expected \"buffer too small\" error message for error code %d: %v", ErrBufferTooSmall, ErrBufferTooSmall)
	}
	if ErrInternalError.Error() != "opus: internal error" {
		t.Errorf("Expected \"internal error\" error message for error code %d: %v", ErrInternalError, ErrInternalError)
	}
	if ErrInvalidPacket.Error() != "opus: corrupted stream" {
		t.Errorf("Expected \"corrupted stream\" error message for error code %d: %v", ErrInvalidPacket, ErrInvalidPacket)
	}
	if ErrUnimplemented.Error() != "opus: request not implemented" {
		t.Errorf("Expected \"request not implemented\" error message for error code %d: %v", ErrUnimplemented, ErrUnimplemented)
	}
	if ErrInvalidState.Error() != "opus: invalid state" {
		t.Errorf("Expected \"invalid state\" error message for error code %d: %v", ErrInvalidState, ErrInvalidState)
	}
	if ErrAllocFail.Error() != "opus: memory allocation failed" {
		t.Errorf("Expected \"memory allocation failed\" error message for error code %d: %v", ErrAllocFail, ErrAllocFail)
	}
}

func TestCodec(t *testing.T) {
	// Create bogus input sound
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
	if len(pcm) != n {
		t.Fatalf("Length mismatch: %d samples in, %d out", len(pcm), n)
	}
}

func TestCodecFloat32(t *testing.T) {
	// Create bogus input sound
	const G4 = 391.995
	const sampleRate = 48000
	const frameSizeMs = 60
	const frameSize = sampleRate * frameSizeMs / 1000
	pcm := make([]float32, frameSize)
	enc, err := NewEncoder(sampleRate, 1, AppVoIP)
	if err != nil || enc == nil {
		t.Fatalf("Error creating new encoder: %v", err)
	}
	addSineFloat32(pcm, sampleRate, G4)
	data := make([]byte, 1000)
	n, err := enc.EncodeFloat32(pcm, data)
	if err != nil {
		t.Fatalf("Couldn't encode data: %v", err)
	}
	data = data[:n]
	dec, err := NewDecoder(sampleRate, 1)
	if err != nil || dec == nil {
		t.Fatalf("Error creating new decoder: %v", err)
	}
	n, err = dec.DecodeFloat32(data, pcm)
	if err != nil {
		t.Fatalf("Couldn't decode data: %v", err)
	}
	if len(pcm) != n {
		t.Fatalf("Length mismatch: %d samples in, %d out", len(pcm), n)
	}
}

func TestCodecFEC(t *testing.T) {
	// Create bogus input sound
	const G4 = 391.995
	const sampleRate = 48000
	const frameSizeMs = 10
	const frameSize = sampleRate * frameSizeMs / 1000
	const numberOfFrames = 6
	pcm := make([]int16, 0)

	enc, err := NewEncoder(sampleRate, 1, AppVoIP)
	if err != nil || enc == nil {
		t.Fatalf("Error creating new encoder: %v", err)
	}
	enc.SetPacketLossPerc(30)
	enc.SetInBandFEC(true)

	dec, err := NewDecoder(sampleRate, 1)
	if err != nil || dec == nil {
		t.Fatalf("Error creating new decoder: %v", err)
	}

	mono := make([]int16, frameSize*numberOfFrames)
	addSine(mono, sampleRate, G4)

	encodedData := make([][]byte, numberOfFrames)
	for i, idx := 0, 0; i < len(mono); i += frameSize {
		data := make([]byte, 1000)
		n, err := enc.Encode(mono[i:i+frameSize], data)
		if err != nil {
			t.Fatalf("Couldn't encode data: %v", err)
		}
		data = data[:n]
		encodedData[idx] = data
		idx++
	}

	lost := false
	for i := 0; i < len(encodedData); i++ {
		// "Dropping" packets 2 and 4
		if i == 2 || i == 4 {
			lost = true
			continue
		}
		if lost {
			samples, err := dec.LastPacketDuration()
			if err != nil {
				t.Fatalf("Couldn't get last packet duration: %v", err)
			}

			pcmBuffer := make([]int16, samples)
			if err = dec.DecodeFEC(encodedData[i], pcmBuffer); err != nil {
				t.Fatalf("Couldn't decode data: %v", err)
			}
			pcm = append(pcm, pcmBuffer...)
			lost = false
		}

		pcmBuffer := make([]int16, numberOfFrames*frameSize)
		n, err := dec.Decode(encodedData[i], pcmBuffer)
		if err != nil {
			t.Fatalf("Couldn't decode data: %v", err)
		}
		pcmBuffer = pcmBuffer[:n]
		if n != frameSize {
			t.Fatalf("Length mismatch: %d samples expected, %d out", frameSize, n)
		}
		pcm = append(pcm, pcmBuffer...)
	}

	if len(mono) != len(pcm) {
		t.Fatalf("Input/Output length mismatch: %d samples in, %d out", len(mono), len(pcm))
	}
}

func TestCodecFECFloat32(t *testing.T) {
	// Create bogus input sound
	const G4 = 391.995
	const sampleRate = 48000
	const frameSizeMs = 10
	const frameSize = sampleRate * frameSizeMs / 1000
	const numberOfFrames = 6
	pcm := make([]float32, 0)

	enc, err := NewEncoder(sampleRate, 1, AppVoIP)
	if err != nil || enc == nil {
		t.Fatalf("Error creating new encoder: %v", err)
	}
	enc.SetPacketLossPerc(30)
	enc.SetInBandFEC(true)

	dec, err := NewDecoder(sampleRate, 1)
	if err != nil || dec == nil {
		t.Fatalf("Error creating new decoder: %v", err)
	}

	mono := make([]float32, frameSize*numberOfFrames)
	addSineFloat32(mono, sampleRate, G4)

	encodedData := make([][]byte, numberOfFrames)
	for i, idx := 0, 0; i < len(mono); i += frameSize {
		data := make([]byte, 1000)
		n, err := enc.EncodeFloat32(mono[i:i+frameSize], data)
		if err != nil {
			t.Fatalf("Couldn't encode data: %v", err)
		}
		data = data[:n]
		encodedData[idx] = data
		idx++
	}

	lost := false
	for i := 0; i < len(encodedData); i++ {
		// "Dropping" packets 2 and 4
		if i == 2 || i == 4 {
			lost = true
			continue
		}
		if lost {
			samples, err := dec.LastPacketDuration()
			if err != nil {
				t.Fatalf("Couldn't get last packet duration: %v", err)
			}

			pcmBuffer := make([]float32, samples)
			if err = dec.DecodeFECFloat32(encodedData[i], pcmBuffer); err != nil {
				t.Fatalf("Couldn't decode data: %v", err)
			}
			pcm = append(pcm, pcmBuffer...)
			lost = false
		}

		pcmBuffer := make([]float32, numberOfFrames*frameSize)
		n, err := dec.DecodeFloat32(encodedData[i], pcmBuffer)
		if err != nil {
			t.Fatalf("Couldn't decode data: %v", err)
		}
		pcmBuffer = pcmBuffer[:n]
		if n != frameSize {
			t.Fatalf("Length mismatch: %d samples expected, %d out", frameSize, n)
		}
		pcm = append(pcm, pcmBuffer...)
	}

	if len(mono) != len(pcm) {
		t.Fatalf("Input/Output length mismatch: %d samples in, %d out", len(mono), len(pcm))
	}
}

func TestStereo(t *testing.T) {
	// Create bogus input sound
	const G4 = 391.995
	const E3 = 164.814
	const sampleRate = 48000
	const frameSizeMs = 60
	const CHANNELS = 2
	const frameSizeMono = sampleRate * frameSizeMs / 1000

	enc, err := NewEncoder(sampleRate, CHANNELS, AppVoIP)
	if err != nil || enc == nil {
		t.Fatalf("Error creating new encoder: %v", err)
	}
	dec, err := NewDecoder(sampleRate, CHANNELS)
	if err != nil || dec == nil {
		t.Fatalf("Error creating new decoder: %v", err)
	}

	// Source signal (left G4, right E3)
	left := make([]int16, frameSizeMono)
	right := make([]int16, frameSizeMono)
	addSine(left, sampleRate, G4)
	addSine(right, sampleRate, E3)
	pcm := interleave(left, right)

	data := make([]byte, 1000)
	n, err := enc.Encode(pcm, data)
	if err != nil {
		t.Fatalf("Couldn't encode data: %v", err)
	}
	data = data[:n]
	n, err = dec.Decode(data, pcm)
	if err != nil {
		t.Fatalf("Couldn't decode data: %v", err)
	}
	if n*CHANNELS != len(pcm) {
		t.Fatalf("Length mismatch: %d samples in, %d out", len(pcm), n*CHANNELS)
	}
}
