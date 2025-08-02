# Yandex Smart Captcha Example

This example demonstrates how to integrate Yandex Smart Captcha with the CrowdSec bouncer Traefik plugin.

## What is Yandex Smart Captcha?

Yandex Smart Captcha is an advanced captcha system that goes beyond traditional image-based challenges. It analyzes user behavior patterns, browser fingerprints, and interaction patterns to determine if a user is human or a bot. Unlike traditional captchas, Yandex Smart Captcha can:

- Analyze mouse movements and click patterns
- Monitor browser behavior and timing
- Detect automated interactions
- Present interactive challenges when needed
- Work silently in the background for legitimate users

## Prerequisites

1. **Yandex Cloud Account**: You need a Yandex Cloud account to access Smart Captcha
2. **Smart Captcha Service**: Enable the Smart Captcha service in your Yandex Cloud console
3. **Site Key and Secret Key**: Obtain your Smart Captcha credentials from Yandex Cloud

## Setup Instructions

### 1. Configure Yandex Smart Captcha

1. Go to [Yandex Cloud Console](https://console.cloud.yandex.com/)
2. Navigate to the Smart Captcha service
3. Create a new captcha instance
4. Note down your **Site Key** and **Secret Key**

### 2. Update Configuration

Edit the `docker-compose.yml` file and replace the placeholder values:

```yaml
# Choose Yandex captcha provider
- "traefik.http.middlewares.crowdsec.plugin.bouncer.captchaProvider=yandex"
# Define Yandex captcha site key
- "traefik.http.middlewares.crowdsec.plugin.bouncer.captchaSiteKey=YOUR_SITE_KEY_HERE"
# Define Yandex captcha secret key
- "traefik.http.middlewares.crowdsec.plugin.bouncer.captchaSecretKey=YOUR_SECRET_KEY_HERE"
```

### 3. Run the Example

```bash
# Start the services
docker-compose up -d

# Check the logs
docker-compose logs -f
```

## Configuration Options

### Traefik Configuration

The minimal configuration for Yandex Smart Captcha:

```yaml
labels:
  # Choose Yandex captcha provider
  - "traefik.http.middlewares.crowdsec.plugin.bouncer.captchaProvider=yandex"
  # Define Yandex captcha site key
  - "traefik.http.middlewares.crowdsec.plugin.bouncer.captchaSiteKey=YOUR_SITE_KEY"
  # Define Yandex captcha secret key
  - "traefik.http.middlewares.crowdsec.plugin.bouncer.captchaSecretKey=YOUR_SECRET_KEY"
  # Define captcha grace period seconds
  - "traefik.http.middlewares.crowdsec.plugin.bouncer.captchaGracePeriodSeconds=1800"
  # Define captcha HTML file path
  - "traefik.http.middlewares.crowdsec.plugin.bouncer.captchaHTMLFilePath=/yandex-captcha.html"
```

### CrowdSec Configuration

The CrowdSec configuration remains the same as other captcha providers. You can choose between:

- **Mixed mode** (`profiles.yaml`): Shows captcha for the first 3 violations, then bans
- **Captcha only mode** (`profiles_captcha_only.yaml`): Always shows captcha for HTTP violations

## Testing the Setup

### 1. Normal Access

Test normal access to the service:

```bash
curl http://localhost:8000/foo
```

This should work normally without any captcha challenge.

### 2. Trigger Captcha

Simulate suspicious activity by adding a captcha decision:

```bash
docker exec crowdsec cscli decisions add --ip 10.0.0.20 -d 4h --type captcha
```

Now when you access `http://localhost:8000/foo`, you should see the Yandex Smart Captcha challenge.

### 3. Test Captcha Validation

Complete the captcha challenge in your browser. The system will:
1. Validate the captcha with Yandex's servers
2. Cache the successful validation for the grace period
3. Allow access to the protected resource

## Yandex Smart Captcha Features

### Advanced Bot Detection

Yandex Smart Captcha uses multiple detection methods:

- **Behavioral Analysis**: Monitors mouse movements, click patterns, and interaction timing
- **Browser Fingerprinting**: Analyzes browser characteristics and consistency
- **Network Analysis**: Examines request patterns and timing
- **Device Analysis**: Checks device characteristics and consistency

### User Experience

- **Silent Operation**: Legitimate users often don't see any captcha
- **Progressive Challenges**: Escalates from simple checks to interactive challenges
- **Mobile Friendly**: Works seamlessly on mobile devices
- **Accessibility**: Supports screen readers and assistive technologies

### Security Features

- **Real-time Analysis**: Continuously monitors for suspicious patterns
- **Adaptive Difficulty**: Adjusts challenge complexity based on risk assessment
- **Geographic Awareness**: Considers regional patterns and behaviors
- **Rate Limiting**: Prevents abuse of the captcha system itself

## Troubleshooting

### Common Issues

1. **Captcha Not Loading**
   - Verify your site key is correct
   - Check that the domain is authorized in Yandex Cloud
   - Ensure the JavaScript is loading properly

2. **Validation Failures**
   - Verify your secret key is correct
   - Check network connectivity to Yandex servers
   - Review the plugin logs for validation errors

3. **Performance Issues**
   - Monitor response times from Yandex validation API
   - Consider implementing caching for successful validations
   - Review network latency to Yandex Cloud

### Debug Mode

Enable debug logging to troubleshoot issues:

```yaml
- "traefik.http.middlewares.crowdsec.plugin.bouncer.loglevel=DEBUG"
```

### Log Analysis

Check the logs for captcha-related events:

```bash
# View Traefik logs
docker-compose logs traefik

# View CrowdSec logs
docker-compose logs crowdsec
```

## Integration with Other Services

### Yandex Cloud Integration

Yandex Smart Captcha integrates well with other Yandex Cloud services:

- **Load Balancer**: Can be used with Yandex Application Load Balancer
- **CDN**: Works with Yandex CDN for global distribution
- **Monitoring**: Integrates with Yandex Cloud Monitoring
- **Logging**: Centralized logging with Yandex Cloud Logging

### Third-party Integrations

- **Analytics**: Can integrate with Google Analytics, Yandex Metrica
- **Security**: Works with WAF and other security services
- **Monitoring**: Compatible with Prometheus, Grafana, and other monitoring tools

## Best Practices

### Security

1. **Keep Keys Secure**: Store secret keys in environment variables or secure vaults
2. **Domain Validation**: Only authorize necessary domains in Yandex Cloud
3. **Rate Limiting**: Implement additional rate limiting for captcha endpoints
4. **Monitoring**: Monitor captcha success/failure rates

### Performance

1. **Caching**: Implement appropriate caching for successful validations
2. **CDN**: Use CDN for static captcha assets
3. **Optimization**: Minimize the impact on page load times
4. **Fallback**: Have a fallback mechanism for captcha failures

### User Experience

1. **Progressive Enhancement**: Don't block legitimate users unnecessarily
2. **Clear Messaging**: Provide clear instructions when captcha is required
3. **Accessibility**: Ensure captcha is accessible to all users
4. **Mobile Optimization**: Test thoroughly on mobile devices

## Comparison with Other Captcha Providers

| Feature | Yandex Smart Captcha | reCAPTCHA | hCaptcha | Turnstile |
|---------|---------------------|-----------|----------|-----------|
| Silent Operation | ✅ | ✅ | ✅ | ✅ |
| Behavioral Analysis | ✅ | ✅ | ✅ | ✅ |
| Mobile Optimization | ✅ | ✅ | ✅ | ✅ |
| Privacy Focus | ✅ | ❌ | ✅ | ✅ |
| Regional Support | ✅ | ✅ | ✅ | ✅ |
| Customization | ✅ | ✅ | ✅ | ✅ |

## Support and Resources

- **Yandex Cloud Documentation**: [Smart Captcha Documentation](https://cloud.yandex.com/docs/smartcaptcha/)
- **CrowdSec Documentation**: [Captcha Profile Documentation](https://docs.crowdsec.net/docs/next/profiles/captcha_profile/)
- **Community Support**: [CrowdSec Community](https://discourse.crowdsec.net/)

## License

This example is provided under the same license as the main CrowdSec bouncer Traefik plugin project. 