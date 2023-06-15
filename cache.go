package main

type Cacher interface {
	Get(key int) (string, bool)
	Set(key int, value string) error
	Remove(key int) error
}

type NopCache struct{}

func (n *NopCache) Get(key int) (string, bool) {
	return "", false
}
func (n *NopCache) Set(key int, value string) error {
	return nil
}
func (n *NopCache) Remove(key int) error {
	return nil
}
