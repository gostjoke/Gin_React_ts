"use client";

import { useState } from "react";
import { useChat } from "@ai-sdk/react";
import { DefaultChatTransport } from "ai";

export function ChatDemo() {
  const [input, setInput] = useState("");

  const { messages, sendMessage, status, stop, error } = useChat({
    transport: new DefaultChatTransport({
      api: "http://127.0.0.1:8081/chat",
    }),
  });

  const isLoading = status === "submitted" || status === "streaming";

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!input.trim()) return;

    await sendMessage({
      role: "user",
      parts: [{ type: "text", text: input }],
    });

    setInput("");
  };

  return (
    <div style={{ padding: 16 }}>
      <h2>Chat Demo</h2>

      <div style={{ marginBottom: 16 }}>
        {messages.map((message) => (
          <div key={message.id} style={{ marginBottom: 12 }}>
            <strong>{message.role === "user" ? "User" : "AI"}:</strong>{" "}
            {message.parts.map((part, index) =>
              part.type === "text" ? (
                <span key={index}>{part.text}</span>
              ) : null
            )}
          </div>
        ))}
      </div>

      {error ? <div style={{ color: "red" }}>{error.message}</div> : null}

      <form onSubmit={handleSubmit}>
        <input
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="請輸入訊息"
          style={{ width: 300, marginRight: 8 }}
        />
        <button type="submit" disabled={isLoading}>
          {isLoading ? "Thinking..." : "Send"}
        </button>
        <button type="button" onClick={stop} style={{ marginLeft: 8 }}>
          Stop
        </button>
      </form>
    </div>
  );
}