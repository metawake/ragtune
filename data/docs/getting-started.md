# Getting Started

Welcome! This guide will help you make your first API call in under 5 minutes.

## Step 1: Create an Account

Sign up at dashboard.example.com. You'll start on the Free tier with 1,000 API calls per month.

## Step 2: Get Your API Key

1. Go to Settings > API Keys
2. Click "Create New Key"
3. Copy the key — it won't be shown again!

Store your API key securely. Never commit it to version control or expose it in client-side code.

## Step 3: Make Your First Request

```bash
curl -X GET "https://api.example.com/v1/documents" \
  -H "Authorization: Bearer YOUR_API_KEY"
```

You should receive a JSON response with your documents (empty array if you haven't created any yet).

## Step 4: Create a Document

```bash
curl -X POST "https://api.example.com/v1/documents" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"name": "My First Document", "content": "Hello, World!"}'
```

## Next Steps

Now that you're set up:

- **Explore the API**: See our full API Reference
- **Set up webhooks**: Get notified when events happen (see Webhooks docs)
- **Upgrade your plan**: Need more API calls? See Billing docs
- **Secure your integration**: Review our Security best practices

## Common Issues

### "401 Unauthorized" Error

- Check that your API key is correct
- Ensure the key hasn't expired or been revoked
- Verify you're using the `Authorization: Bearer` format

### "429 Too Many Requests" Error

You've hit the rate limit. Free tier allows 60 requests per minute. Either:
- Wait and retry with exponential backoff
- Upgrade to Pro for higher limits (see Billing)

### "403 Forbidden" Error

Your API key doesn't have permission for this action. Check the key's scopes in Settings > API Keys.

## Getting Help

- **Documentation**: You're reading it!
- **Community Forum**: community.example.com
- **Email Support**: support@example.com (Pro and Enterprise)
- **Slack**: Enterprise customers get a dedicated Slack channel




