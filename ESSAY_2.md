# Essay 2 - AI Integration Design

## 1. Architecture & Boundaries
AI is treated as an external service, similar to other third-party APIs.

The core Integration Backend communicates with the AI service through well-defined HTTP endpoints. Most AI workloads (such as topic normalisation or content enrichment) run asynchronously via background jobs, since they are not user-facing and may be slow.

The main backend never blocks critical user flows on AI responses. If the AI service is unavailable, the system falls back to deterministic logic or previously stored results.

All AI outputs are persisted in the database with metadata such as input hash and model version to avoid repeated processing and enable safe reprocessing.

## 2. AI vs Deterministic Logic
AI is only used for tasks that benefit from semantic understanding, such as:

- normalising free-text topics
- mapping topics to books
- extracting tags

Core business logic (recommendation ranking, validation, access control, aggregation) remains fully deterministic.

This separation ensures the system stays predictable and debuggable, while AI acts only as an enrichment layer.

## 3. Data Flow & Reliability

Publisher data is synced in batches, while student engagement is ingested near real-time.

AI processing happens after ingestion:

1. Data is saved first.
2. Background jobs trigger AI enrichment.
3. Results are stored back into canonical tables.

All ingestion and AI jobs are idempotent, so retries are safe.

API calls use retry with exponential backoff. Failed jobs are pushed into a dead-letter queue for later inspection and reprocessing.


## 4. Cost, Latency & Scalability

To control cost, AI requests are batched where possible and cached using input hashes to prevent duplicate calls.

Only changed or new records are sent to AI.

The Integration Backend scales horizontally behind a load balancer, while AI workers scale independently based on queue depth.

Because most AI tasks run asynchronously, user-facing latency is not impacted.

---

## 5. Observability & Quality Control

Structured logging is added for each AI job, including input size, duration, and failures.

Metrics track:

- AI request latency
- throughput
- error rate

Periodic validation jobs check for missing mappings or inconsistent outputs.

Manual review tools can be added for edge cases where AI confidence is low.

---

## 6. Security & Privacy

Only minimal required data is sent to the AI service.

Sensitive user information is excluded or anonymised.

All communication is encrypted, and API keys are stored securely.

AI outputs are treated as derived data and follow the same access control rules as core entities.