# Security

## Our Security Commitment

We take security seriously. Our platform is SOC 2 Type II certified and undergoes regular third-party penetration testing.

## Data Protection

### Encryption

- All data encrypted at rest using AES-256
- All data encrypted in transit using TLS 1.3
- API keys are hashed and never stored in plain text

### Data Residency

Enterprise customers can choose data residency:
- US (default)
- EU (GDPR compliant)
- Asia Pacific

## API Security

### Key Management

Protect your API keys:

- Store keys in environment variables, never in code
- Use separate keys for development and production
- Rotate keys immediately if you suspect compromise
- Set expiration dates on keys when possible

If a key is compromised:
1. Go to Settings > API Keys
2. Click "Revoke" on the compromised key
3. Generate a new key
4. Update your applications

### IP Allowlisting

Restrict API access to specific IP addresses:

1. Go to Settings > Security > IP Allowlist
2. Add your server IP addresses
3. Enable "Enforce IP Allowlist"

Requests from non-allowlisted IPs will receive a 403 error.

### Request Signing

For additional security, enable request signing:

1. Generate a signing secret in Settings > Security
2. Sign each request with HMAC-SHA256
3. Include the signature in the `X-Signature` header

## Webhook Security

Verify webhook authenticity:

- Each webhook includes an `X-Signature` header
- Verify the signature using your webhook secret
- Reject requests with invalid signatures
- Use HTTPS endpoints only

See Webhooks documentation for implementation details.

## Rate Limiting as Security

Rate limits protect against abuse:

- Prevents brute-force attacks on authentication
- Mitigates DDoS attempts
- Ensures fair resource usage

Suspicious activity triggers automatic blocking. Contact support if you're unexpectedly blocked.

## Incident Response

If you discover a security vulnerability:

1. Email security@example.com
2. Do not disclose publicly until we've addressed it
3. We aim to respond within 24 hours

We offer a bug bounty program for qualifying vulnerabilities.

## Compliance

- SOC 2 Type II certified
- GDPR compliant
- HIPAA compliant (Enterprise plan)
- PCI DSS compliant for payment processing




