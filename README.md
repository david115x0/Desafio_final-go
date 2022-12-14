# Avaliação final go

## David da Silva Sousa

## packages to install

<pre><code> go get github.com/gin-gonic/gin </code></pre>
<pre><code> go get github.com/joho/godotenv </code></pre>
<pre><code> go get github.com/go-sql-driver/mysql </code></pre>
<pre><code> go get github.com/swaggo/swag/cmd/swag@latest </code></pre>
<pre><code> go get install github.com/swaggo/swag </code></pre>
<pre><code> go get github.com/swaggo/files </code></pre>

## antes de inciar o swagger
<pre><code>swag init -g cmd/server/main.go </code></pre>

# Postamn json

## Pacientes
<pre><code>{
    "nome": "Lionel",
    "sobrenome": "Messi",
    "rg": 2434322,
    "data_cadastro": "2000-02-02"
}
</code></pre>

## Dentistas
<pre><code>{
    "nome": "Pedro",
    "sobrenome": "Almodóvar",
    "matricula": 23144514
}
</code></pre>

## Consultas
<pre><code>{
    "data": "2000-05-23",
    "hora": "08:30:22",
    "descricao": "Não chegar atrasado",
    "paciente_id": 1,
    "dentista_id": 1
}
</code></pre>
