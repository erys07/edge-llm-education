# EXP-001 — SmolLM2-135M-Instruct F16 Baseline

## Objetivo

Avaliar se um modelo de aproximadamente 135 milhões de parâmetros é capaz
de fornecer capacidades educacionais básicas em português antes da aplicação
de técnicas de quantização.

Este experimento representa o primeiro ponto da busca pela menor configuração
computacional capaz de fornecer IA generativa educacional útil de forma
totalmente offline.

## Modelo

- Modelo: HuggingFaceTB/SmolLM2-135M-Instruct
- Parâmetros reportados pelo llama.cpp: 134.52 M
- Formato: GGUF
- Quantização: F16
- Arquivo: smollm2-135m-instruct-f16.gguf
- Tamanho: 270,886,368 bytes (~258 MiB)
- SHA-256: 96f58508ce65f1546e40488ae77c6a5316975a5b94f474f1fc5b2510da945ed9

## Runtime

- llama.cpp / llama-cli
- versão: 0.6.0
- build: 11429
- commit: d81235049
- backend utilizado na avaliação educacional: CPU (-ngl 0)
- temperatura: 0
- máximo de tokens gerados: 128
- execução independente para cada prompt

## Hardware

Baseline executado em Apple M5 com 16 GB de memória.

O hardware utilizado neste experimento não representa o hardware-alvo final.
Ele é utilizado para caracterização inicial e estabelecimento da qualidade
do modelo antes dos testes em dispositivos restritos.

## Desempenho

Benchmark com llama-bench:

- Prompt processing (pp512): ~23k tokens/s
- Text generation (tg128): ~300 tokens/s
- Modelo: ~256.63 MiB
- Parâmetros: 134.52 M

Os números representam execução no hardware de desenvolvimento e não devem
ser extrapolados diretamente para SBCs de baixo custo.

## Memória

Execução CPU medida com `/usr/bin/time -l`:

- Maximum resident set size: 575,963,136 bytes
- Aproximadamente: 549 MiB
- Swaps: 0

Esse valor representa a configuração testada e não deve ser interpretado como
um limite mínimo absoluto de memória necessário para executar o modelo.

## Avaliação educacional

Dataset:

`datasets/educational-v0.1/prompts.json`

Total: 30 prompts.

Categorias:

- aritmética
- raciocínio matemático
- português
- ciências
- pedagogia
- seguimento de instruções

O modelo apresentou falhas substanciais nas diferentes categorias avaliadas.

Exemplos observados incluem:

- erro em 12 × 8;
- erro em 25 + 37;
- incapacidade de calcular 15% de 200;
- erro no plural de "animal";
- falha ao identificar Marte como Planeta Vermelho;
- falhas em explicações pedagógicas básicas;
- dificuldade em seguir instruções simples;
- repetição e geração de texto incoerente em diversos prompts.

## Resultado

O SmolLM2-135M-Instruct em F16 foi considerado insuficiente como modelo
educacional geral em português no protocolo experimental utilizado.

A execução técnica foi bem-sucedida, porém a qualidade das respostas ficou
abaixo do necessário para justificar, neste momento, experimentos adicionais
de quantização desse modelo.

## Decisão

Não avançar imediatamente com Q8, Q4, Q3 ou Q2 para o modelo de 135M.

O próximo experimento deverá aumentar a capacidade do modelo mantendo, na
medida do possível, o mesmo protocolo e o mesmo conjunto de avaliação.

Próximo candidato:

SmolLM2-360M-Instruct F16.

Se o modelo de 360M apresentar qualidade educacional significativamente
superior, ele poderá então ser submetido a quantização para investigar a
fronteira entre qualidade, memória, desempenho e custo computacional.

## Conclusão

O EXP-001 fornece evidência inicial de que aproximadamente 135M parâmetros,
na configuração avaliada, estão abaixo da capacidade necessária para o caso
de uso educacional geral proposto.

Esse resultado estabelece um primeiro ponto inferior na busca experimental
pela menor IA generativa educacional offline considerada útil.
