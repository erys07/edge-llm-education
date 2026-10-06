# Seleção de Modelos

## SmolLM2-135M-Instruct

### Motivação

O SmolLM2-135M-Instruct foi escolhido como o primeiro modelo experimental deste projeto não por ser considerado previamente o melhor modelo para a tarefa, mas por representar um ponto inicial extremamente restrito para investigar a seguinte questão:

> Quão pequeno pode ser um modelo de linguagem generativo e ainda manter utilidade em aplicações educacionais offline?

O objetivo da pesquisa é investigar os limites mínimos de infraestrutura computacional necessários para executar IA generativa educacional em ambientes com restrições severas de memória, processamento, energia, custo e conectividade.

### Por que começar com 135M parâmetros?

Com aproximadamente 135 milhões de parâmetros, o modelo é pequeno quando comparado aos LLMs modernos.

Isso permite investigar cenários de hardware extremamente limitado, incluindo dispositivos com poucas centenas de megabytes de memória.

O modelo também pode ser quantizado para representações menores, permitindo estudar experimentalmente a relação entre:

- tamanho do modelo;
- quantização;
- consumo de memória;
- velocidade de inferência;
- consumo de energia;
- qualidade das respostas;
- utilidade educacional.

### Por que SmolLM2?

A família SmolLM2 possui modelos de diferentes tamanhos, permitindo aumentar gradualmente a capacidade computacional durante os experimentos.

Uma possível progressão experimental é:

SmolLM2-135M → SmolLM2-360M → modelos maiores

Isso permite investigar em qual ponto o aumento de parâmetros passa a produzir qualidade educacional suficiente e qual é o custo computacional dessa melhoria.

### Por que a versão Instruct?

A variante Instruct foi escolhida porque já foi ajustada para seguir instruções e responder a solicitações do usuário.

Isso aproxima o comportamento inicial do modelo do cenário educacional pretendido, como:

- responder perguntas;
- explicar conceitos;
- gerar exercícios;
- fornecer exemplos;
- auxiliar professores e estudantes.

### Hipótese importante

Não assumimos que o SmolLM2-135M será suficientemente bom.

Um possível resultado experimental é descobrir que modelos dessa escala possuem desempenho educacional insuficiente.

Esse resultado continua sendo relevante, pois ajuda a identificar o limite inferior de capacidade necessário para uma aplicação educacional útil.

Por exemplo:

135M → executável no hardware, mas qualidade insuficiente  
360M → qualidade aceitável  
1B+ → qualidade superior, porém custo computacional maior

O objetivo é encontrar experimentalmente esse ponto de equilíbrio.

### Especialização educacional

Uma etapa posterior poderá comparar o modelo original com uma versão especializada para educação e português brasileiro.

Isso permitirá investigar se um modelo pequeno especializado pode apresentar desempenho comparável a um modelo genérico maior.

Exemplo:

SmolLM2-135M genérico
        ↓
avaliação educacional PT-BR
        ↓
fine-tuning educacional
        ↓
SmolLM2-135M especializado
        ↓
nova avaliação

Uma hipótese particularmente relevante é:

> Um modelo pequeno especializado em um domínio educacional pode oferecer utilidade semelhante à de modelos genéricos maiores, utilizando significativamente menos recursos computacionais.

### Papel do SmolLM2-135M na pesquisa

O SmolLM2-135M-Instruct deve ser tratado como um ponto experimental inicial, e não como a escolha definitiva da solução.

A pesquisa poderá concluir que:

- 135M é suficiente;
- 135M é insuficiente;
- especialização torna 135M suficiente;
- 360M ou outra escala representa melhor compromisso;
- outra família de modelos apresenta resultados superiores.

A escolha final deverá ser determinada pelos resultados experimentais.

### Critério central

A pergunta não é:

> "Qual é o melhor LLM?"

Mas sim:

> "Qual é o menor sistema capaz de oferecer uma experiência educacional útil sob restrições severas de hardware, custo e conectividade?"
