# Mercado Bitcoin Api de Dados Go Client

**Um simples cliente escrito em Go para a API de Dados do site Mercado Bitcoin &lt;www.mercadobitcoin.com.br>**

## Exemplo de uso:

```
package main

import (
	"fmt"
	"github.com/awmpietro/mcdbcgoclient"
)

func main() {
	ticker, _ := mcdbitcoindd.GetTicker()
	orderBook, _ := mcdbitcoindd.GetOrderBook()
	trades, _ := mcdbitcoindd.GetTrades()
	fmt.Println(ticker)
	fmt.Println(orderBook)
	fmt.Println(trades)

}
```
### TODO:
Mais opções na função GetTrades()

