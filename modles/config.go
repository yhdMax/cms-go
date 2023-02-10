package modles

type Response struct {
	Data   interface{} `json:"result"`
	Code   int         `json:"code"`
	Status int         `json:"status"`
	Error  string      `json:"msg"`
}

type Token struct {
	Token string `db:"-" json:"token"`
}

type List struct {
	List interface{} `db:"-" json:"list"`
}

type Configuration struct {
	DB struct {
		URL          string `json:"url"`
		MaxIdleConns int    `json:"max_idle_conns"`
		MaxOpenConns int    `json:"max_open_conns"`
	}
}
