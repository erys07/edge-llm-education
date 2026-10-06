# First Inference Test

## Model

- Model: SmolLM2-135M-Instruct
- Format: GGUF
- Precision: F16
- Size: 270886368 bytes (~258 MiB)
- Runtime: llama.cpp b11429-d81235049

## Prompt

Explique de forma simples para um aluno: quanto é 12 vezes 8?

## Response

12 x 8 = 96.

## Observed performance

Prompt processing: 438.1 tokens/s
Generation: 249.2 tokens/s

## Result

The model loaded successfully and produced a correct response in Portuguese.

This test confirms functional compatibility between the converted
SmolLM2-135M-Instruct F16 GGUF artifact and the installed llama.cpp runtime.

The performance values reported here correspond to a single interactive
execution and should not be treated as the project's formal benchmark.
