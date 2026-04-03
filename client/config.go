package client

import kc "github.com/zaibyte/keeper/client"

// Config is the configs of Client.
// I think it makes things easier to share the same config among all kinds of Clients.
type Config struct {
	KeeperConfig kc.FCConfig `toml:"keeper_config"`
	// ZBufConns is the number of concurrent connections the client should establish to the zBuf sever.
	ZBufConns int `toml:"z_buf_conns"`
}
