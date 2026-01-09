# Integrations

Connect our platform with your existing tools and workflows.

## Official Integrations

### Slack

Get notifications directly in Slack:

1. Go to Settings > Integrations > Slack
2. Click "Add to Slack"
3. Choose the channel for notifications
4. Select which events to receive

Events you can subscribe to:
- Document created/updated/deleted
- Payment events
- Security alerts
- Usage warnings (approaching limits)

### Zapier

Connect with 5,000+ apps via Zapier:

1. Search for our app in Zapier
2. Connect your account using an API key
3. Build zaps triggered by our webhooks

Popular zaps:
- New document → Google Drive backup
- Payment failed → Slack alert
- New user → Add to Mailchimp

### GitHub

Sync documents with your repositories:

1. Go to Settings > Integrations > GitHub
2. Install our GitHub App
3. Select repositories to sync
4. Configure sync direction (one-way or bidirectional)

## Building Custom Integrations

### Using Webhooks

Webhooks are the foundation for custom integrations:

1. Set up an endpoint to receive webhook events
2. Register your endpoint in Settings > Webhooks
3. Subscribe to relevant events
4. Process events and trigger your custom logic

See the Webhooks documentation for:
- Available event types
- Payload formats
- Signature verification
- Retry behavior

### Using the API

For pull-based integrations:

1. Generate an API key with appropriate scopes
2. Use our REST API to read/write data
3. Respect rate limits (see Rate Limits docs)
4. Implement proper error handling

### SDKs

Official SDKs available for:
- Python: `pip install example-sdk`
- Node.js: `npm install @example/sdk`
- Go: `go get github.com/example/sdk-go`
- Ruby: `gem install example-sdk`

SDKs handle authentication, retries, and rate limiting automatically.

## Authentication for Integrations

Different integrations need different auth approaches:

| Integration Type | Recommended Auth |
|-----------------|------------------|
| Server-to-server | API Key |
| User-facing app | OAuth 2.0 |
| Webhooks (inbound) | Signature verification |
| Zapier/Integromat | API Key via OAuth flow |

See Authentication documentation for implementation details.

## Rate Limits for Integrations

Integration-specific limits:

| Integration | Limit |
|-------------|-------|
| Slack notifications | 10/minute per channel |
| Webhook deliveries | 1,000/hour |
| GitHub sync | 100 files/sync |

These are separate from your API rate limits.

## Troubleshooting Integrations

### Slack Not Receiving Notifications

1. Check the integration is still connected (Settings > Integrations)
2. Verify the channel still exists
3. Ensure the Slack app wasn't removed from your workspace

### Webhook Integration Failing

1. Check your endpoint is returning 2xx status
2. Verify the endpoint is publicly accessible
3. Confirm you're validating signatures correctly
4. Check for IP allowlisting issues

For persistent issues, enable webhook logging in Settings > Webhooks > Debug Mode.




