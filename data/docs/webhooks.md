# Webhooks

## Overview

Webhooks allow your application to receive real-time notifications when events occur in our system.

## Supported Events

- `document.created` — A new document was created
- `document.updated` — A document was modified
- `document.deleted` — A document was deleted
- `user.signup` — A new user registered
- `payment.success` — Payment was processed successfully
- `payment.failed` — Payment failed

## Setting Up Webhooks

1. Go to Settings > Webhooks
2. Click "Add Endpoint"
3. Enter your HTTPS endpoint URL
4. Select the events you want to receive
5. Save the endpoint

## Webhook Payload

All webhooks are sent as POST requests with JSON body:

```json
{
  "event": "document.created",
  "timestamp": "2024-01-15T10:30:00Z",
  "data": {
    "id": "doc_123",
    "name": "example.pdf"
  }
}
```

## Security

Each webhook request includes an `X-Signature` header. Verify this signature using your webhook secret to ensure the request came from us.

## Retry Policy

Failed webhooks are retried up to 5 times with exponential backoff. After 5 failures, the endpoint is disabled and you'll receive an email notification.




