from rich.console import Console
from rich.markdown import Markdown

console = Console()

def pretty_print_markdown(text: str):
    md = Markdown(text)
    console.print(md)