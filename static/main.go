package main //provê as funções para manter a aplicação de pé

import (
	"net/http" //fornece os métodos de mapeamento e gerenciamento de requisições
)

func handler(w http.ResponseWriter, r *http.Request) { //o primeiro parâmetro será responsável pela resposta da requisiçõa e o segundo para o tratamento da requisição
	//essa função poderia ter outro nome

	http.ServeFile(w, r, "./static/index.html") //o ServeFile é responsável por dispor um arquivo. estamos utilizando nossas variáveis

}

func main() { //função que será o ponto de partida

	http.HandleFunc("/", handler) //o HandleFunc é usado para lidar com as requisições que o servidor receberá
	//o HandleFunc recebe dois parâmetros. O primeiro informa que a URL  http://localhost:8080  deve executar determinada ação
	//o segundo parâmetro indica que chamaremos a função handler ao acessarmos o endereço  http://localhost:8080
	http.ListenAndServe(":8080", nil) //o ListenAndServe será responsável por especificar em qual porta rodará a nossa aplicação e como lidaremos com
	//com o servidor da aplicação
	//no caso, está sendo usada a porta 8080 e o nil é para informar que será usada a configuração padrão do servidor Go
}
