import requests
from config import OLLAMA_URL, MODEL_NAME
from utilities import pretty_print_markdown


def generate_response(prompt: str):
    url = f"{OLLAMA_URL}/api/generate"

    payload = {
        "model": MODEL_NAME,
        "prompt": prompt,
        "stream": False
    }

    r = requests.post(url, json=payload)

    # print("status_code =", r.status_code)
    # print("text =", r.text)

    data = r.json()
    response = data.get("response", None)
    context = data.get("context", None) # 聊天記憶 (如果模型支持的話)
    pretty_print_markdown(response)
    if response:
        return data["response"]
    else:
        raise Exception(f"Error from Ollama API: {data}")