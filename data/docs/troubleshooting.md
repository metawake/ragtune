# Troubleshooting

Common issues and how to resolve them.

## Authentication Errors

### 401 Unauthorized

**Cause**: Invalid or missing authentication credentials.

**Solutions**:
1. Verify your API key is correct
2. Check the key hasn't expired (Settings > API Keys)
3. Ensure proper header format: `Authorization: Bearer YOUR_KEY`
4. If using OAuth, check if the access token has expired

### 403 Forbidden

**Cause**: Valid authentication but insufficient permissions.

**Solutions**:
1. Check your API key's scopes in Settings > API Keys
2. Verify your subscription tier allows this endpoint
3. If IP allowlisting is enabled, ensure your IP is listed
4. Check if the resource belongs to your account

## Rate Limit Errors

### 429 Too Many Requests

**Cause**: You've exceeded your rate limit.

**Solutions**:
1. Check rate limit headers in the response:
   - `X-RateLimit-Remaining`: Requests left
   - `X-RateLimit-Reset`: When limit resets
2. Implement exponential backoff
3. Cache responses where possible
4. Consider upgrading your plan for higher limits

**Rate limits by tier**:
- Free: 60/minute, 1,000/day
- Pro: 600/minute, 100,000/day
- Enterprise: Custom

## Webhook Issues

### Webhooks Not Arriving

1. Verify your endpoint URL is correct and HTTPS
2. Check your endpoint returns 2xx status codes
3. Look for the endpoint in Settings > Webhooks — is it disabled?
4. Check if you've had 5 consecutive failures (auto-disables)

### Webhook Signature Validation Failing

1. Ensure you're using the correct webhook secret
2. Verify you're computing HMAC-SHA256 correctly
3. Check you're using the raw request body (not parsed JSON)

## Payment Issues

### Payment Failed

1. Verify your card details in Settings > Billing
2. Check your card hasn't expired
3. Ensure sufficient funds/credit
4. Contact your bank if declines persist

### Subscription Downgrade Issues

Downgrades take effect at the end of your billing cycle. If you need immediate downgrade, contact support.

## API Response Errors

### 400 Bad Request

Your request body is malformed. Check:
1. JSON syntax is valid
2. Required fields are present
3. Field types match the schema
4. Dates are in ISO 8601 format

### 404 Not Found

The resource doesn't exist. Verify:
1. The resource ID is correct
2. The resource hasn't been deleted
3. You have access to this resource

### 500 Internal Server Error

This is on us. Steps:
1. Retry after a few seconds
2. Check status.example.com for outages
3. If persistent, contact support with your request ID

## Performance Issues

### Slow API Responses

1. Check status.example.com for degradation
2. Verify you're hitting the nearest regional endpoint
3. Reduce payload sizes where possible
4. Use pagination for large result sets

### Timeout Errors

1. Increase your client timeout settings
2. For long operations, use async endpoints with webhooks
3. Break large batch operations into smaller chunks

## Still Stuck?

1. **Search docs**: Use the search bar above
2. **Community**: community.example.com
3. **Support**: support@example.com
4. **Status**: status.example.com




