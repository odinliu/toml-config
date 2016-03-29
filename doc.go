package tmlcfg

// use `github.com/pelletier/go-toml` to bind toml with golang struct
// Example:
//
// ```
// type GeneralConfig struct {
// 	LogLevel         int    `toml:"general.log_level"`
// 	LogFile          string `toml:"general.log_file"`
//	UserAgent        string `toml:"general.user_agent"`
//	EnableCoolie     bool   `toml:"general.enable_cookie"`
//	CrawlerCount     int    `toml:"general.crawler_count"`
//	HbaseWriterCount uint64 `toml:"general.hbase_writer_count"`
//	MaxChannelCount  int    `toml:"general.max_channel_count"`
// }
//
// type HbaseConfig struct {
//	ZkQuorum []string `toml:"hbase.zkQuorum"`
// 	ZkRoot   string   `toml:"hbase.zkRoot"`
// 	Table    string   `toml:"hbase.table"`
// }
//
// type SpiderConfig struct {
//	General GeneralConfig
//	Hbase   HbaseConfig
// }

// func main() {
//	config := SpiderConfig{}
//	tmlcfg.BindFile("spider.toml", &config)
//	fmt.Println(config.General.UserAgent)
//	fmt.Println(config.General.MaxChannelCount)
//	fmt.Println(config.General.EnableCoolie)
//	fmt.Println(config.Hbase.Table)
//	fmt.Println(config.Hbase.ZkQuorum)
//	fmt.Println(config.General.HbaseWriterCount)
//}
// ```
