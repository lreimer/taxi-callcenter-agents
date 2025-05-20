FROM ghcr.io/astral-sh/uv::0.7.6-python3.12-bookworm

WORKDIR /app

COPY . /app
RUN uv sync --locked

ENV PATH="/app/.venv/bin:/root/.local/bin:$PATH"
RUN uv tool install google-adk

EXPOSE 8000
CMD [ "adk", "web" "--port", "8000"]