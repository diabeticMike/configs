package config

import "flag"

func NewFlag() Config {
	var conf Config
	listeningUrl := flag.String("listening-url", "localhost:8080", "Our api listening url")
	delay := flag.Uint("delay", 0, "Some user delay")
	flag.Parse()
	conf.ListeningUrl = *listeningUrl
	conf.Delay = *delay
	return conf
}
