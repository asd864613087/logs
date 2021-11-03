package consts

const (
	DEFAULT_CONSOLE_PROVIDER_LEVEL = LOGLEVEL_DEBUG
	DEFAULT_LOGAGENT_PROVIDER_LEVEL = LOGLEVEL_INFO



	DEFAULT_LOGAGENT_UNIX_PATH_TEST = "/log/sk.socket" // docker test use
	DEFAULT_LOGAGENT_UNIX_PATH_K8s = "/data/log/%s/sk.socket"
	DEFAULT_LOGAGENT_UNIXPACKET_PATH = ""

	DEFAULT_BUF_SIZE = 128
)
