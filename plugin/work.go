package main

import "sdcore"

func SDAuth() func([]byte) error {
	return sdcore.SDAuth()
}

func SDEncrypt(data []byte) []byte {
	return sdcore.SDEncrypt(data)
}

func SDDecrypt(data []byte) []byte {
	return sdcore.SDDecrypt(data)
}

func SDPush(selfID string, recvIds []string, data []byte, isFile bool) (bool, error) {
	return sdcore.SDPush(selfID, recvIds, data, isFile)
}

func SDPull(selfID string) (bool, []string, error) {
	return sdcore.SDPull(selfID)
}