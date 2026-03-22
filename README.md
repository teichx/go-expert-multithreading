# Desafio Multithreading

- Regras consideradas:
  - Para a definição: `O sistema deve aceitar apenas a resposta da API que responder mais rápido e descartar a resposta da outra (mais lenta).` foi considerada a API que responder mais rápida, desde que seja com status de sucesso (status code 200). Caso a API retorne primeiro, mas com erro, a resposta é descartada também.

# Como executar

```bash
go run . <cep>
```

- Exemplos:

```bash
go run . 13330-250
go run . 13330250
```

# Exemplo de saída

## ViaCEP

- OBS: Os dados foram formatados automaticamente pelo editor, no terminal a formatação estará diferente.

```json
{
  "api": "ViaCEP",
  "data": {
    "cep": "13330-250",
    "logradouro": "Rua Hércules Mazzoni",
    "complemento": "",
    "unidade": "",
    "bairro": "Centro",
    "localidade": "Indaiatuba",
    "uf": "SP",
    "estado": "São Paulo",
    "regiao": "Sudeste",
    "ibge": "3520509",
    "gia": "3530",
    "ddd": "19",
    "siafi": "6511"
  }
}
```

## BrasilAPI

```json
{
  "api": "BrasilAPI",
  "data": {
    "cep": "13330250",
    "state": "SP",
    "city": "Indaiatuba",
    "neighborhood": "Centro",
    "street": "Rua Hércules Mazzoni",
    "service": "open-cep"
  }
}
```

## Error

```json
{ "error": "timeout - no API responded within 1 second" }
```
