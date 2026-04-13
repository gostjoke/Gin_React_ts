from fastapi import FastAPI
from ollama_client import generate_response
from pydantic import BaseModel
from OllamaServer.utilities import log_request_response

class ChatRequest(BaseModel):
    message: str


class ChatResponse(BaseModel):
    answer: str

app = FastAPI(title="Company AI Assistant")


@app.get("/")
def health():
    return {"status": "ok"}


@app.post("/chat", response_model=ChatResponse)
def chat(req: ChatRequest):

    answer = generate_response(req.message)

    return {"answer": answer}



# uv run uvicorn main:app --reload

if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, host="127.0.0.1", port=8081)