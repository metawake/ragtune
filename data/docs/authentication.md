# Authentication

## Overview

All API requests must be authenticated. We support multiple authentication methods depending on your use case.

## Authentication Methods

### API Key Authentication

The simplest method. Include your API key in the `Authorization` header:

```
Authorization: Bearer your-api-key
```

API keys are managed in Settings > API Keys. Each key can have scoped permissions and expiration dates. See the API Keys documentation for details on creating, rotating, and revoking keys.

### OAuth 2.0

For applications acting on behalf of users, use OAuth 2.0:

1. Register your application in Settings > OAuth Apps
2. Redirect users to our authorization endpoint
3. Exchange the authorization code for access tokens
4. Use the access token in the `Authorization` header

Access tokens expire after 1 hour. Use refresh tokens to obtain new access tokens.

### Session-Based Authentication

For browser-based applications, you can use session cookies after a user logs in. Sessions expire after 24 hours of inactivity.

## Security Best Practices

- Never expose API keys in client-side code
- Rotate keys regularly (we recommend every 90 days)
- Use the minimum required permissions for each key
- Monitor API key usage in the Analytics dashboard
- Enable IP allowlisting for production keys
- Set up alerts for unusual API activity

## Rate Limits by Auth Method

Different authentication methods have different rate limits:

| Method | Requests/minute |
|--------|-----------------|
| API Key (Free) | 60 |
| API Key (Pro) | 600 |
| OAuth Token | 300 |
| Session | 120 |

If you exceed these limits, you'll receive a 429 error. See Rate Limits documentation for details.

## Troubleshooting Authentication

Common authentication errors:

- `401 Unauthorized`: Invalid or expired credentials
- `403 Forbidden`: Valid credentials but insufficient permissions
- `429 Too Many Requests`: Rate limit exceeded

For persistent issues, check that your API key hasn't been revoked and that your subscription is active.




