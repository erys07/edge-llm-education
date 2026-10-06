# Model

Model: SmolLM2-135M-Instruct
Format: GGUF
Precision: F16

## Artifact

File: smollm2-135m-instruct-f16.gguf
Size: 270886368 bytes (~258 MiB)

SHA-256:
96f58508ce65f1546e40488ae77c6a5316975a5b94f474f1fc5b2510da945ed9

## Status

Conversion completed successfully.

Performance, memory usage and educational quality have not yet been measured.

## CPU Memory Measurement

Command configuration:
- CPU inference (`-ngl 0`)
- Maximum generation: 128 tokens
- Temperature: 0

Measured maximum resident set size:
575963136 bytes (~549.3 MiB)

Swaps: 0

Note: this represents the observed peak RSS for this particular
llama.cpp execution on macOS. It must not be interpreted as the
absolute minimum RAM requirement of the model.

## Qualitative Observation

Prompt:
"Explique para um aluno por que 12 vezes 8 é igual a 96."

Response:
"12 x 96 = 1080."

The response is mathematically incorrect and does not provide the
requested explanation.

This single observation is not sufficient to characterize model
quality. A fixed educational evaluation dataset will be created for
systematic comparison between model variants and quantization levels.
