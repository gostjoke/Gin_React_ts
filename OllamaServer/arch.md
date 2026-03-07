                ┌─────────────────┐
                │  Frontend UI    │
                │ React / Web     │
                └────────┬────────┘
                         │
                         ▼
                ┌─────────────────┐
                │  AI API Server  │
                │  Python         │
                │  FastAPI/Django │
                └───────┬─────────┘
                        │
        ┌───────────────┼────────────────┐
        ▼                               ▼
 ┌───────────────┐               ┌───────────────┐
 │ Vector DB     │               │ Model Server  │
 │ Chroma/Qdrant │               │ Ollama        │
 └───────┬───────┘               └───────┬───────┘
         │                               │
         ▼                               ▼
   Company Documents                LLM Model
   PDF / DOC / Wiki                 gemma / llama