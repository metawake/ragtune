# Rate Limits

## Default Limits

Rate limits are applied per API key:

| Tier       | Requests/minute | Requests/day |
|------------|-----------------|--------------|
| Free       | 60              | 1,000        |
| Pro        | 600             | 100,000      |
| Enterprise | Custom          | Unlimited    |

## Rate Limit Headers

Every API response includes rate limit headers:

- `X-RateLimit-Limit` — Maximum requests allowed
- `X-RateLimit-Remaining` — Requests remaining in window
- `X-RateLimit-Reset` — Unix timestamp when the limit resets

## Handling Rate Limits

When you exceed the rate limit, you'll receive a `429 Too Many Requests` response.

Best practices:
1. Implement exponential backoff
2. Cache responses when possible
3. Batch requests where supported
4. Monitor your usage in the dashboard

## Requesting Higher Limits

Enterprise customers can request custom rate limits. Contact support@example.com with your use case and expected traffic patterns.




